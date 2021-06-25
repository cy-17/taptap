package main

import (
	jwt "github.com/gogf/gf-jwt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"net/http"
	"san616qi/app/common/auth"
	"time"
)

// 注册输入参数
type UserServiceSignUpReq struct {
	Passport string
	Password string
	Nickname string
	Aasdas   string `default:"123"`
	Aqqqqq   string `default:"123"`
	Ammmmm   string `default:"123"`
}

// 注册请求参数，用于前后端交互参数格式约定
type UserApiSignUpReq struct {
	Passport  string `v:"required|length:6,16#账号不能为空|账号长度应当在:min到:max之间"`
	Password  string `v:"required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间"`
	Password2 string `v:"required|length:6,16|same:Password#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等"`
	Nickname  string
}

const (
	//本地保存的文件夹名称
	upload_path string = "/files/"
)

var (
	//BUCKET是你在存储空间的名称
	accessKey = "o1CWnMXIWcff8H4umKJt0_TGkT634fYc8MuOwEAQ"
	secretKey = "Zz3FBrHWN9JVUawhaitzpeRuKIvRo2Lm3BVjtTuK"
	bucket1 = "scutsanqi"
)

// 自定义返回值结构体
type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}

// 定义自定义中间件
var CustomGfJWTMiddleware *jwt.GfJWTMiddleware

// 重写该方法
func init () {
	// 重新定义GfJWTMiddleware
	authMiddleware, err := jwt.New(&jwt.GfJWTMiddleware{
		Realm:           "test zone",           // 用于展示中间件的名称
		Key:             []byte("secret key"),  // 密钥
		Timeout:         time.Minute * 5,       // token过期时间
		MaxRefresh:      time.Minute * 5,
		IdentityKey:     "id",                  // 身份验证的key值
		TokenLookup:     "header: Authorization, query: token, cookie: jwt",    // token检索模式，用于提取token-> Authorization
		TokenHeadName:   "Bearer",              // token在请求头时的名称，默认值为Bearer
		// 客户端在header中传入Authorization 对一个值是Bearer + 空格 + token
		TimeFunc:        time.Now,              // 测试或服务器在其他时区可设置该属性
		Authenticator:   Authenticator,         // 根据登录信息对用户进行身份验证的回调函数
		LoginResponse:   LoginResponse,         // 完成登录后返回的信息，用户可自定义返回数据，默认返回
		RefreshResponse: auth.RefreshResponse,  // 刷新token后返回的信息，用户可自定义返回数据，默认返回
		Unauthorized:    auth.Unauthorized,     // 处理不进行授权的逻辑
		//IdentityHandler: auth.IdentityHandler,  // 解析并设置用户身份信息
		PayloadFunc:     auth.PayloadFunc,      // 登录期间的回调的函数
		//Authorizator:
	})
	if err != nil {
		glog.Fatal("JWT Error:" + err.Error())
	}
	CustomGfJWTMiddleware = authMiddleware
}

// 自定义中间件
func MiddlewareAuth(r *ghttp.Request) {
	// 使用 gf-jwt中间件
	CustomGfJWTMiddleware.MiddlewareFunc()(r)
	r.Middleware.Next()
}

func main() {

	//七牛云测试
	//localFile := "corn.png"
	//bucket := bucket1
	//key := "testqiniu"
	//// 使用 returnBody 自定义回复格式
	//putPolicy := storage.PutPolicy{
	//	Scope:      bucket,
	//	ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	//}
	//mac := qbox.NewMac(accessKey, secretKey)
	//upToken := putPolicy.UploadToken(mac)
	//cfg := storage.Config{}
	//formUploader := storage.NewFormUploader(&cfg)
	//ret := MyPutRet{}
	//putExtra := storage.PutExtra{
	//	Params: map[string]string{
	//		"x:name": "github logo",
	//	},
	//}
	//err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(ret.Bucket, ret.Key, ret.Fsize, ret.Hash, ret.Name)

	//res, err := g.Redis().Do("HGET","ttt","ioul")
	//if err != nil {
	//	fmt.Println(res)
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(err)
	//	fmt.Println(res)
	//}
	//
	//ress, err := g.Redis().DoVar("HGET","ttt","ioul")
	//if err != nil {
	//	fmt.Println(ress)
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(err)
	//	fmt.Println(ress.IsNil())
	//}

	//count, err := dao.CommentLike.Where("comment_id=3 and user_id=2").FindCount()
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(count)
	//}

	//var c *model.CommentLike
	//err := dao.CommentLike.Where("comment_id=4 and user_id=2").Scan(&c)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(c == nil)
	//}

	s := g.Server()

	// 登录，返回token
	s.BindHandler("POST:/login", CustomGfJWTMiddleware.LoginHandler)

	// 使用中间件
	s.Group("/api", func(group *ghttp.RouterGroup) {
		// 使用中间件
		group.Middleware(MiddlewareAuth)

		// 需要验证token的视图函数
		group.GET("/get_info", func(r *ghttp.Request) {
			r.Response.Write("api get_info...")
		})

		// 刷新token
		group.POST("/refresh_token", CustomGfJWTMiddleware.RefreshHandler)
	})

	s.Run()

}

//身份验证器用于验证登录参数。
//它必须返回用户数据作为用户标识符，它将存储在声明数组中。
//检查错误（e）以确定适当的错误消息。
func Authenticator(r *ghttp.Request) (interface{}, error) {
	data := r.GetMap()
	//if e := gvalid.CheckMap(data, auth.ValidationRules); e != nil {
	//	return "", jwt.ErrFailedAuthentication
	//}
	if data["username"] == "admin" && data["password"] == "111111" {
		return g.Map {
			"username": data["username"],
			"id":       1,
		}, nil
	}
	return nil, jwt.ErrFailedAuthentication
}

// LoginResponse用于定义自定义的登录成功回调函数。
func LoginResponse(r *ghttp.Request, code int, token string, expire time.Time) {
	_ = r.Response.WriteJson(g.Map{
		"id":   1,
		"code":   http.StatusOK,
		"token":  token,
		"expire": expire.Format(time.RFC3339),
	})
	r.ExitAll()
}