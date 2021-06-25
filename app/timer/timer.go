package timer

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/os/gtimer"
	"github.com/gogf/gf/util/gconv"
	"san616qi/app/dao"
	"san616qi/app/model"
	"strings"
	"time"
)

func init() {
	gtimer.AddSingleton(10*time.Second, redisCommentLikeFlush)
	gtimer.AddSingleton(10*time.Second, redisCommentLikeCountFlush)
	gtimer.AddSingleton(24*time.Hour,redisUserProfileFlush)
}

//定时刷新点赞数量
func redisCommentLikeCountFlush() {

	//获取某一个分区的评论点赞数量

	for i:=0;i<5;i++ {

		//组装查询的key
		key := "commentlikecount_" + gconv.String(i)

		res, _ := g.Redis().DoVar("HGETALL", key)
		resMap := gconv.Map(res)

		//flush到db中
		//拿到的数据为 commentid count
		for k, v := range resMap {
			if _, err := dao.CommentLikeStat.DB.Exec("insert into comment_like_stat(comment_id,like_count,create_at,update_at) VALUES (?,?,?,?) ON DUPLICATE KEY UPDATE like_count=VALUES(like_count),update_at=VALUES(update_at)",
				gconv.Int64(k), gconv.Int(v), gtime.Now(), gtime.Now())
				err != nil {
				fmt.Println(err)
			}
		}
	}
	time.Sleep(10 * time.Second)

}

//定时刷新点赞状态
func redisCommentLikeFlush() {

	for i:=0;i<5;i++ {

		//组装查询的key
		key := "commentlike_" + gconv.String(i)

		//获取某一个分区的评论的点赞以及点赞状态
		res, _ := g.Redis().DoVar("HGETALL", key)
		resMap := gconv.Map(res)

		//flush到db中
		//拿到的数据为 commentid::userid  status
		for k, v := range resMap {

			keys := strings.Split(k, "::")

			if _, err := dao.CommentLike.DB.Exec("insert into comment_like(user_id,comment_id,comment_like_stat,comment_like_time,create_at,update_at) VALUES (?,?,?,?,?,?) ON DUPLICATE KEY UPDATE comment_like_stat=VALUES(comment_like_stat),update_at=VALUES(update_at)",
				gconv.Int64(keys[1]), gconv.Int64(keys[0]), gconv.Int(v), gtime.Now(), gtime.Now(), gtime.Now())
				err != nil {
				fmt.Println(err)
			}

		}
	}
	time.Sleep(10 * time.Second)
}

//定时刷新用户个人信息
func redisUserProfileFlush() {

	for i:=0;i<10;i++ {

		//组装查询的key
		key := "userprofile_" + gconv.String(i)

		res, _ := g.Redis().DoVar("HGETALL", key)
		resMap := gconv.Map(res)

		for k,v := range resMap	{

			var user *model.UserServiceUpdateProfileReq
			err := gconv.Struct(v, &user)
			if err != nil {
				return
			}

			if _, err := dao.User.Where("user_id=?",gconv.Int64(k)).Update(user); err != nil {
				return
			}

		}

	}

	time.Sleep(24*time.Hour)

}

