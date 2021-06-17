package model

import (
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/net/ghttp"
)

const (
	// 上下文变量存储名
	ContextKey = "ContextKey"
)

// 请求的上下文机构
type Context struct {
	Session *ghttp.Session //统一Session管理对象
	Users   *ContextUsers
}


// 保存已经登录过的用户对象
type ContextUsers struct {
	UsersMap gmap.Map
}
