package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
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
		group.Middleware(service.Middleware.CORS)

	})

}
