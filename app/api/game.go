package api

import (
	"fmt"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"math/rand"
	"san616qi/app/common/consts"
	"san616qi/app/common/response"
	"san616qi/app/dao"
	"san616qi/app/model"
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

// 主游戏列表，排行榜展示
func (game *gameApi) GameMainList(r *ghttp.Request) {

	var (
		classification int
		offset int
	)

	classification = gconv.Int(r.Get("classification"))
	offset = gconv.Int(r.Get("offset"))

	if err, mainList := service.Game.GameMainList(classification, offset); err != nil {
		response.JsonExit(r,consts.CurdSelectFailCode,2,consts.CurdSelectFailMsg,"主游戏列表获取失败")
	} else {
		response.JsonExit(r,consts.CurdStatusOkCode,0,consts.CurdStatusOkMsg,mainList)
	}

}

// mock数据
func (game *gameApi) GameMock(r *ghttp.Request) {

	var (
		apiReq []*model.GameMock
	)

	if err := r.Parse(&apiReq); err != nil {
		response.JsonExit(r,consts.RequestParamLostCode,1,consts.RequestParamLostMsg,"12312")
	}

	for _, v := range apiReq {

		game := &model.Game{
			GameName: v.Title,
			Icon: v.Icon,
			CoverImage: v.Banner,
			Shortdesc: v.ShortDesc,
			Introduction: v.Introduction,
			Author: v.PublisherName,
			Tags: "",
			DetailImages: "",
			ReleaseAt: gtime.Now(),
			Classification: rand.Intn(5)+1,
		}
		length := len(v.Tag)

		for i, v1 := range v.Tag {

			if i != length-1 {
				game.Tags = game.Tags + v1 +","
			} else {
				game.Tags = game.Tags + v1
			}

		}

		lk := len(v.ScreenShot)

		for i, v2 := range v.ScreenShot {

			if i != lk - 1 {
				game.DetailImages = game.DetailImages + v2 +","
			} else {
				game.DetailImages = game.DetailImages + v2
			}
		}

		if _,err  := dao.Game.Save(game); err != nil {
			fmt.Println(err)
		}

	}

}