package commobj

import "time"

type HttpMessageResult struct {
	Start            time.Time
	End              time.Time
	Req              string // 请求
	Res              string // 供应商返回response
	ElapsedTime      time.Duration
	Url              string // 请求url
	ResultCode       int    // code
	ResultMessage    string // 错误原因，错误消息 加上供应商code，message
	SplResultCode    int    // code
	SplResultMessage string // 错误原因，错误消息 加上供应商code，message
	Success          bool   //  是否成功
}


const (
	DATA_TYPE_JSON         = "json"
	DATA_TYPE_XML          = "xml"
	ERROR_DESERIALIZE_XML  = 40005
	ERROR_DESERIALIZE_JSON = 40004
	SUCCESS                = 40000
	ERROR_REQUEST_TIMEOUT  = 40006
	ERROR_SEND_REQUEST     = 40007
)
