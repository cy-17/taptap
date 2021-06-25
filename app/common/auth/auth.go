package auth

import (
	"errors"
	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/util/gconv"
	"golang.org/x/crypto/bcrypt"
	"san616qi/app/common/consts"
	"san616qi/app/common/response"
	"san616qi/app/dao"
	"san616qi/app/model"
	"time"
)

// 定义认证中间件
var CustomGfJWTMiddleware *jwt.GfJWTMiddleware

// 重写该方法
func init() {
	// 重新定义GfJWTMiddleware
	authMiddleware, err := jwt.New(&jwt.GfJWTMiddleware{
		Realm:         "test zone",          // 用于展示中间件的名称
		Key:           []byte("secret key"), // 密钥
		Timeout:       time.Hour * 24 * 7,   // token过期时间
		MaxRefresh:    time.Hour * 24 * 7,
		IdentityKey:   "user_id",                                          // 身份验证的key值
		TokenLookup:   "header: Authorization, query: token, cookie: jwt", // token检索模式，用于提取token-> Authorization
		TokenHeadName: "Bearer",                                           // token在请求头时的名称，默认值为Bearer
		// 客户端在header中传入Authorization 对一个值是Bearer + 空格 + token
		TimeFunc:        time.Now,        // 测试或服务器在其他时区可设置该属性
		Authenticator:   Authenticator,   // 根据登录信息对用户进行身份验证的回调函数
		LoginResponse:   LoginResponse,   // 完成登录后返回的信息，用户可自定义返回数据，默认返回
		RefreshResponse: RefreshResponse, // 刷新token后返回的信息，用户可自定义返回数据，默认返回
		Unauthorized:    Unauthorized,    // 处理不进行授权的逻辑
		//IdentityHandler: auth.IdentityHandler,  // 解析并设置用户身份信息
		PayloadFunc: PayloadFunc, // 登录期间的回调的函数
		//Authorizator:
	})
	if err != nil {
		glog.Fatal("JWT Error:" + err.Error())
	}
	CustomGfJWTMiddleware = authMiddleware
}

//身份验证器用于验证登录参数。
//它必须返回用户数据作为用户标识符，它将存储在声明数组中。
//检查错误（e）以确定适当的错误消息。
func Authenticator(r *ghttp.Request) (interface{}, error) {
	data := r.GetMap()
	//if e := gvalid.CheckMap(data, auth.ValidationRules); e != nil {
	//	return "", jwt.ErrFailedAuthentication
	//}
	var user *model.UserProfileRep
	var u *model.User
	if err := dao.User.Where("passport=?", data["passport"]).Scan(&u); err != nil {
		return nil, errors.New("mysql查询错误")
	} else {

		if u == nil {
			response.JsonExit(r, consts.CurdSelectFailCode, 2, consts.CurdSelectFailMsg, "账户错误")
		} else {

			if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(gconv.String(data["password"]))); err != nil {
				response.JsonExit(r, consts.CurdSelectFailCode, 2, consts.CurdSelectFailMsg, "密码错误")
			}

			err := gconv.Struct(u, &user)
			if err != nil {
				response.JsonExit(r, consts.ServerOccurredErrorCode, 2, consts.ServerOccurredErrorMsg, "类型转换错误")
			}

			var uMap g.Map
			uMap = make(g.Map, 0)

			//赋值
			uMap["user_id"] = user.UserId
			uMap["nickname"] = user.Nickname
			uMap["passport"] = user.Passport

			if user != nil {
				return uMap, nil
			}
		}
	}

	return nil, jwt.ErrFailedAuthentication
}

// LoginResponse用于定义自定义的登录成功回调函数。
func LoginResponse(r *ghttp.Request, code int, token string, expire time.Time) {

	r.Response.Status = code

	_ = r.Response.WriteJson(g.Map{
		"code":   code,
		"token":  token,
		"expire": expire.Format(time.RFC3339),
		"data": "登陆成功",
	})
	r.ExitAll()
}

func RefreshResponse(r *ghttp.Request, code int, token string, expire time.Time) {

	r.Response.Status = code

	_ = r.Response.WriteJson(g.Map{
		"code":   code,
		"token":  token,
		"expire": expire.Format(time.RFC3339),
		"data": "刷新成功",
	})
	r.ExitAll()
}

func Unauthorized(r *ghttp.Request, code int, message string) {

	r.Response.Status = 400

	_ = r.Response.WriteJson(g.Map{
		"code": code,
		"msg":  message,
		"data": "token缺失",
	})
	r.ExitAll()
}

// PayloadFunc is a callback function that will be called during login.
// Using this function it is possible to add additional payload data to the webtoken.
// The data is then made available during requests via c.Get("JWT_PAYLOAD").
// Note that the payload is not encrypted.
// The attributes mentioned on jwt.io can't be used as keys for the map.
// Optional, by default no additional data will be set.
func PayloadFunc(data interface{}) jwt.MapClaims {
	claims := jwt.MapClaims{}
	params := data.(map[string]interface{})
	if len(params) > 0 {
		for k, v := range params {
			claims[k] = v
		}
	}
	return claims
}

// IdentityHandler get the identity from JWT and set the identity for every request
// Using this function, by r.GetParam("id") get identity
//func IdentityHandler(r *ghttp.Request) interface{} {
//	claims := jwt.ExtractClaims(r)
//	return claims[Auth.IdentityKey]
//}
