// 订单评论服务 
package elemeOpenApi

// 获取指定订单的评论
// orderId 订单id
func (ugc *Ugc) GetOrderRateByOrderId(orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	return APIInterface(ugc.config, "eleme.ugc.getOrderRateByOrderId", params)
}

// 获取指定订单的评论
// orderIds 订单id
func (ugc *Ugc) GetOrderRatesByOrderIds(orderIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderIds"] = orderIds_
	return APIInterface(ugc.config, "eleme.ugc.getOrderRatesByOrderIds", params)
}

// 获取未回复的评论
// orderIds 订单id
func (ugc *Ugc) GetUnreplyOrderRatesByOrderIds(orderIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderIds"] = orderIds_
	return APIInterface(ugc.config, "eleme.ugc.getUnreplyOrderRatesByOrderIds", params)
}

// 获取指定店铺的评论
// shopId  餐厅id
// startTime   开始时间,只能查询最近90天的数据
// endTime   结束时间
// offset 页面偏移量
// pageSize 页面大小
func (ugc *Ugc) GetOrderRatesByShopId(shopId_ string, startTime_ interface{}, endTime_ interface{}, offset_ int, pageSize_ int) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["startTime"] = startTime_
	params["endTime"] = endTime_
	params["offset"] = offset_
	params["pageSize"] = pageSize_
	return APIInterface(ugc.config, "eleme.ugc.getOrderRatesByShopId", params)
}

// 获取指定店铺的评论
// shopIds 店铺id
// startTime   开始时间,只能查询最近90天的数据
// endTime   结束时间
// offset 页面偏移量
// pageSize 页面大小
func (ugc *Ugc) GetOrderRatesByShopIds(shopIds_ interface{}, startTime_ interface{}, endTime_ interface{}, offset_ int, pageSize_ int) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopIds"] = shopIds_
	params["startTime"] = startTime_
	params["endTime"] = endTime_
	params["offset"] = offset_
	params["pageSize"] = pageSize_
	return APIInterface(ugc.config, "eleme.ugc.getOrderRatesByShopIds", params)
}

// 获取未回复的评论
// shopIds 店铺id
// startTime   开始时间,只能查询最近90天的数据
// endTime   结束时间
// offset 页面偏移量
// pageSize 页面大小
func (ugc *Ugc) GetUnreplyOrderRatesByShopIds(shopIds_ interface{}, startTime_ interface{}, endTime_ interface{}, offset_ int, pageSize_ int) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopIds"] = shopIds_
	params["startTime"] = startTime_
	params["endTime"] = endTime_
	params["offset"] = offset_
	params["pageSize"] = pageSize_
	return APIInterface(ugc.config, "eleme.ugc.getUnreplyOrderRatesByShopIds", params)
}

// 获取店铺的满意度评价信息
// shopId  餐厅id
// score 满意度,取值范围为1~5，1为最不满意，5为非常满意
// startTime   开始时间,只能查询最近90天的数据
// endTime   结束时间
// offset 页面偏移量
// pageSize 页面大小
func (ugc *Ugc) GetOrderRatesByShopAndRating(shopId_ string, score_ int, startTime_ interface{}, endTime_ interface{}, offset_ int, pageSize_ int) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["score"] = score_
	params["startTime"] = startTime_
	params["endTime"] = endTime_
	params["offset"] = offset_
	params["pageSize"] = pageSize_
	return APIInterface(ugc.config, "eleme.ugc.getOrderRatesByShopAndRating", params)
}

// 获取单个商品的评论
// itemId  商品id
// startTime   开始时间,只能查询最近90天的数据
// endTime   结束时间
// offset 页面偏移量
func (ugc *Ugc) GetItemRatesByItemId(itemId_ string, startTime_ interface{}, endTime_ interface{}, offset_ int) (interface{}, error) {
	params := make(map[string]interface{})
	params["itemId"] = itemId_
	params["startTime"] = startTime_
	params["endTime"] = endTime_
	params["offset"] = offset_
	return APIInterface(ugc.config, "eleme.ugc.getItemRatesByItemId", params)
}

// 获取多个商品的评论
// itemIds 商品id
// startTime   开始时间,只能查询最近90天的数据
// endTime   结束时间
// offset 页面偏移量
// pageSize 页面大小
func (ugc *Ugc) GetItemRatesByItemIds(itemIds_ interface{}, startTime_ interface{}, endTime_ interface{}, offset_ int, pageSize_ int) (interface{}, error) {
	params := make(map[string]interface{})
	params["itemIds"] = itemIds_
	params["startTime"] = startTime_
	params["endTime"] = endTime_
	params["offset"] = offset_
	params["pageSize"] = pageSize_
	return APIInterface(ugc.config, "eleme.ugc.getItemRatesByItemIds", params)
}

// 获取多个商品未回复的评论
// itemIds 店铺id
// startTime   开始时间,只能查询最近90天的数据
// endTime   结束时间
// offset 页面偏移量
// pageSize 页面大小
func (ugc *Ugc) GetUnreplyItemRatesByItemIds(itemIds_ interface{}, startTime_ interface{}, endTime_ interface{}, offset_ int, pageSize_ int) (interface{}, error) {
	params := make(map[string]interface{})
	params["itemIds"] = itemIds_
	params["startTime"] = startTime_
	params["endTime"] = endTime_
	params["offset"] = offset_
	params["pageSize"] = pageSize_
	return APIInterface(ugc.config, "eleme.ugc.getUnreplyItemRatesByItemIds", params)
}

// 回复指定类型的评论
// rateId 评论编号
// replyType 评论类型
// reply 回复的内容
func (ugc *Ugc) ReplyRateByRateId(rateId_ string, replyType_ interface{}, reply_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["rateId"] = rateId_
	params["replyType"] = replyType_
	params["reply"] = reply_
	return APIInterface(ugc.config, "eleme.ugc.replyRateByRateId", params)
}

// 回复指定类型的评论
// rateIds  评论编号
// replyType 评论类型
// reply 回复的内容
func (ugc *Ugc) ReplyRateByRateIds(rateIds_ interface{}, replyType_ interface{}, reply_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["rateIds"] = rateIds_
	params["replyType"] = replyType_
	params["reply"] = reply_
	return APIInterface(ugc.config, "eleme.ugc.replyRateByRateIds", params)
}

// 回复订单未回复的评论
// orderId 订单id
// reply 回复内容
func (ugc *Ugc) ReplyRateByOrderId(orderId_ string, reply_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	params["reply"] = reply_
	return APIInterface(ugc.config, "eleme.ugc.replyRateByOrderId", params)
}

// 批量回复订单未回复的评论
// orderIds 订单id
// reply 回复信息
// reply 回复信息
func (ugc *Ugc) ReplyCommentByOrderIds(orderIds_ interface{}, reply_ string, reply_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderIds"] = orderIds_
	params["reply"] = reply_
	params["reply"] = reply_
	return APIInterface(ugc.config, "eleme.ugc.replyCommentByOrderIds", params)
}

// 回复商品回复的评论
// itemId 商品id
// reply 回复内容
// startTime   开始时间,只能查询最近90天的数据
// endTime   结束时间
func (ugc *Ugc) ReplyRatesByItemId(itemId_ string, reply_ string, startTime_ interface{}, endTime_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["itemId"] = itemId_
	params["reply"] = reply_
	params["startTime"] = startTime_
	params["endTime"] = endTime_
	return APIInterface(ugc.config, "eleme.ugc.replyRatesByItemId", params)
}

// 回复多个商品评论
// itemIds 商品d
// reply 回复信息
// startTime 开始时间,只能查询最近90天的数据
// endTime 结束时间
func (ugc *Ugc) ReplyRatesByItemIds(itemIds_ interface{}, reply_ string, startTime_ interface{}, endTime_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["itemIds"] = itemIds_
	params["reply"] = reply_
	params["startTime"] = startTime_
	params["endTime"] = endTime_
	return APIInterface(ugc.config, "eleme.ugc.replyRatesByItemIds", params)
}

