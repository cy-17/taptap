package consts

const (

	//以下是400请求错误系列
	// 请求错误
	RequestParamLostCode int    = 400100
	RequestParamLostMsg  string = "请求参数有误"

	//以下是500服务器错误系列
	// MYSQL CURD 常用业务状态码
	CurdStatusOkCode         int    = 200
	CurdStatusOkMsg          string = "Success"
	CurdCreatFailCode        int    = -500200
	CurdCreatFailMsg         string = "新增失败"
	CurdUpdateFailCode       int    = -500201
	CurdUpdateFailMsg        string = "更新失败"
	CurdDeleteFailCode       int    = -500202
	CurdDeleteFailMsg        string = "删除失败"
	CurdSelectFailCode       int    = -500203
	CurdSelectFailMsg        string = "查询无数据"
	CurdRegisterFailCode     int    = -500204
	CurdRegisterFailMsg      string = "注册失败"
	CurdLoginFailCode        int    = -500205
	CurdLoginFailMsg         string = "登录失败"
	CurdRefreshTokenFailCode int    = -500206
	CurdRefreshTokenFailMsg  string = "刷新Token失败"

	// REDIS CURD 常用业务状态码
	RedisCurdStatusOkCode         int    = 200
	RedisCurdStatusOkMsg          string = "Success"
	RedisCurdCreatFailCode        int    = -500300
	RedisCurdCreatFailMsg         string = "新增失败"
	RedisCurdUpdateFailCode       int    = -500301
	RedisCurdUpdateFailMsg        string = "更新失败"
	RedisCurdDeleteFailCode       int    = -500302
	RedisCurdDeleteFailMsg        string = "删除失败"
	RedisCurdSelectFailCode       int    = -500303
	RedisCurdSelectFailMsg        string = "查询无数据"
	RedisCurdRegisterFailCode     int    = -500304
	RedisCurdRegisterFailMsg      string = "注册失败"
	RedisCurdLoginFailCode        int    = -500305
	RedisCurdLoginFailMsg         string = "登录失败"
	RedisCurdRefreshTokenFailCode int    = -500306
	RedisCurdRefreshTokenFailMsg  string = "刷新Token失败"

	// 服务器代码发生错误
	ServerOccurredErrorCode int    = -500100
	ServerOccurredErrorMsg  string = "服务器内部发生代码执行错误"

	//Redis过期时间设定
	RedisExpireTime int = 1200
)
