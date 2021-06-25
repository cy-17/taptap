package service

import (
	"errors"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"san616qi/app/dao"
	"san616qi/app/model"
	"strings"
)

// 中间件管理服务
var Game = gameService{}

type gameService struct{}

// 卡片式游戏列表推荐
func (s *gameService) RecList(offset int) (error, *model.GameRecEntity) {

	//返回数据结构准备
	entity := &model.GameRecEntity{
		GameRecListRep: make([]*model.GameRecRep, 0),
	}
	//为简化信息list初始化
	gameInfoList := make([]*model.GameInfo, 0)

	//offset处理,设定limit
	limit := 7
	if offset == 0 {
		offset = 1
	}

	if err := dao.Game.Offset((offset - 1) * 7).Limit(limit).Scan(&gameInfoList); err != nil {
		return errors.New("数据库查询失败"), nil
	}

	//组装GameRecRep
	for _, v := range gameInfoList {

		//初始化结构体
		tempGameRecRep := &model.GameRecRep{}

		if err, temp := GameComment.DetailScore(v.GameId); err != nil {
			return errors.New("数据库查询错误"), nil
		} else {
			tempGameRecRep.Score = temp.TotalScore
			tempGameRecRep.GameInfo = v
		}

		//添加数据
		entity.GameRecListRep = append(entity.GameRecListRep, tempGameRecRep)

	}

	return nil, entity
}

// 游戏详情获取
// redis逻辑
func (s *gameService) GameProfile(gameid int) (error, *model.GameProfile) {

	//准备返回的结构体
	gameProfile := &model.GameProfile{
		Game: &model.GameWithSlice{
			Tags:         make([]string, 0),
			DetailImages: make([]string, 0),
		},
	}
	//用于查询游戏的结构体
	game := &model.Game{}

	//现在redis里面寻找
	profileKey := "gameprofile_" + gconv.String(gameid%5)
	profileField := gconv.String(gameid)

	if v, err := g.Redis().DoVar("HGET", profileKey, profileField); err != nil {
		return errors.New("redis查询错误"), nil
	} else {
		//redis内存在
		if !v.IsNil() {

			err := gconv.Struct(v, &game)
			if err != nil {
				return errors.New("类型转换错误"), nil
			}
			//把该游戏的详情延时
			if _, err := g.Redis().Do("EXPIRE", profileKey, 600); err != nil {
				return errors.New("设置延时错误"), nil
			}
		}
		//else {
		//	//redis内不存在，要在mysql里面查询
		//	if err := dao.Game.Where("game_id=?", gameid).Scan(&game); err != nil {
		//		return errors.New("数据库查询错误"), nil
		//	}
		//	//如果mysql存在，那就写入redis
		//	if game != nil {
		//		//写入redis
		//		if _, err := g.Redis().Do("HSET", profileKey, profileField, game); err != nil {
		//			return errors.New("redis设置错误"), nil
		//		}
		//		//把该游戏的详情延时
		//		if _, err := g.Redis().Do("EXPIRE", profileKey, 1200); err != nil {
		//			return errors.New("设置延时错误"), nil
		//		}
		//	}
		//}
	}

	if err := dao.Game.Where("game_id=?", gameid).Scan(&game); err != nil {
		return errors.New("数据库查询错误"), nil
	}
	if err, score := GameComment.DetailScore(gameid); err != nil {
		return errors.New("数据库查询错误"), nil
	} else {
		gameProfile.GameCommentScore = score
	}

	gameProfile.Game.GameId = game.GameId
	gameProfile.Game.GameName = game.GameName
	gameProfile.Game.Icon = game.Icon
	gameProfile.Game.Author = game.Author
	gameProfile.Game.Classification = game.Classification
	gameProfile.Game.ReleaseAt = game.ReleaseAt
	gameProfile.Game.CoverImage = game.CoverImage
	gameProfile.Game.Shortdesc = game.Shortdesc
	gameProfile.Game.Introduction = game.Introduction
	gameProfile.Game.Tags = strings.Split(game.Tags, ",")
	gameProfile.Game.DetailImages = strings.Split(game.DetailImages, ",")

	//设置游戏详情到redis中
	if _, err := g.Redis().Do("HSET",profileKey,profileField,gameProfile); err != nil {
		return errors.New("redis设置错误"), nil
	}

	//把该游戏的详情延时
	if _, err := g.Redis().Do("EXPIRE", profileKey, 600); err != nil {
		return errors.New("设置延时错误"), nil
	}

	return nil, gameProfile
}

// 主游戏列表获取
func (s *gameService) GameMainList(classification, offset int) (error, *model.GameMainEntity) {

	//准备返回的entity
	entity := &model.GameMainEntity{
		GameMainList: make([]*model.GameMainRep, 0),
	}

	//处理offset,每次加载10个
	if offset == 0 {
		offset = 1
	}

	//处理分类，如果是0，那么就是主列表，从所有游戏中取，如果是1~5，那么就是分类的列表
	if classification == 0 {

		var gameMainInfoList []*model.GameMainInfo

		if err := dao.Game.Offset((offset - 1) * 10).Limit(10).Scan(&gameMainInfoList); err != nil {
			return errors.New("数据库查询错误"), nil
		}

		for _, v := range gameMainInfoList {

			gameMainRep := &model.GameMainRep{
				GameMainInfo: v,
			}

			if err, sc := GameComment.DetailScore(v.GameId); err != nil {
				return errors.New("获取评分失败"), nil
			} else {
				gameMainRep.Score = sc.TotalScore
			}

			entity.GameMainList = append(entity.GameMainList, gameMainRep)

		}

		return nil, entity

	} else {

		var gameMainInfoList []*model.GameMainInfo

		if err := dao.Game.Where("classification=?", classification).
			Offset((offset - 1) * 10).Limit(10).Scan(&gameMainInfoList); err != nil {
			return errors.New("数据库查询错误"), nil
		}

		for _, v := range gameMainInfoList {

			gameMainRep := &model.GameMainRep{
				GameMainInfo: v,
			}

			if err, sc := GameComment.DetailScore(v.GameId); err != nil {
				return errors.New("获取评分失败"), nil
			} else {
				gameMainRep.Score = sc.TotalScore
			}

			entity.GameMainList = append(entity.GameMainList, gameMainRep)

		}

		return nil, entity
	}
}
