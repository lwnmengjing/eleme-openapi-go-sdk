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

