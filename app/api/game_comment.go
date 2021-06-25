package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"san616qi/app/common/consts"
	"san616qi/app/common/response"
	"san616qi/app/model"
	"san616qi/app/service"
)

// 游戏的评论评价管理对象
var GameComment = new(gameCommentApi)

type gameCommentApi struct{}

// 为指定游戏增加一个评论（考虑有无pid，repliedId，commentId,userId）
// token必选
func (gc *gameCommentApi) AddComment(r *ghttp.Request) {

	var (
		//接收传递过来的发起评论的请求数据
		apiReq = &model.GameAddCommentApiReq{}
	)

	//token逻辑
	umap := r.GetParam("JWT_PAYLOAD")
	umap = gconv.Map(umap)
	if t, ok := umap.(map[string]interface{}); !ok {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, "请求缺乏用户id")
	} else {
		apiReq.Userid = gconv.Int(t["user_id"])
	}

	if apiReq.Userid == 0 {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, "请求缺乏用户id")
	}

	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, err.Error())
	}
	//把创建时间赋值处理
	apiReq.CreateAt = gtime.Now()

	if err := service.GameComment.AddComment(apiReq); err != nil {
		response.JsonExit(r, consts.CurdCreatFailCode, 2, consts.CurdCreatFailMsg, err.Error())
	} else {
		response.JsonExit(r, consts.CurdStatusOkCode, 0, consts.CurdStatusOkMsg, "评论成功，请继续踊跃发言")
	}

}

//更新自己的评论
// token必选
func (gc *gameCommentApi) UpdateComment(r *ghttp.Request) {

	var (
		apiReq = &model.GameAddCommentApiReq{}
	)

	//token逻辑
	umap := r.GetParam("JWT_PAYLOAD")
	umap = gconv.Map(umap)
	if t, ok := umap.(map[string]interface{}); !ok {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, "请求缺乏用户id")
	} else {
		apiReq.Userid = gconv.Int(t["user_id"])
	}

	if apiReq.Userid == 0 {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, "请求缺乏用户id")
	}

	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, consts.RequestParamLostCode,1,consts.RequestParamLostMsg,err.Error())
	}
	//更新创建时间
	apiReq.CreateAt = gtime.Now()

	if err := service.GameComment.UpdateComment(apiReq); err != nil {
		response.JsonExit(r, consts.CurdUpdateFailCode, 2, consts.CurdUpdateFailMsg, err.Error())
	} else {
		response.JsonExit(r, consts.CurdStatusOkCode, 0, consts.CurdStatusOkMsg, "评论更新")
	}
}

// 删除一条指定的评论
// token必选
func (gc *gameCommentApi) DelComment(r *ghttp.Request) {

	var (
		//接收传递过来的删除评论数据
		apiReq = &model.GameDelCommentApiReq{}
	)

	//token逻辑
	umap := r.GetParam("JWT_PAYLOAD")
	umap = gconv.Map(umap)
	if t, ok := umap.(map[string]interface{}); !ok {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, "请求缺乏用户id")
	} else {
		apiReq.Userid = gconv.Int(t["user_id"])
	}

	if apiReq.Userid == 0 {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, "请求缺乏用户id")
	}

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
// token
func (gc *gameCommentApi) SelComment(r *ghttp.Request) {

	//准备结构体接收参数
	var (
		//接收传递过来的查询游戏评论数据
		apiReq = &model.GameSelCommentApiReq{}
	)

	//token逻辑
	umap := r.GetParam("JWT_PAYLOAD")
	umap = gconv.Map(umap)
	if t, ok := umap.(map[string]interface{}); !ok {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, "请求缺乏用户id")
	} else {
		apiReq.Userid = gconv.Int(t["user_id"])
	}

	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, err.Error())
	}

	if err, commentList := service.GameComment.SelComment(apiReq); err != nil {
		response.JsonExit(r, consts.CurdSelectFailCode, 2, consts.CurdSelectFailMsg, err.Error())
	} else {
		response.JsonExit(r, consts.CurdStatusOkCode, 0, consts.CurdStatusOkMsg, commentList)
	}
}

// 查询评论下的子评论
// token
func (gc *gameCommentApi) SelChildComment(r *ghttp.Request) {

	//准备结构体接收参数
	var (
		//接收传递过来查询子评论的数据
		apiReq = &model.GameSelChildCommentApiReq{}
	)

	//token逻辑
	umap := r.GetParam("JWT_PAYLOAD")
	umap = gconv.Map(umap)
	if t, ok := umap.(map[string]interface{}); !ok {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, "请求缺乏用户id")
	} else {
		apiReq.Userid= gconv.Int(t["user_id"])
	}

	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, err.Error())
	}

	if err, childCommentList := service.GameComment.SelChildComment(apiReq); err != nil {
		response.JsonExit(r, consts.CurdSelectFailCode, 2, consts.CurdSelectFailMsg, err.Error())
	} else {
		response.JsonExit(r, consts.CurdStatusOkCode, 0, consts.CurdStatusOkMsg, childCommentList)
	}

}


// 查询游戏的评分统计(游戏详情用）
func (gc *gameCommentApi) DetailScore(r *ghttp.Request) {

	var (
		//接收游戏id参数
		gameid int
	)

	gameid = gconv.Int(r.Get("gameid"))

	if err,resultMap := service.GameComment.DetailScore(gameid); err != nil {
		response.JsonExit(r, consts.CurdSelectFailCode, 2, consts.CurdSelectFailMsg, err.Error())
	} else {
		response.JsonExit(r, consts.CurdStatusOkCode, 0, consts.CurdStatusOkMsg, resultMap)
	}

}