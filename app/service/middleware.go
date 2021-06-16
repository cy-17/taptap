package service

import "github.com/gogf/gf/net/ghttp"

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
