// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// ArticleLike is the golang structure for table article_like.
type ArticleLike struct {
	ArticleLikeId     int64       `orm:"article_like_id,primary" json:"articleLikeId"`     //
	UserId            int64       `orm:"user_id"                 json:"userId"`            //
	ArticleId         int64       `orm:"article_id"              json:"articleId"`         //
	ArticleLikeStatus int         `orm:"article_like_status"     json:"articleLikeStatus"` //
	ArticleLikeTime   *gtime.Time `orm:"article_like_time"       json:"articleLikeTime"`   //
	CreateAt          *gtime.Time `orm:"create_at"               json:"createAt"`          //
	UpdateAt          *gtime.Time `orm:"update_at"               json:"updateAt"`          //
}
