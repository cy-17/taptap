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
		//group.ALL("/game", api.Game)
		group.Group("/", func(group *ghttp.RouterGroup) {
			//group.Middleware(service.Middleware.Auth)

			//查询用户登录状态
			group.ALL("/user/issignedin/:passport", api.User.IsSignedIn)
			//更新用户信息
			group.ALL("/user/updateprofile/:userid", api.User.UpdateProfile)
			//查询用户信息
			group.ALL("/user/queryprofile/:userid", api.User.QueryProfile)
		})

		//group.ALL("/game", api.Game)
		group.Group("/", func(group *ghttp.RouterGroup) {

			//卡片式推荐
			group.GET("/game/reclist/:offset", api.Game.RecList)
			//游戏详情
			group.GET("/game/gameprofile/:gameid", api.Game.GameProfile)
			//主游戏列表
			group.GET("/game/mainlist/:classification/:offset", api.Game.GameMainList)

			group.POST("/game/mock", api.Game.GameMock)

		})

		//group.ALL("/game", api.GameComment)
		group.Group("/", func(group *ghttp.RouterGroup) {

			group.POST("/game/comment", api.GameComment.AddComment)
			group.DELETE("/game/comment", api.GameComment.DelComment)
			group.GET("/game/comment/", api.GameComment.SelComment)
			group.GET("/game/detailscore/:gameid", api.GameComment.DetailScore)
			group.GET("/game/childcomment", api.GameComment.SelChildComment)

		})

	})

}
