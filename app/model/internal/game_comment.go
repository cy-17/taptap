// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// GameComment is the golang structure for table game_comment.
type GameComment struct {
	CommentId int64       `orm:"comment_id,primary" json:"commentId"` //
	UserId    int64       `orm:"user_id"            json:"userId"`    //
	GameId    int64       `orm:"game_id"            json:"gameId"`    //
	RepliedId int64       `orm:"replied_id"         json:"repliedId"` //
	Pid       int64       `orm:"pid"                json:"pid"`       //
	Content   string      `orm:"content"            json:"content"`   //
	CreateAt  *gtime.Time `orm:"create_at"          json:"createAt"`  //
	Score     int         `orm:"score"              json:"score"`     //
}
