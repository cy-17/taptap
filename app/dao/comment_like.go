// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"san616qi/app/dao/internal"
)

// commentLikeDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type commentLikeDao struct {
	*internal.CommentLikeDao
}

var (
	// CommentLike is globally public accessible object for table comment_like operations.
	CommentLike commentLikeDao
)

func init() {
	CommentLike = commentLikeDao{
		internal.NewCommentLikeDao(),
	}
}

// Fill with you ideas below.
