// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"san616qi/app/dao/internal"
)

// articleDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type articleDao struct {
	*internal.ArticleDao
}

var (
	// Article is globally public accessible object for table article operations.
	Article articleDao
)

func init() {
	Article = articleDao{
		internal.NewArticleDao(),
	}
}

// Fill with you ideas below.
