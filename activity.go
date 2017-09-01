// 活动服务 
package elemeOpenApi

// 查询店铺邀约活动信息
// shopId 店铺Id
func (flash *Flash) GetInvitedActivityInfos(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(flash.config, "eleme.activity.flash.getInvitedActivityInfos", params)
}

// 报名限量抢购活动
// activityId 活动Id
// activityApplyInfo 活动报名信息
func (flash *Flash) ApplyFlashActivity(activityId_ int64, activityApplyInfo_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["activityId"] = activityId_
	params["activityApplyInfo"] = activityApplyInfo_
	return APIInterface(flash.config, "eleme.activity.flash.applyFlashActivity", params)
}

// 通过店铺Id和活动Id分页查询报名详情
// activityId 活动Id
// shopId 店铺Id
// pageNo 页码
// pageSize 每页数量
func (flash *Flash) GetActivityApplyInfos(activityId_ int64, shopId_ int64, pageNo_ int64, pageSize_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["activityId"] = activityId_
	params["shopId"] = shopId_
	params["pageNo"] = pageNo_
	params["pageSize"] = pageSize_
	return APIInterface(flash.config, "eleme.activity.flash.getActivityApplyInfos", params)
}

// 修改活动菜品库存
// activityId 活动Id
// shopId 店铺Id
// itemId 菜品Id
// stock 库存
func (flash *Flash) UpdateActivityItemStock(activityId_ int64, shopId_ int64, itemId_ int64, stock_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["activityId"] = activityId_
	params["shopId"] = shopId_
	params["itemId"] = itemId_
	params["stock"] = stock_
	return APIInterface(flash.config, "eleme.activity.flash.updateActivityItemStock", params)
}

// 取消活动菜品
// activityId 活动Id
// shopId 店铺Id
// itemId 菜品Id
func (flash *Flash) OfflineFlashActivityItem(activityId_ int64, shopId_ int64, itemId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["activityId"] = activityId_
	params["shopId"] = shopId_
	params["itemId"] = itemId_
	return APIInterface(flash.config, "eleme.activity.flash.offlineFlashActivityItem", params)
}

// 作废店铺与活动的关联关系
// activityId 活动Id
// shopId 店铺Id
func (flash *Flash) InvalidShopActivity(activityId_ int64, shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["activityId"] = activityId_
	params["shopId"] = shopId_
	return APIInterface(flash.config, "eleme.activity.flash.invalidShopActivity", params)
}

