package response

type RespCode int

const (
	SUCCESS                    RespCode = 0    //操作成功
	ErrorServer                RespCode = -1   //操作失败
	ErrorParams                RespCode = 4000 //请求参数错误
	ErrorAuthCheckTokenFail    RespCode = 4001 //Token鉴权失败
	ErrorAuthCheckTokenTimeout RespCode = 4002 //Token已超时
	ErrorAuthToken             RespCode = 4003 //Token生成失败
	ErrorAuth                  RespCode = 4004 //Token错误
	ErrorRequestTooFrequent    RespCode = 4005 //请求太频繁
)
