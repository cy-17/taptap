// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package model

import (
	"github.com/gogf/gf/container/gmap"
	"github.com/gogf/gf/os/gtime"
	"san616qi/app/model/internal"
)

// GameComment is the golang structure for table game_comment.
type GameComment internal.GameComment

// Fill with you ideas below.

//以下部分是Request传过来的BO
type GameAddCommentApiReq struct {
	Userid    int `v:"required#评论者用户id不可为空"`
	Gameid    int `v:"required#评论的游戏id不可为空"`
	Repliedid int
	Pid       int
	Content   string `v:"required#评论内容不可为空"`
	CreateAt  *gtime.Time
	Score     int
	Commentid int
}

//删除自己的评论
type GameDelCommentApiReq struct {
	CommentId int `v:"required#评论的id不可为空"`
	Userid    int `v:"required#删除评论时用户id不可为空"`
	Gameid    int `v:"required#评论的游戏id不可为空"`
}

//查询一个游戏下的所有评论
type GameSelCommentApiReq struct {
	Gameid int `v:"required#获取评论的游戏id不可为空"`
	Userid int `v:"required#获取评论时候用户id不可为空"`
	Offset int `v:"between:0,100000#offset异常，要在0-100之间"`
	//Limit  int
}

//查询一个评论的子评论
type GameSelChildCommentApiReq struct {
	Comment_id int `v:"required#获取子评论时，评论id不可为空"`
	Offset     int `v:"between:0,100000#offset异常，要在0-100之间"`
}

//以下部分是Reponse回去的VO

//带有用户头像，昵称的父评论结构体
type ParentComment struct {
	CommentId int64
	UserId    int64
	GameId    int64
	RepliedId int64
	Pid       int64
	Content   string
	CreateAt  *gtime.Time
	Score     int    //评论分数
	Nickname  string //用户昵称
	Avatar    string //用户头像
}

//带有用户头像，昵称，回复昵称的子评论结构体，没有分数
type ChildComment struct {
	CommentId       int64
	UserId          int64
	GameId          int64
	RepliedId       int64
	Pid             int64
	Content         string
	CreateAt        *gtime.Time
	UserNickname    string //用户昵称
	RepliedNickname string //回复谁的用户昵称
	Avatar          string //用户头像
}

//用户自己的评论
type GameUserCommentRep struct {
	GameParentComment *ParentComment
	LikeCount         int
	ChildCount        int
}

//评论列表
type GameCommentRep struct {
	GameParentComment    *ParentComment
	GameChildCommentList []*ChildComment
	LikeStatus           int
	LikeCount            int
	ChildCount           int
}

//评论列表entity
type GameCommentEntity struct {
	GameCommentRep    []*GameCommentRep
	GameCommentStatus bool
	ParentCommentNum  int
	TotalCommentNum   int
}

//子评论列表
type GameChildCommentRep struct {
	GameCommentList []*ChildComment
}

// 子评论列表entity
type GameChildCommentEntity struct {
	GameChildCommentRep *GameChildCommentRep
}

//评分包装VO
type GameCommentScoreRep struct {
	Score    int
	Scorenum int
}

//评分包装EntityVO
type GameCommentScoreEntity struct {
	GameCommentScore gmap.Map
	TotalScore       float64
}

//以下部分是参与Service的业务数据
//type GameCommentServiceReq struct {}
