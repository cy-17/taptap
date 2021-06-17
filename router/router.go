package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"san616qi/app/api"
	"san616qi/app/service"
)

func init() {

	//服务器初始化
	s := g.Server()
	//中间件注册
	s.Use()

	//分组路由注册
	s.Group("/", func(group *ghttp.RouterGroup) {
		//分组中间件注册
		group.Middleware(
			service.Middleware.Ctx,
			service.Middleware.CORS,)

		group.ALL("/user", api.User)
		group.Group("/", func(group *ghttp.RouterGroup) {
			//group.Middleware(service.Middleware.Auth)

			//查询用户登录状态
			group.ALL("/user/issignedin/:passport", api.User.IsSignedIn)
			//更新用户信息
			group.ALL("/user/updateprofile/:userid", api.User.UpdateProfile)
			//查询用户信息
			group.ALL("/user/queryprofile/:userid", api.User.QueryProfile)
		})
		

	})

}
