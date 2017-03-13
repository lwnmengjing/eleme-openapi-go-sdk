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

