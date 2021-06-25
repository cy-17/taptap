package router

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"san616qi/app/api"
	"san616qi/app/common/auth"
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
			service.Middleware.Log,
			service.Middleware.Ctx,
			service.Middleware.CORS)

		//用户部分
		group.Group("/", func(group *ghttp.RouterGroup) {

			//用户登录返回token
			group.POST("/user/signin", auth.CustomGfJWTMiddleware.LoginHandler)
			//用户注册
			group.POST("/user/signup", api.User.SignUp)

			//以下请求需要携带token
			group.Middleware(service.Middleware.Jwt)

			//查询用户登录状态
			//group.GET("/user/issignedin/:passport", api.User.IsSignedIn)
			//更新用户信息
			group.PUT("/user/updateprofile", api.User.UpdateProfile)
			//查询用户信息
			group.GET("/user/queryprofile", api.User.QueryProfile)
			//刷新用户token
			group.GET("/user/refreshtoken", auth.CustomGfJWTMiddleware.RefreshHandler)
		})

		//游戏内容部分
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

		//游戏评论部分
		//group.ALL("/game", api.GameComment)
		group.Group("/", func(group *ghttp.RouterGroup) {

			group.GET("/game/comment", api.GameComment.SelComment)
			group.GET("/game/detailscore/:gameid", api.GameComment.DetailScore)
			group.GET("/game/childcomment", api.GameComment.SelChildComment)

			group.Middleware(service.Middleware.Jwt)
			group.PUT("/game/comment", api.GameComment.UpdateComment)
			group.POST("/game/comment", api.GameComment.AddComment)
			group.DELETE("/game/comment", api.GameComment.DelComment)
			group.GET("/game/usercomment/:offset", api.GameComment.SelUserComment)

		})

		//论坛部分
		group.Group("/", func(group *ghttp.RouterGroup) {

			//获取论坛列表
			group.GET("/luntan/list/:offset", api.LunTan.GetForumList)

		})

		//文章模块
		group.Group("/", func(group *ghttp.RouterGroup) {

			//新增文章
			group.POST("/article/addarticle", api.Article.AddArticle)
			//查看文章详情
			group.GET("/article/getarticle/:articleid", api.Article.GetArticle)
			//删除文章
			group.DELETE("/article/delarticle", api.Article.DelArticle)
			//更新文章
			group.POST("/article/updatearticle", api.Article.UpdateArticle)
		})

		//点赞模块
		group.Group("/", func(group *ghttp.RouterGroup) {

			//点赞需要token
			group.Middleware(service.Middleware.Jwt)
			group.POST("/like/comment", api.CommentLike.CommentLike)

		})

	})

}
