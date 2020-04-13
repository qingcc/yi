package commobj

type OrderStatus int

var (
	//订单状态默认不能重复提交，可以重复提交的带有后缀 `_RESUBMIT`
	ORDER_STATUS_APPENDING       OrderStatus = 1 //待处理
	ORDER_STATUS_CONFIRM         OrderStatus = 2 //交易成功
	ORDER_STATUS_CHECK_NEEDED    OrderStatus = 3 //需人工处理
	ORDER_STATUS_FAILED          OrderStatus = 4 //交易失败
	ORDER_STATUS_FAILED_RESUBMIT OrderStatus = 5 //交易取消, 可重复提交
	ORDER_STATUS_CANCELED        OrderStatus = 5 //交易取消

)
