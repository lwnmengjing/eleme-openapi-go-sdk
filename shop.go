// 店铺服务 
package elemeOpenApi

// 查询店铺信息
// shopId 店铺Id
func (shop *Shop) GetShop(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(shop.config, "eleme.shop.getShop", params)
}

// 更新店铺基本信息
// shopId 店铺Id
// properties 店铺属性
func (shop *Shop) UpdateShop(shopId_ int64, properties_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["properties"] = properties_
	return APIInterface(shop.config, "eleme.shop.updateShop", params)
}

// 批量获取店铺简要
// shopIds 店铺Id的列表
func (shop *Shop) MgetShopStatus(shopIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopIds"] = shopIds_
	return APIInterface(shop.config, "eleme.shop.mgetShopStatus", params)
}

// 设置送达时间
// shopId 店铺Id
// deliveryBasicMins 配送基准时间(单位分钟)
// deliveryAdjustMins 配送调整时间(单位分钟)
func (shop *Shop) SetDeliveryTime(shopId_ int64, deliveryBasicMins_ int, deliveryAdjustMins_ int) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["deliveryBasicMins"] = deliveryBasicMins_
	params["deliveryAdjustMins"] = deliveryAdjustMins_
	return APIInterface(shop.config, "eleme.shop.setDeliveryTime", params)
}

// 设置是否支持在线退单
// shopId 店铺Id
// enable 是否支持
func (shop *Shop) SetOnlineRefund(shopId_ int64, enable_ bool) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["enable"] = enable_
	return APIInterface(shop.config, "eleme.shop.setOnlineRefund", params)
}

// 设置是否支持预定单及预定天数
// shopId 店铺id
// enabled 是否支持预订
// maxBookingDays 最大预定天数
func (shop *Shop) SetBookingStatus(shopId_ int64, enabled_ bool, maxBookingDays_ int) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["enabled"] = enabled_
	params["maxBookingDays"] = maxBookingDays_
	return APIInterface(shop.config, "eleme.shop.setBookingStatus", params)
}

// 批量通过店铺Id获取Oid
// shopIds 店铺Id的列表
func (shop *Shop) GetOidByShopIds(shopIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopIds"] = shopIds_
	return APIInterface(shop.config, "eleme.shop.getOidByShopIds", params)
}

// 更新店铺营业时间预设置
// shopId 店铺 id 
// weekSetting 一周营业时间预设置, 参考 OShopBusyLevelSetting weekSetting 字段定义 
// dateSetting 特定日期营业时间预设置, 参考 OShopBusyLevelSetting dateSetting 字段定义 
func (shop *Shop) UpdateBusyLevelSetting(shopId_ int64, weekSetting_ interface{}, dateSetting_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["weekSetting"] = weekSetting_
	params["dateSetting"] = dateSetting_
	return APIInterface(shop.config, "eleme.shop.updateBusyLevelSetting", params)
}

// 获取店铺营业时间预设置
// shopId 店铺 id 
func (shop *Shop) GetBusyLevelSetting(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(shop.config, "eleme.shop.getBusyLevelSetting", params)
}

// 提交开店申请接口
// openStoreMessage 开店申请表单
func (setup *Setup) SubmitOpenStoreMessageAudit(openStoreMessage_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["openStoreMessage"] = openStoreMessage_
	return APIInterface(setup.config, "eleme.shop.setup.submitOpenStoreMessageAudit", params)
}

// 星巴克提交开店申请接口
// openStoreMessage 开店申请表单
func (setup *Setup) SubmitOpenStoreForMermaid(openStoreMessage_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["openStoreMessage"] = openStoreMessage_
	return APIInterface(setup.config, "eleme.shop.setup.submitOpenStoreForMermaid", params)
}

// 更新申请信息接口
// updateStoreMessageBody 开店申请表单
func (setup *Setup) UpdateOpenStoreMessageAudit(updateStoreMessageBody_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["updateStoreMessageBody"] = updateStoreMessageBody_
	return APIInterface(setup.config, "eleme.shop.setup.updateOpenStoreMessageAudit", params)
}

// 查询请求状态接口
// submitId 请求提交id
func (setup *Setup) QueryProcessStatusBySubmitId(submitId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["submitId"] = submitId_
	return APIInterface(setup.config, "eleme.shop.setup.queryProcessStatusBySubmitId", params)
}

// 图片上传处理接口（5M以内图片）
// imageBase64 base64字节流
func (setup *Setup) UploadImage(imageBase64_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["imageBase64"] = imageBase64_
	return APIInterface(setup.config, "eleme.shop.setup.uploadImage", params)
}

// 图片上传处理接口（500K以内图片）
// imageBase64 base64字节流
func (setup *Setup) UploadMinImage(imageBase64_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["imageBase64"] = imageBase64_
	return APIInterface(setup.config, "eleme.shop.setup.uploadMinImage", params)
}

// 远程上传图片接口
// url 图片url
func (setup *Setup) UploadImageWithRemoteUrl(url_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["url"] = url_
	return APIInterface(setup.config, "eleme.shop.setup.uploadImageWithRemoteUrl", params)
}

