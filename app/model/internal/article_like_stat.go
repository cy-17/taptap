// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// ArticleLikeStat is the golang structure for table article_like_stat.
type ArticleLikeStat struct {
	ArticleLikeStatusId int64       `orm:"article_like_status_id,primary" json:"articleLikeStatusId"` //
	ArticleId           int64       `orm:"article_id,unique"              json:"articleId"`           //
	LikeCount           int         `orm:"like_count"                     json:"likeCount"`           //
	CreateAt            *gtime.Time `orm:"create_at"                      json:"createAt"`            //
	UpdateAt            *gtime.Time `orm:"update_at"                      json:"updateAt"`            //
}