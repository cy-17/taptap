package api

import (
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/util/gconv"
	"san616qi/app/common/consts"
	"san616qi/app/common/response"
	"san616qi/app/service"
)

var Game = new(gameApi)

type gameApi struct{}

// 卡片式游戏列表
func (game *gameApi) RecList(r *ghttp.Request) {

	var (
		offset int
	)

	offset = gconv.Int(r.Get("offset"))

	if err, mainGameList := service.Game.RecList(offset); err != nil {
		response.JsonExit(r, consts.CurdSelectFailCode, 2, consts.CurdSelectFailMsg, "查询失败")
	} else {
		response.JsonExit(r, consts.CurdStatusOkCode, 0, consts.CurdStatusOkMsg, mainGameList)
	}

}

// 游戏详情
func (game *gameApi) GameProfile(r *ghttp.Request) {

	var (
		gameid int
	)

	gameid = gconv.Int(r.Get("gameid"))
	if gameid == 0 {
		response.JsonExit(r, consts.RequestParamLostCode, 1, consts.RequestParamLostMsg, "请求的游戏id缺失")
	}
	if err, gameProfile := service.Game.GameProfile(gameid); err != nil {
		response.JsonExit(r, consts.CurdSelectFailCode, 2, consts.CurdSelectFailMsg, "服务器查询出错")
	} else {
		response.JsonExit(r, consts.CurdStatusOkCode, 0, consts.CurdStatusOkMsg, gameProfile)
	}

}
