package service

import (
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"san616qi/app/model"
	"time"
)

var Middleware = middlewareService{}

type middlewareService struct{}

//
//	CORS
//	@Description:
//	@receiver s *middlewareService
//	@param r *ghttp.Request
//
func (s *middlewareService) CORS(r *ghttp.Request) {
	r.Response.CORSDefault()
	r.Middleware.Next()
}

func (s *middlewareService) Ctx(r *ghttp.Request) {

	// 初始化，务必最开始执行，作为一开始的中间件
	customCtx := &model.Context{
		Session: r.Session,
	}
	Context.Init(r,customCtx)
	if users := Session.GetUser(r.Context()); users != nil {
		customCtx.Users = users
	} else {
		customCtx.Users = &model.ContextUsers{
			UsersMap: gmap.Map{},
		}
	}

	// 执行下一个中间件
	r.Middleware.Next()
}

func (s *middlewareService) Log(r *ghttp.Request) {

	bT := time.Now()
	r.Middleware.Next()
	g.Log().Println(r.Response.Status, r.URL.Path, time.Since(bT))

}