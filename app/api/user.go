package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"san616qi/app/common/consts"
	"san616qi/app/common/response"
	"san616qi/app/model"
	"san616qi/app/service"
)

// 用户api管理对象
var User = new(userApi)

type userApi struct{}

//
//	SignUp
//	@Description:
//	@receiver a *userApi
//	@param r *ghttp.Request
//
func (a *userApi) SignUp(r *ghttp.Request) {

	//需要使用的结构变量
	var (
		//接传递
		apiReq *model.UserApiSignUpReq
		//往下的业务逻辑
		serviceReq *model.UserServiceSignUpReq
	)

	if err := r.ParseForm(&apiReq); err != nil {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, err.Error())
	}
	if err := gconv.Struct(apiReq, &serviceReq); err != nil {
		response.JsonExit(r, consts.ServerOccurredErrorCode, 2, consts.ServerOccurredErrorMsg, err.Error())
	}

	if err := service.User.SignUp(serviceReq); err != nil {
		response.JsonExit(r, consts.CurdCreatFailCode, 2, consts.CurdCreatFailMsg, err.Error())
	} else {
		response.JsonExit(r, consts.CurdStatusOkCode, 0, consts.CurdStatusOkMsg, "注册成功")
	}

}

func (a *userApi) SignIn(r *ghttp.Request) {

	var (
		data *model.UserApiSignInReq
	)
	if err := r.Parse(&data); err != nil {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, err.Error())
	}
	if err, userData := service.User.SignIn(r.Context(), data.Passport, data.Password); err != nil {
		response.JsonExit(r, consts.CurdLoginFailCode, 2, consts.CurdLoginFailMsg, err.Error())
	} else {
		response.JsonExit(r, consts.CurdStatusOkCode, 0, consts.CurdStatusOkMsg,
			userData)
	}

}

func (a *userApi) IsSignedIn(r *ghttp.Request) {

	var (
		passport string
	)
	passport = gconv.String(r.Get("passport"))
	if r.GetParam("Passport"); len(passport) == 0 {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, "passport参数缺失")
	}
	if err := service.User.IsSignedIn(r.Context(), passport); !err {
		response.JsonExit(r, consts.CurdStatusOkCode, 0, consts.CurdStatusOkMsg, false)
	} else {
		response.JsonExit(r, consts.CurdStatusOkCode, 0, consts.CurdStatusOkMsg, true)
	}
}

// 更新个人信息
func (a *userApi) UpdateProfile(r *ghttp.Request) {

	var (
		apiUpdateUser *model.UserApiUpdateProfileReq
		serviceUpdateUser *model.UserServiceUpdateProfileReq
		userId     int
	)

	//token逻辑
	umap := r.GetParam("JWT_PAYLOAD")
	umap = gconv.Map(umap)
	if t, ok := umap.(map[string]interface{}); !ok {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, "请求缺乏用户id")
	} else {
		userId = gconv.Int(t["user_id"])
	}

	if userId == 0 {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, "请求缺乏用户id")
	}
	if err := r.Parse(&apiUpdateUser); err != nil {
		response.JsonExit(r, consts.RequestParamLostCode,1,consts.RequestParamLostMsg,"请求参数有误")
	}
	if err := gconv.Struct(apiUpdateUser, &serviceUpdateUser); err != nil {
		response.JsonExit(r,consts.ServerOccurredErrorCode,2,consts.ServerOccurredErrorMsg,err.Error())
	}

	if err := service.User.UpdateProfile(userId,serviceUpdateUser); err != nil {
		response.JsonExit(r,consts.CurdUpdateFailCode,2,consts.CurdUpdateFailMsg,err.Error())
	} else {
		response.JsonExit(r,consts.CurdStatusOkCode,0,consts.CurdStatusOkMsg,"更新个人信息成功")
	}

}

// 查询个人信息
func (a *userApi) QueryProfile(r *ghttp.Request) {

	var (
		userId int
	)

	//token逻辑
	umap := r.GetParam("JWT_PAYLOAD")
	umap = gconv.Map(umap)
	if t, ok := umap.(map[string]interface{}); !ok {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, "请求缺乏用户id")
	} else {
		userId = gconv.Int(t["user_id"])
	}

	if userId == 0 {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, "请求缺乏用户id")
	}
	if err, user := service.User.QueryProfile(userId); err != nil {
		response.JsonExit(r,consts.CurdSelectFailCode,2,consts.CurdSelectFailMsg,err.Error())
	} else {
		response.JsonExit(r,consts.CurdStatusOkCode,0,consts.CurdStatusOkMsg,user)
	}

}

