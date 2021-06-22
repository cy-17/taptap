package api

import (
	"github.com/gogf/gf/net/ghttp"
	"san616qi/app/common/consts"
	"san616qi/app/common/response"
	"san616qi/app/model"
	"san616qi/app/service"
)

var CommentLike = new(commentLikeApi)

type commentLikeApi struct{}

//对评论进行点赞
func (cla *commentLikeApi) CommentLike(r *ghttp.Request) {

	var (
		apiReq *model.CommentLikeApiReq
	)

	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, err.Error())
	}
	if err := service.CommentLike.CommentLike(apiReq); err != nil {
		response.JsonExit(r, consts.RedisCurdCreatFailCode,2,consts.RedisCurdCreatFailMsg,err.Error())
	} else {
		response.JsonExit(r, consts.CurdStatusOkCode,0,consts.CurdStatusOkMsg,"点赞成功")
	}

}
