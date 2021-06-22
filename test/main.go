package main

import (
	"fmt"
	"github.com/gogf/gf/frame/g"
)

// 注册输入参数
type UserServiceSignUpReq struct {
	Passport string
	Password string
	Nickname string
	Aasdas   string `default:"123"`
	Aqqqqq   string `default:"123"`
	Ammmmm   string `default:"123"`
}

// 注册请求参数，用于前后端交互参数格式约定
type UserApiSignUpReq struct {
	Passport  string `v:"required|length:6,16#账号不能为空|账号长度应当在:min到:max之间"`
	Password  string `v:"required|length:6,16#请输入确认密码|密码长度应当在:min到:max之间"`
	Password2 string `v:"required|length:6,16|same:Password#密码不能为空|密码长度应当在:min到:max之间|两次密码输入不相等"`
	Nickname  string
}

const (
	//本地保存的文件夹名称
	upload_path string = "/files/"
)

var (
	//BUCKET是你在存储空间的名称
	accessKey = "o1CWnMXIWcff8H4umKJt0_TGkT634fYc8MuOwEAQ"
	secretKey = "Zz3FBrHWN9JVUawhaitzpeRuKIvRo2Lm3BVjtTuK"
	bucket1 = "scutsanqi"
)

// 自定义返回值结构体
type MyPutRet struct {
	Key    string
	Hash   string
	Fsize  int
	Bucket string
	Name   string
}


func main() {

	//七牛云测试
	//localFile := "corn.png"
	//bucket := bucket1
	//key := "testqiniu"
	//// 使用 returnBody 自定义回复格式
	//putPolicy := storage.PutPolicy{
	//	Scope:      bucket,
	//	ReturnBody: `{"key":"$(key)","hash":"$(etag)","fsize":$(fsize),"bucket":"$(bucket)","name":"$(x:name)"}`,
	//}
	//mac := qbox.NewMac(accessKey, secretKey)
	//upToken := putPolicy.UploadToken(mac)
	//cfg := storage.Config{}
	//formUploader := storage.NewFormUploader(&cfg)
	//ret := MyPutRet{}
	//putExtra := storage.PutExtra{
	//	Params: map[string]string{
	//		"x:name": "github logo",
	//	},
	//}
	//err := formUploader.PutFile(context.Background(), &ret, upToken, key, localFile, &putExtra)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(ret.Bucket, ret.Key, ret.Fsize, ret.Hash, ret.Name)

	//res, err := g.Redis().Do("HGET","ttt","ioul")
	//if err != nil {
	//	fmt.Println(res)
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(err)
	//	fmt.Println(res)
	//}
	//
	//ress, err := g.Redis().DoVar("HGET","ttt","ioul")
	//if err != nil {
	//	fmt.Println(ress)
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(err)
	//	fmt.Println(ress.IsNil())
	//}

	//count, err := dao.CommentLike.Where("comment_id=3 and user_id=2").FindCount()
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(count)
	//}

	//var c *model.CommentLike
	//err := dao.CommentLike.Where("comment_id=4 and user_id=2").Scan(&c)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println(c == nil)
	//}


	if _, err := g.Redis().Do("EXPIRE", "123",20); err != nil {
		//redis更新失败
		fmt.Println(err)
	}

}
