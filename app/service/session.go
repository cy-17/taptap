package service

import (
	"context"
	"san616qi/app/model"
)

var Session = sessionService{}

type sessionService struct{}

const (
	// 用户信息存放在Session中的key
	sessionKeyUser = "SessionKeyUser"
)

// 设置登陆成功的用户的Session
func (s *sessionService) SetUser(ctx context.Context, users *model.ContextUsers) error {
	return Context.Get(ctx).Session.Set(sessionKeyUser,users)
}

// 获取Session中缓存的用户登录信息
func (s *sessionService) GetUser(ctx context.Context) *model.ContextUsers {

	customCtx := Context.Get(ctx)
	if customCtx != nil {
		if v := customCtx.Session.GetVar(sessionKeyUser); !v.IsNil() {
			var users *model.ContextUsers
			_ = v.Struct(&users)
			return users
		}
	}
	return nil

}