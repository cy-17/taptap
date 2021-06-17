package main

import (
	"fmt"
	"github.com/gogf/gf/os/gtime"
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

func main() {

	//m := gmap.New()
	m := gtime.Now()

	fmt.Println(m)


}
