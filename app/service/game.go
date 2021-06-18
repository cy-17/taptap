package service

import (
	"errors"
	"san616qi/app/dao"
	"san616qi/app/model"
)

// 中间件管理服务
var Game = gameService{}

type gameService struct{}

// 卡片式游戏列表推荐
func (s *gameService) MainList(offset int) (error, *model.GameMainEntity) {

	//返回数据结构准备
	entity := &model.GameMainEntity{
		GameMainListRep: make([]*model.GameMainRep,0),
	}
	//为简化信息list初始化
	gameInfoList := make([]*model.GameInfo,0)

	//offset处理,设定limit
	limit := 7
	if offset == 0 {
		offset = 1
	}

	if err := dao.Game.Offset((offset-1)*7).Limit(limit).Scan(&gameInfoList); err != nil {
		return errors.New("数据库查询失败"), nil
	}

	//组装GameMainRep
	for _,v := range gameInfoList {

		//初始化结构体
		tempGameMainRep := &model.GameMainRep{}

		if err, temp := GameComment.DetailScore(v.GameId); err != nil {
			return errors.New("数据库查询错误"), nil
		} else {
			tempGameMainRep.Score = temp.TotalScore
			tempGameMainRep.GameInfo = v
		}

		//添加数据
		entity.GameMainListRep = append(entity.GameMainListRep, tempGameMainRep)

	}
	
	return nil, entity
}

// 游戏详情获取
func (s *gameService) GameProfile(gameid int) (error, *model.GameProfile) {

	//准备返回的结构体
	gameProfile := &model.GameProfile{}
	var game *model.Game

	if err := dao.Game.Where("game_id=?",gameid).Scan(&game); err != nil {
		return errors.New("数据库查询错误"), nil
	}
	if err, score := GameComment.DetailScore(gameid); err != nil {
		return errors.New("数据库查询错误"), nil
	} else {
		gameProfile.GameCommentScore = score
	}
	gameProfile.Game = game

	return nil, gameProfile
}
