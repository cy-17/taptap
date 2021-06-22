package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"san616qi/app/common/consts"
	"san616qi/app/common/response"
	"san616qi/app/model"
	"san616qi/app/service"
)

var Article = new(articleApi)

type articleApi struct{}


//新增一篇文章
func (a *articleApi) AddArticle(r *ghttp.Request) {

	var (
		apiReq *model.ArticleAddApiReq
	)

	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r,consts.RequestParamLostCode,1,consts.RequestParamLostMsg,"请求参数缺失")
	}
	if err := service.Article.AddArticle(apiReq); err != nil {
		response.JsonExit(r,consts.CurdCreatFailCode,2,consts.CurdCreatFailMsg,"新增文章失败")
	} else {
		response.JsonExit(r,consts.CurdStatusOkCode,0,consts.CurdStatusOkMsg,"新增文章成功")
	}

}

//更新文章
func (a *articleApi) UpdateArticle(r *ghttp.Request) {

	var (
		apiReq *model.ArticleUpdateApiReq
	)

	if err:= r.Parse(&apiReq); err != nil {
		response.JsonExit(r,consts.RequestParamLostCode,1,consts.RequestParamLostMsg,err.Error())
	}

	if err := service.Article.UpdateArticle(apiReq); err != nil {
		response.JsonExit(r,consts.CurdUpdateFailCode,2,consts.CurdUpdateFailMsg,"更新失败")
	} else {
		response.JsonExit(r,consts.CurdStatusOkCode,0,consts.CurdStatusOkMsg,"文章更新成功")
	}


}

//删除文章
func (a *articleApi) DelArticle(r *ghttp.Request) {

	var (
		apiReq *model.ArticleDelApiReq
	)

	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r,consts.RequestParamLostCode,1,consts.RequestParamLostMsg,err.Error())
	}
	if err := service.Article.DelArticle(apiReq); err != nil {
		response.JsonExit(r,consts.CurdDeleteFailCode,2,consts.CurdDeleteFailMsg,err.Error())
	} else {
		response.JsonExit(r,consts.CurdStatusOkCode,0,consts.CurdStatusOkMsg,"文章删除成功")
	}
}

//查看文章详情
func (a *articleApi) GetArticle(r *ghttp.Request) {

	var (
		articleid int
	)
	articleid = gconv.Int(r.Get("articleid"))
	if articleid == 0 {
		response.JsonExit(r,consts.RequestParamLostCode,1,consts.RequestParamLostMsg,"文章id缺失")
	}
	if err, article := service.Article.GetArticle(articleid); err != nil {
		response.JsonExit(r,consts.CurdSelectFailCode,2,consts.CurdSelectFailMsg,err.Error())
	} else {
		response.JsonExit(r,consts.CurdStatusOkCode,0,consts.CurdStatusOkMsg,article)
	}

}