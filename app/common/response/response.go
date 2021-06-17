package response

import "github.com/gogf/gf/net/ghttp"

// 数据返回的Json结构
type JsonResponse struct {
	Code    int         `json:"code"`    // 错误码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据(业务接口定义具体数据结构)
}

// 标准返回结果数据结构封装
// Status代表请求结果，0成功，1 请求有误，2服务器内部错误//
//	Json
//	@Description:
//	@param r *ghttp.Request
//	@param code int
//	@param status int
//	@param message string
//	@param data ...interface{}
//
func Json(r *ghttp.Request, code int, status int, message string, data interface{}) {

	//进行状态请求判断，是否请求成功
	switch status {
	case 0:
		r.Response.Status = 200
	case 1:
		r.Response.Status = 400
		case 2:
		r.Response.Status = 500
	}

	//同一封装返回json数据
	r.Response.WriteJson(JsonResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// 返回JSON数据并退出当前HTTP执行函数
func JsonExit(r *ghttp.Request,code int, status int, msg string, data interface{}){
	Json(r,code,status,msg,data)
	r.Exit()
}
