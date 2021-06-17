package api

import (
	"github.com/gogf/gf/net/ghttp"
	"san616qi/app/common/consts"
	"san616qi/app/common/response"
	"san616qi/app/service"
)

var Game = new(gameApi)

type gameApi struct{}

func (game *gameApi) MainList(r *ghttp.Request) {

	if err, mainGameList := service.Game.MainList(); err != nil {
		response.JsonExit(r,consts.CurdSelectFailCode,2,consts.CurdSelectFailMsg,"查询失败")
	} else {
		response.JsonExit(r,consts.CurdStatusOkCode,0,consts.CurdStatusOkMsg,mainGameList)
	}

}