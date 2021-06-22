package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"github.com/gomodule/redigo/redis"
	"golang.org/x/crypto/bcrypt"
	"san616qi/app/common/consts"
	"san616qi/app/dao"
	"san616qi/app/model"
)

// 中间件管理服务
var User = userService{}

type userService struct{}

//用户注册
func (s *userService) SignUp(r *model.UserServiceSignUpReq) error {

	if r.Nickname == "" {
		r.Nickname = r.Passport
	}

	//先走redis，redis如果查询报错，也返回错误
	//ifExist1, err := redis.Bool(g.Redis().Do("EXISTS", fmt.Sprintf(r.Passport, "-port")))
	ifExist1, err := redis.Bool(g.Redis().Do("EXISTS", r.Passport+"-port"))
	if err != nil {
		return errors.New(fmt.Sprintf("redis 查询出现错误:%s", err))
	}

	//redis已经有了，那就直接返回错误，不能存在重复的账号
	if ifExist1 {
		return errors.New(fmt.Sprintf("redis 账号 %s 已经存在", r.Passport))
	} else {

		// 账号唯一性数据检查,走mysql
		if !s.CheckPassport(r.Passport) {
			//redis内的账号信息丢失，补充回redis
			//为了防止A用户的nickname和B用户的passport相同导致的判断错误
			//在设置key值得时候加一个后缀
			if _, errSetPassport := g.Redis().Do("SET", r.Passport+"-port",
				r.Passport, "EX", consts.RedisExpireTime); errSetPassport != nil {
				return errSetPassport
			}

			if _, errSetNickname := g.Redis().Do("SET", r.Nickname+"-nick",
				r.Nickname, "EX", consts.RedisExpireTime); errSetNickname != nil {
				return errSetNickname
			}
			return errors.New(fmt.Sprintf("mysql 账号 %s 已经存在", r.Passport))
		}

		//检测redis有无重复的nickname
		ifExist2, err := redis.Bool(g.Redis().Do("EXISTS", r.Nickname+"-nick"))
		if err != nil {
			return errors.New(fmt.Sprintf("查询出现错误:%s", err))
		}
		if ifExist2 {
			return errors.New(fmt.Sprintf("redis 昵称 %s 已经存在", r.Nickname))
		} else {

			// 昵称唯一性数据检查,走mysql
			if !s.CheckNickName(r.Nickname) {
				//redis昵称信息丢失，重新写回
				//为了防止A用户的nickname和B用户的passport相同导致的判断错误
				//在设置key值得时候加一个后缀
				if _, errSetPassport := g.Redis().Do("SET", r.Passport+"-port", r.Passport); errSetPassport != nil {
					return errSetPassport
				}
				if _, errSetNickname := g.Redis().Do("SET", r.Nickname+"-nick", r.Nickname); errSetNickname != nil {
					return errSetNickname
				}
				return errors.New(fmt.Sprintf("mysql 昵称 %s 已经存在", r.Nickname))
			}

			//设置到mysql。并在保存前进行一些必要数据的处理
			//进行密码加密
			var err2 error
			r.Password, err2 = s.PasswordEncode(r.Password)
			if err2 != nil {
				return err2
			}
			//把创建时间赋值
			r.CreateAt = gtime.Now()
			r.LastAccessAt = gtime.Now()

			if _, err := dao.User.Save(r); err != nil {
				return err
			}

			//成功注册后设置到redis中，passport和nickname
			if _, errSetPassport := g.Redis().Do("SET", r.Passport+"-port", r.Passport, "EX", consts.RedisExpireTime); errSetPassport != nil {
				return errSetPassport
			}
			if _, errSetNickname := g.Redis().Do("SET", r.Nickname+"-nick", r.Nickname, "EX", consts.RedisExpireTime); errSetNickname != nil {
				return errSetNickname
			}

		}
	}
	return nil

}

// 用户登录，成功的话就放进context里的hashmap
func (s *userService) SignIn(ctx context.Context, passport, password string) (error, *model.UserProfileRep) {

	//查询数据
	var user *model.User
	//返回VO
	var userVO *model.UserProfileRep

	err := dao.User.Where("passport=?", passport).Scan(&user)
	if err != nil {
		return errors.New("数据库查询错误"), nil
	}
	if user == nil {
		return errors.New("账号错误,不存在该用户"), nil
	}
	if err := s.PasswordDecode(password, user.Password); !err {
		return errors.New("密码错误，请再次输入"), nil
	}

	//进行处理,存入当前登录的user
	//为了防止这里面的usermap是空的，如果为空要进行一次新的赋值
	if err := Context.Get(ctx).Users.UsersMap.IsEmpty(); err {
		Context.Get(ctx).Users.UsersMap = *gmap.New()
	}

	if err := gconv.Struct(user, &userVO); err != nil {
		return errors.New("数据转换失败"), nil
	}

	//进行登录用户的记录
	Context.Get(ctx).Users.UsersMap.Set(passport, userVO)

	// 赋值到Session，以便其他请求检测是否有登陆
	contextUser := &model.ContextUsers{
		UsersMap: Context.Get(ctx).Users.UsersMap,
	}
	if err := Session.SetUser(ctx, contextUser); err != nil {
		return err, nil
	}

	//设定上下文的已登录用户
	Context.SetUsers(ctx, Context.Get(ctx).Users)
	return nil, userVO

}

// 判断用户是否已经登录
func (s *userService) IsSignedIn(ctx context.Context, passport string) bool {
	if v := Context.Get(ctx); v != nil && v.Users.UsersMap.Contains(passport) {
		return true
	}
	return false
}

// 更新用户个人信息
// 增加redis
func (s *userService) UpdateProfile(userId int, r *model.UserServiceUpdateProfileReq) error {

	//更新一下用户access时间
	r.LastAccessAt = gtime.Now()

	//组装redis个人信息key
	key := "userprofile_" + gconv.String(userId%10)
	//组装个人信息field
	field := gconv.String(userId)

	if _, err := g.Redis().Do("HSET",key,field,r); err != nil {
		return errors.New("redis插入失败")
	}

	return nil

}

// 查询用户个人信息
// 增加redis
func (s *userService) QueryProfile(userId int) (error, *model.UserServiceUpdateProfileReq) {

	////查询用户返回的信息
	var user *model.UserServiceUpdateProfileReq

	//组装redis个人信息key
	key := "userprofile_" + gconv.String(userId%10)
	//组装个人信息field
	field := gconv.String(userId)

	//从redis读取并组装
	if v, err := g.Redis().DoVar("HGET", key, field); err != nil {
		return errors.New("redis查询错误"), nil
	} else {
		if !v.IsNil() {
			err := gconv.Struct(v, &user)
			if err != nil {
				return errors.New("数据转换失败"), nil
			}
			return nil, user
		} else {
			if err := dao.User.Where("user_id=?", userId).Scan(&user); err != nil {
				return errors.New("数据库查询失败"), nil
			} else {
				if _, err := g.Redis().Do("HSET", key, field, user); err != nil {
					return errors.New("redis设置错误"), nil
				} else {
					return nil, user
				}
			}
		}
	}

	return nil, nil
}

//
//	CheckPassport
//	@Description:
//	@receiver s *userService
//	@param passport string
//	@return bool
//
// 检查账号是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *userService) CheckPassport(passport string) bool {
	//dao.User是一个userService实例
	if i, err := dao.User.FindCount("passport", passport); err != nil {
		return false
	} else {
		return i == 0
	}
}

//
//	CheckNickName
//	@Description:
//	@receiver s *userService
//	@param nickname string
//	@return bool
//
// 检查昵称是否符合规范(目前仅检查唯一性),存在返回false,否则true
func (s *userService) CheckNickName(nickname string) bool {
	if i, err := dao.User.FindCount("nickname", nickname); err != nil {
		return false
	} else {
		return i == 0
	}
}

// 密码加密
func (s *userService) PasswordEncode(password string) (string, error) {

	//进行加密
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.New("加密处理错误")
	} else {
		return string(hash), nil
	}

}

// 密码解密
func (s *userService) PasswordDecode(loginword, password string) bool {

	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(loginword));
		err != nil {
		return false
	} else {
		return true
	}

}
