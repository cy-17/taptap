package service

import (
	"errors"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/util/gconv"
	"san616qi/app/dao"
	"san616qi/app/model"
)

var CommentLike = new(commentLikeService)

type commentLikeService struct{}

//点赞评论
func (cls *commentLikeService) CommentLike(r *model.CommentLikeApiReq) error {

	//先获取用户id和评论id
	userid := r.UserId
	commentid := r.CommentId
	status := r.Status

	//在redis内进行操作
	//存储数据格式 key field value key是commentid % 5 | field是 commentid::userid | value为点赞状态
	//拼装点赞状态field
	field := gconv.String(commentid) + "::" + gconv.String(userid)
	//拼装点赞状态key
	key := "commentlike_" + gconv.String(commentid%5)
	//设定点赞状态
	value := status

	//点赞逻辑处理
	//1.要处理用户点赞
	res, err := g.Redis().DoVar("HGET", key, field)
	if err != nil {
		return errors.New("redis查询出错")
	} else {
		//redis查出来的结果是这个field有东西,那么直接更新就可以了
		if !res.IsNil() {
			//更新redis对应的key-field-value
			if _, err := g.Redis().Do("HSET", key, field, value); err != nil {
				//redis更新失败
				return errors.New("redis更新出错")
			} else {
				//更新成功，要设置过期时间
				if _, err := g.Redis().Do("EXPIRE", key, 900); err != nil {
					//redis更新失败
					return errors.New("redis更新出错")
				}
			}
		} else {
			//redis查出来的结果是这个field没有东西
			//那么就要到mysql里面查，准备一个数据结构接收查询
			var c *model.CommentLike
			err := dao.CommentLike.Where("comment_id=? and user_id=?", commentid, userid).Scan(&c)
			if err != nil {
				return errors.New("mysql查询点赞信息失败")
			} else {

				//写到redis中，再次操作提高效率
				if _, err := g.Redis().Do("HSET", key, field, value); err != nil {
					//redis更新失败
					return errors.New("redis更新出错")
				}
				//同时延长过期时间到下一次flush
				if _, err := g.Redis().Do("EXPIRE", key, 900); err != nil {
					return errors.New("redis更新过期时间出错")
				}

				//如果说mysql有这个数据，那么就把他写回redis中，可能会再次被操作,不用再次操作db
				if c == nil {
					//组织数据
					c = &model.CommentLike{
						CommentLikeTime: gtime.Now(),
						CreateAt:        gtime.Now(),
						UpdateAt:        gtime.Now(),
						UserId:          int64(userid),
						CommentId:       int64(commentid),
						CommentLikeStat: status,
					}
					// 插入形成mysql记录
					if _, err := dao.CommentLike.Save(c); err != nil {
						return errors.New("mysql插入出错")
					}
				}
			}
		}
	}

	//拼装数量key
	keyCount := "commentlikecount_" + gconv.String(commentid%5)
	//拼装数量field
	fieldCount := gconv.String(commentid)

	//2.要处理点赞数量
	//value 为1，那么就进行点赞，数量加一，value为0，那么就取消点赞，点赞的数量要减一
	var incr int
	if value == 1 {
		incr = 1
	} else if value == 0 {
		incr = -1
	} else {
		return errors.New("点赞状态错误")
	}

	if count, err := g.Redis().DoVar("HGET", keyCount, fieldCount); err != nil {
		return errors.New("redis查询出错")
	} else {
		//如果count查出来不是nil，说明redis有这个
		if !count.IsNil() {
			//更新redis中的结果
			if _, err := g.Redis().Do("HINCRBY", keyCount, fieldCount, incr); err != nil {
				return errors.New("redis自增出错")
			}
			//延长时间
			if _, err := g.Redis().Do("EXPIRE", keyCount, 900); err != nil {
				return errors.New("redis设置延时出错")
			}
		} else {
			//查出来count为nil，说明redis没有这个,那么就在mysql里面查
			var likeStat *model.CommentLikeStat

			if err := dao.CommentLikeStat.Where("comment_id=?", commentid).Scan(&likeStat); err != nil {
				return errors.New("mysql查询出错")
			} else {
				//mysql也没有，那就需要插入一条记录
				if likeStat == nil {
					//准备一条数据插入
					likeStat = &model.CommentLikeStat{
						CommentId: int64(commentid),
						LikeCount: 1,
						CreateAt:  gtime.Now(),
						UpdateAt:  gtime.Now(),
					}

					if _, err := dao.CommentLikeStat.Save(likeStat); err != nil {
						return errors.New("mysql插入出错")
					}
					//设置回redis
					if _, err := g.Redis().Do("HSET", keyCount, fieldCount, 1); err != nil {
						return errors.New("redis设置失败")
					}
				} else {
					//更新redis中的结果
					if _, err := g.Redis().Do("HSET", keyCount, fieldCount, likeStat.LikeCount+1); err != nil {
						return errors.New("redis自增出错")
					}
				}
				//延长时间
				if _, err := g.Redis().Do("EXPIRE", keyCount, 900); err != nil {
					return errors.New("redis设置延时出错")
				}
			}

		}
	}

	return nil
}
