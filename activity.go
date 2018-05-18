// 活动服务 
package elemeOpenApi

// 创建红包活动
// createInfo 创建红包活动的结构体
func (coupon *Coupon) CreateCouponActivity(createInfo_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["createInfo"] = createInfo_
	return APIInterface(coupon.config, "eleme.activity.coupon.createCouponActivity", params)
}

// 向指定用户发放红包
// shopId 店铺Id
// couponActivityId 红包活动Id
// mobiles 需要发放红包的用户手机号列表
func (coupon *Coupon) GiveOutCoupons(shopId_ int64, couponActivityId_ int64, mobiles_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["couponActivityId"] = couponActivityId_
	params["mobiles"] = mobiles_
	return APIInterface(coupon.config, "eleme.activity.coupon.giveOutCoupons", params)
}

// 定向赠红包
// shopId 店铺Id
// mobile 需要发放红包的用户手机号
// couponTemplate 定向赠红包的模板信息
func (coupon *Coupon) PresentCoupon(shopId_ int64, mobile_ string, couponTemplate_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["mobile"] = mobile_
	params["couponTemplate"] = couponTemplate_
	return APIInterface(coupon.config, "eleme.activity.coupon.presentCoupon", params)
}

// 分页查询店铺红包活动信息
// shopId 店铺Id
// couponActivityType 红包活动类型
// activityStatus 活动状态
// pageNo 页码（第几页）
// pageSize 每页数量
func (coupon *Coupon) QueryCouponActivities(shopId_ int64, couponActivityType_ interface{}, activityStatus_ interface{}, pageNo_ int, pageSize_ int) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["couponActivityType"] = couponActivityType_
	params["activityStatus"] = activityStatus_
	params["pageNo"] = pageNo_
	params["pageSize"] = pageSize_
	return APIInterface(coupon.config, "eleme.activity.coupon.queryCouponActivities", params)
}

// 分页查询店铺红包领取详情
// shopId 店铺Id
// couponActivityId 红包活动Id
// couponStatus 红包状态
// pageNo 页码（第几页）
// pageSize 每页数量
func (coupon *Coupon) QueryReceivedCouponDetails(shopId_ int64, couponActivityId_ int64, couponStatus_ interface{}, pageNo_ int, pageSize_ int) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["couponActivityId"] = couponActivityId_
	params["couponStatus"] = couponStatus_
	params["pageNo"] = pageNo_
	params["pageSize"] = pageSize_
	return APIInterface(coupon.config, "eleme.activity.coupon.queryReceivedCouponDetails", params)
}

// 托管单店红包服务
// shopIds 餐厅id列表,长度不能超过20
// hostedType 红包服务业务类型,暂只支持超级会员,"SUPER_VIP"
// discounts 扣减额,请设置在[4,15]元,小数点后最多1位
func (coupon *Coupon) HostShops(shopIds_ interface{}, hostedType_ interface{}, discounts_ float64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopIds"] = shopIds_
	params["hostedType"] = hostedType_
	params["discounts"] = discounts_
	return APIInterface(coupon.config, "eleme.activity.coupon.hostShops", params)
}

// 查询红包服务托管情况
// shopIds 餐厅id列表,长度不能超过20
// hostedType 红包服务业务类型,暂只支持超级会员,"SUPER_VIP"
func (coupon *Coupon) QueryHostInfo(shopIds_ interface{}, hostedType_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopIds"] = shopIds_
	params["hostedType"] = hostedType_
	return APIInterface(coupon.config, "eleme.activity.coupon.queryHostInfo", params)
}

// 取消托管单店红包服务
// shopIds 餐厅id列表,长度不能超过20
// hostedType 红包服务业务类型,暂只支持超级会员,"SUPER_VIP"
func (coupon *Coupon) UnhostShops(shopIds_ interface{}, hostedType_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopIds"] = shopIds_
	params["hostedType"] = hostedType_
	return APIInterface(coupon.config, "eleme.activity.coupon.unhostShops", params)
}

// 更改单店红包服务托管方式
// shopId 店铺Id
// hostedType 红包服务业务类型,暂只支持超级会员,"SUPER_VIP"
// oActivityServiceDetails 服务内容
func (coupon *Coupon) RehostShop(shopId_ int64, hostedType_ interface{}, oActivityServiceDetails_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["hostedType"] = hostedType_
	params["oActivityServiceDetails"] = oActivityServiceDetails_
	return APIInterface(coupon.config, "eleme.activity.coupon.rehostShop", params)
}

// 定向赠红包(单店红包)
// shopId 店铺id
// targetList 目标列表
// targetListType 目标类型
// targetCouponDetail 定向赠红包模板细节
func (coupon *Coupon) PresentTargetCoupons(shopId_ int64, targetList_ interface{}, targetListType_ interface{}, targetCouponDetail_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["targetList"] = targetList_
	params["targetListType"] = targetListType_
	params["targetCouponDetail"] = targetCouponDetail_
	return APIInterface(coupon.config, "eleme.activity.coupon.presentTargetCoupons", params)
}

// 定向赠通用红包
// chainId 连锁店id
// targetList 目标列表
// targetListType 目标类型
// commonTargetCouponDetail 通用定向赠红包模板细节
func (coupon *Coupon) PresentCommonTargetCoupons(chainId_ int64, targetList_ interface{}, targetListType_ interface{}, commonTargetCouponDetail_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["chainId"] = chainId_
	params["targetList"] = targetList_
	params["targetListType"] = targetListType_
	params["commonTargetCouponDetail"] = commonTargetCouponDetail_
	return APIInterface(coupon.config, "eleme.activity.coupon.presentCommonTargetCoupons", params)
}

// 分页查询店铺的定向赠红包信息
// targetCouponQueryRequest 定向赠红包查询入参对象
func (coupon *Coupon) QueryTargetCouponInfo(targetCouponQueryRequest_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["targetCouponQueryRequest"] = targetCouponQueryRequest_
	return APIInterface(coupon.config, "eleme.activity.coupon.queryTargetCouponInfo", params)
}

// 通过店铺Id查询该店铺被邀约的美食活动
// shopId 店铺Id
func (food *Food) QueryInvitedFoodActivities(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(food.config, "eleme.activity.food.queryInvitedFoodActivities", params)
}

// 报名美食活动
// activityId 活动Id
// activityApplyInfo 活动报名信息
func (food *Food) ApplyFoodActivity(activityId_ int64, activityApplyInfo_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["activityId"] = activityId_
	params["activityApplyInfo"] = activityApplyInfo_
	return APIInterface(food.config, "eleme.activity.food.applyFoodActivity", params)
}

// 通过店铺Id和活动Id分页查询店铺已报名的美食活动
// activityId 活动Id
// shopId 店铺Id
// pageNo 页码
// pageSize 每页数量
func (food *Food) QueryFoodActivities(activityId_ int64, shopId_ int64, pageNo_ int64, pageSize_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["activityId"] = activityId_
	params["shopId"] = shopId_
	params["pageNo"] = pageNo_
	params["pageSize"] = pageSize_
	return APIInterface(food.config, "eleme.activity.food.queryFoodActivities", params)
}

// 修改美食活动的菜品库存
// activityId 活动Id
// shopId 店铺Id
// itemId 菜品Id
// stock 库存
func (food *Food) UpdateFoodActivityItemStock(activityId_ int64, shopId_ int64, itemId_ int64, stock_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["activityId"] = activityId_
	params["shopId"] = shopId_
	params["itemId"] = itemId_
	params["stock"] = stock_
	return APIInterface(food.config, "eleme.activity.food.updateFoodActivityItemStock", params)
}

// 取消参与了美食活动的菜品
// activityId 活动Id
// shopId 店铺Id
// itemId 菜品Id
func (food *Food) OfflineFoodActivityItem(activityId_ int64, shopId_ int64, itemId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["activityId"] = activityId_
	params["shopId"] = shopId_
	params["itemId"] = itemId_
	return APIInterface(food.config, "eleme.activity.food.offlineFoodActivityItem", params)
}

// 作废店铺与美食活动的关联关系
// activityId 活动Id
// shopId 店铺Id
func (food *Food) UnbindFoodActivity(activityId_ int64, shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["activityId"] = activityId_
	params["shopId"] = shopId_
	return APIInterface(food.config, "eleme.activity.food.unbindFoodActivity", params)
}

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

