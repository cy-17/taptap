package service

import (
	"errors"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/util/gconv"
	"san616qi/app/dao"
	"san616qi/app/model"
)

var commentLikeStat = new(commentLikeStatService)

type commentLikeStatService struct{}



//获取点赞数量
func(clss *commentLikeStatService) LikeCount(userid,commentid int) (error, int) {

	//点赞数量
	var count int

	//"commentlikecount_commentid%5"---"commentid---value"
	keyCount := "commentlikecount" + gconv.String(commentid%5)
	fieldCount := gconv.String(commentid)

	//再获取点赞数量
	if v, err := g.Redis().DoVar("HGET", keyCount, fieldCount); err != nil {
		return errors.New("redis查询错误"), 0
	} else {
		if !v.IsNil() {
			count = gconv.Int(v)
		} else {
			//redis为空，去mysql查询
			var cls *model.CommentLikeStat
			if err := dao.CommentLikeStat.Where("comment_id=?",commentid).Scan(&cls); err != nil {
				return errors.New("mysql查询错误"), 0
			} else {
				// mysql都没有这个的点赞数量，说明根本没有点赞过
				if cls == nil {
					count = 0
				} else {
					count = cls.LikeCount
				}
			}
		}
	}

	return nil, count

}