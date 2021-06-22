package service

import (
	"errors"
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/util/gconv"
	"san616qi/app/dao"
	"san616qi/app/model"
)

// 中间件管理服务
var GameComment = gameCommentService{}

type gameCommentService struct{}

// 用户添加一条评论(针对游戏)
func (gc *gameCommentService) AddComment(r *model.GameAddCommentApiReq) error {
	//先进行参数校验
	if err := gc.CheckComment(r); err != nil {
		return err
	}

	if _, err := dao.GameComment.Save(r); err != nil {
		return errors.New("数据库插入失败")
	}

	return nil
}

//// 用户添加一条评论(针对评论)
//func (gc *gameCommentService) AddChildComment(r *model.GameAddCommentApiReq) error {
//
//	//进行参数校验
//	if err := gc.CheckComment(r);  err != nil {
//		return err
//	}
//
//	if _, err := dao.GameComment.Save(r); err != nil {
//		return errors.New("数据库插入失败")
//	}
//
//	return nil
//
//}

// 用户删除一条评论，只能删除自己的
func (gc *gameCommentService) DelComment(r *model.GameDelCommentApiReq) error {

	if _, err := dao.GameComment.Delete("user_id=? and game_id=? and comment_id=?", r.Userid, r.Gameid, r.CommentId);
		err != nil {
		return errors.New("删除指定评论失败")
	}

	return nil

}

// 查询某游戏的所有评论（分页加载，加载10条主评论，每条主评论加载5条子评论)
func (gc *gameCommentService) SelComment(r *model.GameSelCommentApiReq) (error, *model.GameCommentEntity) {

	//准备返回评论的Entity结构
	var entity = &model.GameCommentEntity{
		GameCommentRep:    make([]*model.GameCommentRep, 0),
		GameCommentStatus: false,
	}
	//评论Entity准备
	var commentEntity = make([]*model.GameCommentRep, 0)
	//准备页数统计
	var parentCommentNum int
	var totalCommentNum int

	//获取评论的用户id和该游戏id
	gameid := r.Gameid
	userid := r.Userid
	//获取分页参数offset 和 limit，并进行处理
	offset := r.Offset
	limit := 10
	childLimit := 5
	if offset == 0 {
		offset = 1
	}

	//先检测两个id是否符合
	if gameid == 0 || userid == 0 {
		return errors.New("游戏id和用户id不可缺失"), nil
	}

	//获取所有主评论的总数
	if temp, err := dao.GameComment.Where("pid=0 and game_id=?", gameid).FindCount(); err != nil {
		return errors.New("数据库查询错误"), nil
	} else {
		parentCommentNum = temp
	}

	//获取所有评论的总数
	if temp, err := dao.GameComment.Where("game_id=?",gameid).FindCount(); err != nil {
		return errors.New("数据库查询错误"), nil
	} else {
		totalCommentNum = temp
	}

	//先获取一级所有评论
	var commentList []*model.GameComment
	if err := dao.GameComment.Where("pid=0 and game_id=?", gameid).Offset((offset - 1) * 10).Limit(limit).Scan(&commentList); err != nil {
		return errors.New("数据库查询错误"), nil
	}

	//把所有获取的评论的子评论也添加
	for _, v := range commentList {

		//准备评论的子评论列表
		var commentChildList []*model.GameComment
		//准备返回的评论Entity结构
		var commentRep = model.GameCommentRep{
			GameParentComment:    &model.ParentComment{},
			GameChildCommentList: make([]*model.ChildComment, 0),
			LikeStatus: 0,
			LikeCount: 0,
		}

		//获取子评论总数
		if childCount, err := dao.GameComment.Where("pid=?",v.CommentId).FindCount(); err != nil {
			return errors.New("数据库查询错误"), nil
		} else {
			commentRep.ChildCount = childCount
		}

		//查询前五条子评论
		if err := dao.GameComment.Where("game_id=? and pid=?", v.GameId, v.CommentId).Offset(0).Limit(childLimit).Scan(&commentChildList);
			err != nil {
			return errors.New("数据库查询错误"), nil
		}

		//组装主评论给类型
		err := gconv.Struct(v, commentRep.GameParentComment)
		if err != nil {
			return errors.New("类型转化错误"), nil
		}
		//获取评论的头像和昵称
		var an *model.UserAvatarAndNickname
		if err := dao.User.Where("user_id=?",v.UserId).Scan(&an); err != nil {
			return errors.New("数据库查询错误"), nil
		}
		//赋值头像和昵称
		commentRep.GameParentComment.Nickname = an.Nickname
		commentRep.GameParentComment.Avatar = an.Avatar

		//处理子评论
		for _, v := range commentChildList {

			var childComment *model.ChildComment
			err := gconv.Struct(v, &childComment)
			if err != nil {
				return errors.New("类型转化错误"), nil
			}

			//获取评论者头像昵称
			var anc *model.UserAvatarAndNickname
			if err := dao.User.Where("user_id=?",v.UserId).Scan(&anc); err != nil {
				return errors.New("数据库查询错误"), nil
			}
			//获取被评论者昵称
			var anr *model.UserAvatarAndNickname
			if err := dao.User.Where("user_id=?",v.RepliedId).Scan(&anr); err != nil {
				return errors.New("数据库查询错误"), nil
			}

			//组装子评论给
			childComment.Avatar = anc.Avatar
			childComment.UserNickname = anc.Nickname
			childComment.RepliedNickname = anr.Nickname

			commentRep.GameChildCommentList = append(commentRep.GameChildCommentList, childComment)
		}

		//添加结构体
		commentEntity = append(commentEntity, &commentRep)
	}

	//包装封装数据
	entity.GameCommentRep = commentEntity
	entity.GameCommentStatus = gc.CheckCommentStatus(userid, gameid)
	entity.ParentCommentNum = parentCommentNum
	entity.TotalCommentNum = totalCommentNum

	return nil, entity

}

// 获取更多子评论（每次加载10条)
func (gc *gameCommentService) SelChildComment(r *model.GameSelChildCommentApiReq) (error, *model.GameChildCommentRep) {

	var entity *model.GameChildCommentRep

	//准备返回子评论的Entity结构
	entity = &model.GameChildCommentRep{
		GameCommentList: make([]*model.ChildComment,0),
	}
	//获取子评论列表准备
	var childCommentEntity = make([]*model.GameComment, 0)

	//获取父评论id
	commentid := r.Comment_id
	//获取offset和limit进行处理,limit限定为5
	offset := r.Offset
	limit := 10
	if offset == 0 || offset == 1 {
		offset = 2
	}

	//先检测id是否缺失
	if commentid == 0 {
		return errors.New("评论id不可缺失"), nil
	}

	if err := dao.GameComment.Where("pid=?",commentid).Offset((offset - 1)*5).Limit(limit).Scan(&childCommentEntity); err != nil {
		return errors.New("数据库查询错误"), nil
	}

	for _, v := range childCommentEntity {

		var childComment *model.ChildComment

		err := gconv.Struct(v, &childComment)
		if err != nil {
			return errors.New("类型转换错误"), nil
		}

		//获取评论者头像昵称
		var anc *model.UserAvatarAndNickname
		if err := dao.User.Where("user_id=?",v.UserId).Scan(&anc); err != nil {
			return errors.New("数据库查询错误"), nil
		}
		//获取被评论者昵称
		var anr *model.UserAvatarAndNickname
		if err := dao.User.Where("user_id=?",v.RepliedId).Scan(&anr); err != nil {
			return errors.New("数据库查询错误"), nil
		}

		//组装子评论数据
		childComment.UserNickname = anc.Nickname
		childComment.Avatar = anc.Avatar
		childComment.RepliedNickname = anr.Nickname

		entity.GameCommentList = append(entity.GameCommentList, childComment)

	}

	return nil, entity

}

// 获取游戏评分统计(游戏详情)
func (gc *gameCommentService) DetailScore(gameid int) (error, *model.GameCommentScoreEntity) {

	//返回Entity准备
	entity := &model.GameCommentScoreEntity{}
	//分数统计准备
	var sum float64

	resultMap := *gmap.New()

	list, err := dao.GameComment.DB.GetAll("select score,count(*) as score_num from game_comment where game_id=? and pid=0 group by score", gameid)
	if err != nil {
		return errors.New("数据库查询错误"), nil
	}

	for _, v := range list {

		//结构体准备
		var convertion = &model.GameCommentScoreRep{}

		if err = gconv.Struct(v, &convertion); err != nil {
			return errors.New("数据转换错误"), nil
		}
		sum += gconv.Float64(convertion.Score)
		resultMap.Set(convertion.Score, convertion.Scorenum)
	}

	//返回总分
	entity.GameCommentScore = resultMap
	entity.TotalScore = (sum / 5.0) * 2.0

	return nil, entity

}

// 检测评论是否合法
func (gc *gameCommentService) CheckComment(r *model.GameAddCommentApiReq) error {

	//先获取传进来的参数

	//评论item特性相关
	userid := r.Userid
	repliedid := r.Repliedid
	gameid := r.Gameid
	pid := r.Pid

	//评论item内容
	content := r.Content
	//获取分数
	score := r.Score

	//先校验评论内容
	if len(content) == 0 {
		return errors.New("评论内容不可为空")
	}

	//再检验评论的特性合法性
	//1.仅对游戏评论 那么要有gameid，userid
	//2.对他人的评论进行评论 那么要有gameid userid repliedid（回复谁）pid（也就是父评论的id）
	if gameid == 0 || userid == 0 {
		return errors.New("评论游戏的gameid或者评论者的userid不可为空")
	} else if (pid == 0 && repliedid != 0) || (pid != 0 && repliedid == 0) {
		return errors.New("如果要评论一条评论，那么被回复者的id和主评论id都要齐全")
	}

	//如果不是对评论进行评论，那么分数一定不能为空
	if pid == 0 && score == 0 {
		return errors.New("对游戏评论分数不可为空")
	}


	return nil

}

// 检测是否为游戏评分过?
func (gc *gameCommentService) CheckCommentStatus(userid, gameid int) bool {

	if count, err := dao.GameComment.Where("user_id=? and game_id=?", userid, gameid).FindCount(); err != nil {
		return false
	} else if count > 0 {
		return true
	} else {
		return false
	}

}
