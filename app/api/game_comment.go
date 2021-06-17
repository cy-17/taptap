package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"san616qi/app/common/consts"
	"san616qi/app/common/response"
	"san616qi/app/model"
	"san616qi/app/service"
)

// 游戏的评论评价管理对象
var GameComment = new(gameCommentApi)

type gameCommentApi struct{}

// 为指定游戏增加一个评论（考虑有无pid，repliedId，commentId,userId）
func (gc *gameCommentApi) AddComment(r *ghttp.Request) {

	var (
		//接收传递过来的发起评论的请求数据
		apiReq *model.GameAddCommentApiReq
	)

	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, err.Error())
	}
	//把创建时间赋值处理
	apiReq.CreateAt = gtime.Now()

	if err := service.GameComment.AddComment(apiReq); err != nil {
		response.JsonExit(r, consts.CurdCreatFailCode, 2, consts.CurdCreatFailMsg, err.Error())
	} else {
		response.JsonExit(r, consts.CurdStatusOkCode, 0, consts.CurdStatusOkMsg, "asda")
	}

}

// 删除一条指定的评论
func (gc *gameCommentApi) DelComment(r *ghttp.Request) {

	var (
		//接收传递过来的删除评论数据
		apiReq *model.GameDelCommentApiReq
	)

	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, err.Error())
	}

	if err := service.GameComment.DelComment(apiReq); err != nil {
		response.JsonExit(r, consts.CurdDeleteFailCode, 2, consts.CurdDeleteFailMsg, err.Error())
	} else {
		response.JsonExit(r, consts.CurdStatusOkCode, 0, consts.CurdStatusOkMsg, "删除成功")
	}

}

// 查询某游戏所有评论
func (gc *gameCommentApi) SelComment(r *ghttp.Request) {
	var (
		//接收传递过来的查询游戏评论数据
		apiReq *model.GameSelCommentApiReq
	)

	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, err.Error())
	}

	if err,commentList := service.GameComment.SelComment(apiReq); err != nil {
		response.JsonExit(r, consts.CurdSelectFailCode, 2, consts.CurdSelectFailMsg, err.Error())
	} else {
		response.JsonExit(r, consts.CurdStatusOkCode, 0, consts.CurdStatusOkMsg, commentList)
	}
}
