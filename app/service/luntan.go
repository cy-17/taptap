package service

import (
	"errors"
	"san616qi/app/dao"
	"san616qi/app/model"
	"strings"
)

// 中间件管理服务
var LunTan = luntanService{}

type luntanService struct{}

// 获取论坛列表
func (lt *luntanService) GetForumList(offset int) (error, *model.LuntanEntity) {

	//准备返回的entity数据
	entity := &model.LuntanEntity{
		LuntanRepList: make([]*model.LuntanRep, 0),
	}

	//准备相应数据结构体
	var result []*model.Luntan

	//获取offset，并处理，limit为15
	if offset == 0 {
		offset = 1
	}
	limit := 15

	//获取数据
	if err := dao.Luntan.Offset((offset-1)*limit).Limit(limit).Scan(&result); err != nil {
		return errors.New("数据库查询错误"), nil
	}

	//处理数据，准备返回
	for _, v := range result {

		var rep = &model.LuntanRep{
			LuntanId: v.LuntanId,
			LuntanName: v.LuntanName,
			Shortdesc: v.Shortdesc,
			Icon: v.Icon,
			GameId: v.GameId,
		}
		rep.Tags = strings.Split(v.Tags,",")

		entity.LuntanRepList = append(entity.LuntanRepList, rep)

	}

	return nil, entity

}
