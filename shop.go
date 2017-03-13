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

