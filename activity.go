// 活动服务 
package elemeOpenApi

// 创建减配送费活动
// createInfo 创建减配送费活动的结构体
// shopId 店铺Id
func (shippingfee *Shippingfee) CreateShippingFeeActivity(createInfo_ interface{}, shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["createInfo"] = createInfo_
	params["shopId"] = shopId_
	return APIInterface(shippingfee.config, "eleme.activity.shippingFee.createShippingFeeActivity", params)
}

// 作废减配送费活动
// activityId 活动Id
// shopId 店铺Id
// comment 作废原因
func (shippingfee *Shippingfee) InvalidShippingFeeActivity(activityId_ int64, shopId_ int64, comment_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["activityId"] = activityId_
	params["shopId"] = shopId_
	params["comment"] = comment_
	return APIInterface(shippingfee.config, "eleme.activity.shippingFee.invalidShippingFeeActivity", params)
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

// 定向赠通用商品券
// chainId 连锁店id
// targetList 目标列表
// targetListType 目标类型
// commonTargetSkuCouponDetail 通用定向赠连锁商品券模板细节
func (coupon *Coupon) PresentCommonTargetSkuCoupons(chainId_ int64, targetList_ interface{}, targetListType_ interface{}, commonTargetSkuCouponDetail_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["chainId"] = chainId_
	params["targetList"] = targetList_
	params["targetListType"] = targetListType_
	params["commonTargetSkuCouponDetail"] = commonTargetSkuCouponDetail_
	return APIInterface(coupon.config, "eleme.activity.coupon.presentCommonTargetSkuCoupons", params)
}

// 定向赠连锁通用商品券
// chainId 连锁店id
// targetList 目标列表
// targetListType 目标类型
// chainSkuCouponDetail 通用定向赠连锁商品券模板细节
func (coupon *Coupon) PresentChainSkuCoupons(chainId_ int64, targetList_ interface{}, targetListType_ interface{}, chainSkuCouponDetail_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["chainId"] = chainId_
	params["targetList"] = targetList_
	params["targetListType"] = targetListType_
	params["chainSkuCouponDetail"] = chainSkuCouponDetail_
	return APIInterface(coupon.config, "eleme.activity.coupon.presentChainSkuCoupons", params)
}

// 定向赠指定商品券
// targetList 目标列表
// targetListType 目标类型
// skuCouponDetail 商品券模板细节
func (coupon *Coupon) PresentSkuCoupons(targetList_ interface{}, targetListType_ interface{}, skuCouponDetail_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["targetList"] = targetList_
	params["targetListType"] = targetListType_
	params["skuCouponDetail"] = skuCouponDetail_
	return APIInterface(coupon.config, "eleme.activity.coupon.presentSkuCoupons", params)
}

// 券状态变更
// criteria 券状态修改对象
// type 操作类型
func (coupon *Coupon) UpdateCouponStatus(criteria_ interface{}, type_ int) (interface{}, error) {
	params := make(map[string]interface{})
	params["criteria"] = criteria_
	params["type"] = type_
	return APIInterface(coupon.config, "eleme.activity.coupon.updateCouponStatus", params)
}

// 查询订单内营销相关数据
// orderId 饿了么订单Id
func (marketing *Marketing) QueryOrderSubsidy(orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderId"] = orderId_
	return APIInterface(marketing.config, "eleme.activity.marketing.queryOrderSubsidy", params)
}

// 创建并绑定连锁店特价活动
// activity 活动创建信息
// chainId 连锁店Id
// shopApplyInfo  绑定的商品信息
func (skuchain *Skuchain) CreateAndParticipateChainPriceActivity(activity_ interface{}, chainId_ int64, shopApplyInfo_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["activity"] = activity_
	params["chainId"] = chainId_
	params["shopApplyInfo"] = shopApplyInfo_
	return APIInterface(skuchain.config, "eleme.activity.skuchain.createAndParticipateChainPriceActivity", params)
}

// 根据活动Id和店铺Id和商品规格Id，作废参与关系
// activityId 活动Id
// shopId 店铺Id
// specId 商品规格Id
// comment 作废原因
func (skuchain *Skuchain) InValidSkuActivityById(activityId_ int64, shopId_ int64, specId_ int64, comment_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["activityId"] = activityId_
	params["shopId"] = shopId_
	params["specId"] = specId_
	params["comment"] = comment_
	return APIInterface(skuchain.config, "eleme.activity.skuchain.inValidSkuActivityById", params)
}

