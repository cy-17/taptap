package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"san616qi/app/common/consts"
	"san616qi/app/common/response"
	"san616qi/app/service"
)

var LunTan = new(luntanApi)

type luntanApi struct{}

//获取论坛的话题列表
func (lt *luntanApi) GetForumList(r *ghttp.Request) {

	var (
		offset int
	)

	offset = gconv.Int(r.Get("offset"))

	if err, forumList := service.LunTan.GetForumList(offset); err != nil {
		response.JsonExit(r,consts.CurdSelectFailCode,2,consts.CurdSelectFailMsg,"查询错误")
	} else {
		response.JsonExit(r,consts.CurdStatusOkCode,0,consts.CurdStatusOkMsg,forumList)
	}


}