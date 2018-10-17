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
func (ugc *Ugc) GetOrderRatesByShopId(shopId_ string, startTime_ string, endTime_ string, offset_ int, pageSize_ int) (interface{}, error) {
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
func (ugc *Ugc) GetOrderRatesByShopIds(shopIds_ interface{}, startTime_ string, endTime_ string, offset_ int, pageSize_ int) (interface{}, error) {
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
func (ugc *Ugc) GetUnreplyOrderRatesByShopIds(shopIds_ interface{}, startTime_ string, endTime_ string, offset_ int, pageSize_ int) (interface{}, error) {
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
func (ugc *Ugc) GetOrderRatesByShopAndRating(shopId_ string, score_ int, startTime_ string, endTime_ string, offset_ int, pageSize_ int) (interface{}, error) {
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
// pageSize 页面大小
func (ugc *Ugc) GetItemRatesByItemId(itemId_ string, startTime_ string, endTime_ string, offset_ int, pageSize_ int) (interface{}, error) {
	params := make(map[string]interface{})
	params["itemId"] = itemId_
	params["startTime"] = startTime_
	params["endTime"] = endTime_
	params["offset"] = offset_
	params["pageSize"] = pageSize_
	return APIInterface(ugc.config, "eleme.ugc.getItemRatesByItemId", params)
}

// 获取多个商品的评论
// itemIds 商品id
// startTime   开始时间,只能查询最近90天的数据
// endTime   结束时间
// offset 页面偏移量
// pageSize 页面大小
func (ugc *Ugc) GetItemRatesByItemIds(itemIds_ interface{}, startTime_ string, endTime_ string, offset_ int, pageSize_ int) (interface{}, error) {
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
func (ugc *Ugc) GetUnreplyItemRatesByItemIds(itemIds_ interface{}, startTime_ string, endTime_ string, offset_ int, pageSize_ int) (interface{}, error) {
	params := make(map[string]interface{})
	params["itemIds"] = itemIds_
	params["startTime"] = startTime_
	params["endTime"] = endTime_
	params["offset"] = offset_
	params["pageSize"] = pageSize_
	return APIInterface(ugc.config, "eleme.ugc.getUnreplyItemRatesByItemIds", params)
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
func (ugc *Ugc) ReplyCommentByOrderIds(orderIds_ interface{}, reply_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderIds"] = orderIds_
	params["reply"] = reply_
	return APIInterface(ugc.config, "eleme.ugc.replyCommentByOrderIds", params)
}

// 回复商品回复的评论
// itemId 商品id
// reply 回复内容
// startTime   开始时间,只能查询最近90天的数据
// endTime   结束时间
func (ugc *Ugc) ReplyRatesByItemId(itemId_ string, reply_ string, startTime_ string, endTime_ string) (interface{}, error) {
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
func (ugc *Ugc) ReplyRatesByItemIds(itemIds_ interface{}, reply_ string, startTime_ string, endTime_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["itemIds"] = itemIds_
	params["reply"] = reply_
	params["startTime"] = startTime_
	params["endTime"] = endTime_
	return APIInterface(ugc.config, "eleme.ugc.replyRatesByItemIds", params)
}

// 通过rateId和shopId 回复指定类型的评论
// rateId 评论编号
// shopId  餐厅id
// replyType 评论类型
// reply 回复的内容
func (ugc *Ugc) ReplyRateByRateIdAndShopId(rateId_ string, shopId_ string, replyType_ interface{}, reply_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["rateId"] = rateId_
	params["shopId"] = shopId_
	params["replyType"] = replyType_
	params["reply"] = reply_
	return APIInterface(ugc.config, "eleme.ugc.replyRateByRateIdAndShopId", params)
}

// 通过rateIds和shopId 批量回复指定类型的评论
// rateIds  评论编号
// shopId  餐厅id
// replyType 评论类型
// reply 回复的内容
func (ugc *Ugc) ReplyRateByRateIdsAndShopId(rateIds_ interface{}, shopId_ string, replyType_ interface{}, reply_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["rateIds"] = rateIds_
	params["shopId"] = shopId_
	params["replyType"] = replyType_
	params["reply"] = reply_
	return APIInterface(ugc.config, "eleme.ugc.replyRateByRateIdsAndShopId", params)
}

// 根据订单ID赠送代金券给该订单的评价用户
// orderId  订单编号
// coupon 需要赠送的代金券信息
func (ugc *Ugc) SendCouponByOrderId(orderId_ string, coupon_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	params["coupon"] = coupon_
	return APIInterface(ugc.config, "eleme.ugc.sendCouponByOrderId", params)
}

// 根据订单ID获取该订单评价用户的可赠券状态
// orderId  订单编号
func (ugc *Ugc) GetOrderCouponStatus(orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	return APIInterface(ugc.config, "eleme.ugc.getOrderCouponStatus", params)
}

// 根据订单ID集合获取该订单的已赠券信息集合
// orderIds 订单编号集合
func (ugc *Ugc) GetCouponsByOrderIds(orderIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderIds"] = orderIds_
	return APIInterface(ugc.config, "eleme.ugc.getCouponsByOrderIds", params)
}

// 获取店铺的推荐赠送代金券信息
// shopId 餐厅ID
func (ugc *Ugc) GetRecommendCouponByShopId(shopId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(ugc.config, "eleme.ugc.getRecommendCouponByShopId", params)
}

// 查询评价信息
// rateQuery 评价查询参数
func (ugc *Ugc) GetORateResult(rateQuery_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["rateQuery"] = rateQuery_
	return APIInterface(ugc.config, "eleme.ugc.getORateResult", params)
}

// 统计评价信息数量
// rateQuery 评价查询参数
func (ugc *Ugc) CountORateResult(rateQuery_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["rateQuery"] = rateQuery_
	return APIInterface(ugc.config, "eleme.ugc.countORateResult", params)
}

// 通过rateIds和shopId 回复饿了么星选评论
// rateIds  评论编号(订单维度)
// shopId  饿了么侧餐厅id
// reply 回复的内容
func (ugc *Ugc) ReplyBaiduRate(rateIds_ interface{}, shopId_ string, reply_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["rateIds"] = rateIds_
	params["shopId"] = shopId_
	params["reply"] = reply_
	return APIInterface(ugc.config, "eleme.ugc.replyBaiduRate", params)
}

// 根据rateId和shopId 赠送代金券给该饿了么星选评价对应订单的评价用户
// rateId  评论编号(订单维度)
// shopId  餐厅id
// coupon 需要赠送的代金券信息
func (ugc *Ugc) SendBaiduCoupon(rateId_ string, shopId_ string, coupon_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["rateId"] = rateId_
	params["shopId"] = shopId_
	params["coupon"] = coupon_
	return APIInterface(ugc.config, "eleme.ugc.sendBaiduCoupon", params)
}

// 根据rateId和shopId获取该订单评价用户的可赠券状态
// rateId  评论编号(订单维度)
// shopId  餐厅id
// rateDataType 评价数据类型
func (ugc *Ugc) GetRateCouponStatus(rateId_ string, shopId_ string, rateDataType_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["rateId"] = rateId_
	params["shopId"] = shopId_
	params["rateDataType"] = rateDataType_
	return APIInterface(ugc.config, "eleme.ugc.getRateCouponStatus", params)
}

