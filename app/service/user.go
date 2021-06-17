package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gomodule/redigo/redis"
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
				r.Passport,"EX",consts.RedisExpireTime); errSetPassport != nil {
				return errSetPassport
			}

			if _, errSetNickname := g.Redis().Do("SET", r.Nickname+"-nick",
				r.Nickname,"EX",consts.RedisExpireTime); errSetNickname != nil {
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
			if _, err := dao.User.Save(r); err != nil {
				return err
			}

			//成功注册后设置到redis中，passport和nickname
			if _, errSetPassport := g.Redis().Do("SET", r.Passport+"-port", r.Passport,"EX",consts.RedisExpireTime); errSetPassport != nil {
				return errSetPassport
			}
			if _, errSetNickname := g.Redis().Do("SET", r.Nickname+"-nick", r.Nickname,"EX",consts.RedisExpireTime); errSetNickname != nil {
				return errSetNickname
			}

		}
	}
	return nil

}

// 用户登录，成功的话就放进context里的hashmap
func (s *userService) SignIn(ctx context.Context, passport, password string) (error, *model.User){

	var user *model.User
	err := dao.User.Where("passport=? and password=?", passport, password).Scan(&user)
	if err != nil {
		return errors.New("数据库查询错误"), &model.User{}
	}
	if user == nil {
		return errors.New("账号或者密码错误"), &model.User{}
	}

	//进行处理,存入当前登录的user
	//为了防止这里面的usermap是空的，如果为空要进行一次新的赋值
	if err := Context.Get(ctx).Users.UsersMap.IsEmpty(); err {
		Context.Get(ctx).Users.UsersMap = *gmap.New()
	}

	//进行登录用户的记录
	Context.Get(ctx).Users.UsersMap.Set(passport,user)

	// 赋值到Session，以便其他请求检测是否有登陆
	contextUser := &model.ContextUsers{
		UsersMap: Context.Get(ctx).Users.UsersMap,
	}
	if err := Session.SetUser(ctx, contextUser); err != nil {
		return err,nil
	}

	//设定上下文的已登录用户
	Context.SetUsers(ctx,Context.Get(ctx).Users)
	return nil,user

}

// 判断用户是否已经登录
func (s *userService) IsSignedIn(ctx context.Context, passport string) bool {
	if v := Context.Get(ctx); v!= nil && v.Users.UsersMap.Contains(passport)  {
		return true
	}
	return false
}

// 更新用户个人信息
func (s *userService) UpdateProfile(userId int, r *model.UserServiceUpdateProfileReq) error {

	//更新一下用户access时间
	r.LastAccessAt = gtime.Now()

	if _, err := dao.User.Where("user_id=?",userId).Update(r); err != nil{
		return errors.New("数据库更新出错")
	}
	return nil

}

// 查询用户个人信息
func (s *userService) QueryProfile(userId int) (error, *model.User) {

	////查询用户返回的信息
	var user *model.User

	if err := dao.User.Where("user_id=?",userId).Scan(&user); err != nil {
		return errors.New("数据库查询失败"),nil
	} else {
		return nil, user
	}

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
