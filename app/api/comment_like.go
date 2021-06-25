package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
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
		apiReq = &model.CommentLikeApiReq{}
	)

	//token逻辑
	umap := r.GetParam("JWT_PAYLOAD")
	umap = gconv.Map(umap)
	if t, ok := umap.(map[string]interface{}); !ok {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, "请求缺乏用户id")
	} else {
		apiReq.UserId = gconv.Int(t["user_id"])
	}

	if apiReq.UserId == 0 {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, "请求缺乏用户id")
	}

	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, err.Error())
	}
	if err := service.CommentLike.CommentLike(apiReq); err != nil {
		response.JsonExit(r, consts.RedisCurdCreatFailCode,2,consts.RedisCurdCreatFailMsg,err.Error())
	} else {
		response.JsonExit(r, consts.CurdStatusOkCode,0,consts.CurdStatusOkMsg,"点赞成功")
	}

}
