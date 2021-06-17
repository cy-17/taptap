package service

import (
	"context"
	"github.com/gogf/gf/net/ghttp"
	"san616qi/app/model"
)

var Context = contextService{}

type contextService struct{}

// 初始化上下文对象指针到上下文对象中，以便后续的请求流程中可以修改。
func (s *contextService) Init(r *ghttp.Request, customCtx *model.Context) {
	r.SetCtxVar(model.ContextKey, customCtx)
}

// 获得上下文变量，如果没有设置，那么返回nil
func (s *contextService) Get(ctx context.Context) *model.Context {
	value := ctx.Value(model.ContextKey)
	if value == nil {
		return nil
	}
	if localCtx, ok := value.(*model.Context); ok {
		return localCtx
	}
	return nil
}

// 将上下文信息设置到上下文请求中
func (s *contextService) SetUsers(ctx context.Context, ctxUsers *model.ContextUsers) {
	s.Get(ctx).Users.UsersMap = ctxUsers.UsersMap
}