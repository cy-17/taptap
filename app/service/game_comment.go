package service

import (
	"errors"
	"fmt"
	"san616qi/app/dao"
	"san616qi/app/model"
)

// 中间件管理服务
var GameComment = gameCommentService{}

type gameCommentService struct{}

// 用户添加一条评论
func (gc *gameCommentService) AddComment(r *model.GameAddCommentApiReq) error {
	//先进行参数校验
	if err := gc.CheckComment(r); err != nil {
		return err
	}

	if _, err := dao.GameComment.Save(r); err != nil {
		return err
	}

	return nil
}

// 用户删除一条评论，只能删除自己的
func (gc *gameCommentService) DelComment(r *model.GameDelCommentApiReq) error {

	if _, err := dao.GameComment.Delete("user_id=? and game_id=? and comment_id=?", r.Userid, r.Gameid, r.CommentId);
		err != nil {
		return errors.New("删除指定评论失败")
	}

	return nil

}

// 查询某游戏的所有评论
func (gc *gameCommentService) SelComment(r *model.GameSelCommentApiReq) (error, *model.GameCommentEntity) {

	//准备返回评论的Entity结构
	var entity = &model.GameCommentEntity{
		GameCommentRep: make([]*model.GameCommentRep,0),
		GameCommentStatus: false,
	}

	//评论Entity准备
	var commentEntity = make([]*model.GameCommentRep, 0)

	//获取评论的用户id和该游戏id
	gameid := r.Gameid
	userid := r.Userid

	//先检测两个id是否符合
	if gameid == 0 || userid == 0 {
		return errors.New("游戏id和用户id不可缺失"), nil
	}

	//先获取一级所有评论
	var commentList []*model.GameComment
	if err := dao.GameComment.Where("pid=0 and game_id=?", gameid).Scan(&commentList); err != nil {
		return errors.New("数据库查询错误"), nil
	}

	fmt.Println(len(commentList))

	//把所有获取的评论的子评论也添加
	for _, v := range commentList {

		//准备评论的子评论列表
		var commentChildList []*model.GameComment
		//准备返回的评论Entity结构
		var commentRep = model.GameCommentRep{
			GameParentComment: &model.GameComment{},
			GameChildCommentList: make([]*model.GameComment, 0),
		}

		//查询所有子评论
		if err := dao.GameComment.Where("game_id=? and pid=?", v.GameId, v.CommentId).Scan(&commentChildList);
			err != nil {
			return errors.New("数据库查询错误"), nil
		}
		commentRep.GameParentComment = v
		commentRep.GameChildCommentList = commentChildList

		commentEntity = append(commentEntity,&commentRep)
	}

	//包装封装数据
	entity.GameCommentRep = commentEntity
	entity.GameCommentStatus = gc.CheckCommentStatus(userid,gameid)

	return nil, entity

}

func (gc *gameCommentService) CheckComment(r *model.GameAddCommentApiReq) error {

	//先获取传进来的参数

	//评论item特性相关
	userid := r.Userid
	repliedid := r.Repliedid
	gameid := r.Gameid
	pid := r.Pid

	//评论item内容
	score := r.Score
	content := r.Content

	//先校验分数或者评论内容
	if score == 0 || len(content) == 0 {
		return errors.New("评论分数或者内容不可为空")
	}

	//再检验评论的特性合法性
	//1.仅对游戏评论 那么要有gameid，userid
	//2.对他人的评论进行评论 那么要有gameid userid repliedid（回复谁）pid（也就是父评论的id）
	if gameid == 0 || userid == 0 {
		return errors.New("评论游戏的gameid或者评论者的userid不可为空")
	} else if (pid == 0 && repliedid != 0) || (pid != 0 && repliedid == 0) {
		return errors.New("如果要评论一条评论，那么被回复者的id和主评论id都要齐全")
	}

	return nil

}

func (gc *gameCommentService) CheckCommentStatus(userid, gameid int) bool {

	if count, err := dao.GameComment.Where("user_id=? and game_id=?", userid, gameid).FindCount(); err != nil {
		return false
	} else if count > 0 {
		return true
	} else {
		return false
	}

}
