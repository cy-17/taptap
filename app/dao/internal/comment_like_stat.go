// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// CommentLikeStatDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type CommentLikeStatDao struct {
	gmvc.M                         // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB                 // DB is the raw underlying database management object.
	Table   string                 // Table is the table name of the DAO.
	Columns commentLikeStatColumns // Columns contains all the columns of Table that for convenient usage.
}

// CommentLikeStatColumns defines and stores column names for table comment_like_stat.
type commentLikeStatColumns struct {
	CommentLikeStat string //
	CommentId       string //
	LikeCount       string //
	CreateAt        string //
	UpdateAt        string //
}

func NewCommentLikeStatDao() *CommentLikeStatDao {
	return &CommentLikeStatDao{
		M:     g.DB("default").Model("comment_like_stat").Safe(),
		DB:    g.DB("default"),
		Table: "comment_like_stat",
		Columns: commentLikeStatColumns{
			CommentLikeStat: "comment_like_stat",
			CommentId:       "comment_id",
			LikeCount:       "like_count",
			CreateAt:        "create_at",
			UpdateAt:        "update_at",
		},
	}
}
