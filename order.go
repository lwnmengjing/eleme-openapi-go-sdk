// 订单服务 
package elemeOpenApi

// 获取订单
// orderId 订单Id
func (order *Order) GetOrder(orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	return APIInterface(order.config, "eleme.order.getOrder", params)
}

// 批量获取订单
// orderIds 订单Id的列表
func (order *Order) MgetOrders(orderIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderIds"] = orderIds_
	return APIInterface(order.config, "eleme.order.mgetOrders", params)
}

// 确认订单
// orderId 订单Id
func (order *Order) ConfirmOrder(orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	return APIInterface(order.config, "eleme.order.confirmOrder", params)
}

// 取消订单
// orderId 订单Id
// type 取消原因
// remark 备注说明
func (order *Order) CancelOrder(orderId_ string, type_ interface{}, remark_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	params["type"] = type_
	params["remark"] = remark_
	return APIInterface(order.config, "eleme.order.cancelOrder", params)
}

// 同意退单/取消单
// orderId 订单Id
func (order *Order) AgreeRefund(orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	return APIInterface(order.config, "eleme.order.agreeRefund", params)
}

// 不同意退单/取消单
// orderId 订单Id
// reason 商家不同意退单原因
func (order *Order) DisagreeRefund(orderId_ string, reason_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	params["reason"] = reason_
	return APIInterface(order.config, "eleme.order.disagreeRefund", params)
}

// 获取订单配送记录
// orderId 订单Id
func (order *Order) GetDeliveryStateRecord(orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	return APIInterface(order.config, "eleme.order.getDeliveryStateRecord", params)
}

// 批量获取订单最新配送记录
// orderIds 订单Id列表
func (order *Order) BatchGetDeliveryStates(orderIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderIds"] = orderIds_
	return APIInterface(order.config, "eleme.order.batchGetDeliveryStates", params)
}

// 配送异常或者物流拒单后选择自行配送
// orderId 订单Id
func (order *Order) DeliveryBySelf(orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	return APIInterface(order.config, "eleme.order.deliveryBySelf", params)
}

// 配送异常或者物流拒单后选择不再配送
// orderId 订单Id
func (order *Order) NoMoreDelivery(orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	return APIInterface(order.config, "eleme.order.noMoreDelivery", params)
}

// 订单确认送达
// orderId 订单ID
func (order *Order) ReceivedOrder(orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	return APIInterface(order.config, "eleme.order.receivedOrder", params)
}

// 回复催单
// remindId 催单Id
// type 回复类别
// content 回复内容,如果type为custom,content必填,回复内容不能超过30个字符
func (order *Order) ReplyReminder(remindId_ string, type_ interface{}, content_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["remindId"] = remindId_
	params["type"] = type_
	params["content"] = content_
	return APIInterface(order.config, "eleme.order.replyReminder", params)
}

// 获取指定订单菜品活动价格.
// orderId 订单Id
func (order *Order) GetCommodities(orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	return APIInterface(order.config, "eleme.order.getCommodities", params)
}

// 批量获取订单菜品活动价格
// orderIds 订单Id列表
func (order *Order) MgetCommodities(orderIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderIds"] = orderIds_
	return APIInterface(order.config, "eleme.order.mgetCommodities", params)
}

// 获取订单退款信息
// orderId 订单Id
func (order *Order) GetRefundOrder(orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	return APIInterface(order.config, "eleme.order.getRefundOrder", params)
}

// 批量获取订单退款信息
// orderIds 订单Id列表
func (order *Order) MgetRefundOrders(orderIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderIds"] = orderIds_
	return APIInterface(order.config, "eleme.order.mgetRefundOrders", params)
}

// 取消呼叫配送
// orderId 订单Id
func (order *Order) CancelDelivery(orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	return APIInterface(order.config, "eleme.order.cancelDelivery", params)
}

// 呼叫配送
// orderId 订单Id
// fee 小费,1-8之间的整数
func (order *Order) CallDelivery(orderId_ string, fee_ int) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	params["fee"] = fee_
	return APIInterface(order.config, "eleme.order.callDelivery", params)
}

