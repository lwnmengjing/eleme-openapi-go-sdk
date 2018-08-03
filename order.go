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

// 确认订单(推荐)
// orderId 订单Id
func (order *Order) ConfirmOrderLite(orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	return APIInterface(order.config, "eleme.order.confirmOrderLite", params)
}

// 取消订单(推荐)
// orderId 订单Id
// type 取消原因
// remark 备注说明
func (order *Order) CancelOrderLite(orderId_ string, type_ interface{}, remark_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	params["type"] = type_
	params["remark"] = remark_
	return APIInterface(order.config, "eleme.order.cancelOrderLite", params)
}

// 同意退单/同意取消单(推荐)
// orderId 订单Id
func (order *Order) AgreeRefundLite(orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	return APIInterface(order.config, "eleme.order.agreeRefundLite", params)
}

// 不同意退单/不同意取消单(推荐)
// orderId 订单Id
// reason 商家不同意退单原因
func (order *Order) DisagreeRefundLite(orderId_ string, reason_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	params["reason"] = reason_
	return APIInterface(order.config, "eleme.order.disagreeRefundLite", params)
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

// 配送异常或者物流拒单后选择自行配送(推荐)
// orderId 订单Id
func (order *Order) DeliveryBySelfLite(orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	return APIInterface(order.config, "eleme.order.deliveryBySelfLite", params)
}

// 配送异常或者物流拒单后选择不再配送(推荐)
// orderId 订单Id
func (order *Order) NoMoreDeliveryLite(orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	return APIInterface(order.config, "eleme.order.noMoreDeliveryLite", params)
}

// 订单确认送达
// orderId 订单ID
func (order *Order) ReceivedOrderLite(orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	return APIInterface(order.config, "eleme.order.receivedOrderLite", params)
}

// 订单确认送出(自配送)
// orderId 订单ID
// phone 配送者电话
func (order *Order) StartDeliveryBySelf(orderId_ string, phone_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	params["phone"] = phone_
	return APIInterface(order.config, "eleme.order.startDeliveryBySelf", params)
}

// 订单确认送达(自配送)
// orderId 订单ID
// phone 配送者电话
func (order *Order) CompleteDeliveryBySelf(orderId_ string, phone_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	params["phone"] = phone_
	return APIInterface(order.config, "eleme.order.completeDeliveryBySelf", params)
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

// 获取店铺未回复的催单
// shopId 店铺id
func (order *Order) GetUnreplyReminders(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(order.config, "eleme.order.getUnreplyReminders", params)
}

// 查询店铺未处理订单
// shopId 店铺id
func (order *Order) GetUnprocessOrders(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(order.config, "eleme.order.getUnprocessOrders", params)
}

// 查询店铺未处理的取消单
// shopId 店铺id
func (order *Order) GetCancelOrders(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(order.config, "eleme.order.getCancelOrders", params)
}

// 查询店铺未处理的退单
// shopId 店铺id
func (order *Order) GetRefundOrders(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(order.config, "eleme.order.getRefundOrders", params)
}

// 查询全部订单
// shopId 店铺id
// pageNo 页码。取值范围:大于零的整数最大限制为100
// pageSize 每页获取条数。最小值1，最大值50。
// date 日期,默认当天,格式:yyyy-MM-dd
func (order *Order) GetAllOrders(shopId_ int64, pageNo_ int, pageSize_ int, date_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["pageNo"] = pageNo_
	params["pageSize"] = pageSize_
	params["date"] = date_
	return APIInterface(order.config, "eleme.order.getAllOrders", params)
}

// 批量查询订单是否支持索赔
// orderIds 索赔订单Id的列表
func (order *Order) QuerySupportedCompensationOrders(orderIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderIds"] = orderIds_
	return APIInterface(order.config, "eleme.order.querySupportedCompensationOrders", params)
}

// 批量申请索赔
// requests 索赔请求的列表
func (order *Order) BatchApplyCompensations(requests_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["requests"] = requests_
	return APIInterface(order.config, "eleme.order.batchApplyCompensations", params)
}

// 批量查询索赔结果
// orderIds 索赔订单Id的列表
func (order *Order) QueryCompensationOrders(orderIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderIds"] = orderIds_
	return APIInterface(order.config, "eleme.order.queryCompensationOrders", params)
}

// 众包订单询价，获取配送费
// orderId 订单Id
func (order *Order) GetDeliveryFeeForCrowd(orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	return APIInterface(order.config, "eleme.order.getDeliveryFeeForCrowd", params)
}

// 评价骑手
// orderId 订单Id
// evaluationInfo 评价信息
func (order *Order) EvaluateRider(orderId_ string, evaluationInfo_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	params["evaluationInfo"] = evaluationInfo_
	return APIInterface(order.config, "eleme.order.evaluateRider", params)
}

// 批量获取骑手评价信息
// orderIds 订单Id的列表
func (order *Order) MgetEvaluationInfos(orderIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderIds"] = orderIds_
	return APIInterface(order.config, "eleme.order.mgetEvaluationInfos", params)
}

// 批量获取是否可以评价骑手
// orderIds 订单Id的列表
func (order *Order) MgetEvaluationStatus(orderIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderIds"] = orderIds_
	return APIInterface(order.config, "eleme.order.mgetEvaluationStatus", params)
}

// 批量获取订单加小费信息
// orderIds 订单Id的列表
func (order *Order) MgetDeliveryTipInfos(orderIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderIds"] = orderIds_
	return APIInterface(order.config, "eleme.order.mgetDeliveryTipInfos", params)
}

// 订单加小费
// orderId 订单Id
// tip 小费金额
func (order *Order) AddDeliveryTipByOrderId(orderId_ string, tip_ int) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	params["tip"] = tip_
	return APIInterface(order.config, "eleme.order.addDeliveryTipByOrderId", params)
}

// 主动发起退单
// orderId 订单Id
// type 取消原因
// remark 备注说明
func (order *Order) ApplyRefund(orderId_ string, type_ interface{}, remark_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	params["type"] = type_
	params["remark"] = remark_
	return APIInterface(order.config, "eleme.order.applyRefund", params)
}

// 非自配送餐厅标记已出餐
// orderId 订单Id
func (order *Order) SetOrderPrepared(orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	return APIInterface(order.config, "eleme.order.setOrderPrepared", params)
}

// 查询已出餐列表
// orderIds 查询已出餐订单Id的列表
func (order *Order) GetPreparedTimesByOrderIds(orderIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderIds"] = orderIds_
	return APIInterface(order.config, "eleme.order.getPreparedTimesByOrderIds", params)
}

// 查询顾客联系方式
// orderIds 订单ID的列表
func (order *Order) MgetUserSimpleInfoByOrderIds(orderIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderIds"] = orderIds_
	return APIInterface(order.config, "eleme.order.mgetUserSimpleInfoByOrderIds", params)
}

// 商家部分退款
// orderId 订单id
// refundOrderMessage 退款详情
func (order *Order) RefundPart(orderId_ string, refundOrderMessage_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	params["refundOrderMessage"] = refundOrderMessage_
	return APIInterface(order.config, "eleme.order.refundPart", params)
}

// 设置订单开票地址
// orderId 订单id
// invoiceUrl 开票地址
func (order *Order) SetInvoiceUrl(orderId_ string, invoiceUrl_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	params["invoiceUrl"] = invoiceUrl_
	return APIInterface(order.config, "eleme.order.setInvoiceUrl", params)
}

// 获取订单配送轨迹
// orderId 订单Id
func (delivery *Delivery) GetDeliveryRoutes(orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	return APIInterface(delivery.config, "eleme.order.delivery.getDeliveryRoutes", params)
}

