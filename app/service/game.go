package service

import (
	"errors"
	"san616qi/app/dao"
	"san616qi/app/model"
)

// 中间件管理服务
var Game = gameService{}

type gameService struct{}

// 主游戏列表推荐
func (s *gameService) MainList() (error, []model.GameMainEntity) {

	var games []model.GameMainEntity

	////先进行初始化
	//var tempList = make([]*model.Game,0)
	//var gameMainListRep = &model.GameMainListRep{
	//	GameMainList: tempList,
	//}

	if err := dao.Game.ScanList(&games,"Games"); err != nil {
		return errors.New("数据库查询失败"), nil
	} else {
		return nil, games
	}
}