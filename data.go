// 商户数据服务 
package elemeOpenApi

// 查询指定时间段内单店营业数据汇总（历史数据）
// shopId 店铺Id
// startTime 查询起始日期
// endTime 查询结束日期
func (single *Single) GetRestaurantSaleDetail(shopId_ int64, startTime_ string, endTime_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["startTime"] = startTime_
	params["endTime"] = endTime_
	return APIInterface(single.config, "eleme.data.single.getRestaurantSaleDetail", params)
}

// 获取单店今日实时的营业数据汇总
// shopId 店铺Id
func (single *Single) GetRestaurantRealTimeSaleDetail(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(single.config, "eleme.data.single.getRestaurantRealTimeSaleDetail", params)
}

// 查询指定时间段内单店相关营业数据指标增长率
// shopId 店铺Id
// startTime 查询起始日期
// endTime 查询结束日期
func (single *Single) GetRestaurantSaleRatio(shopId_ int64, startTime_ string, endTime_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["startTime"] = startTime_
	params["endTime"] = endTime_
	return APIInterface(single.config, "eleme.data.single.getRestaurantSaleRatio", params)
}

// 查询指定时间段内连锁店营业数据汇总(历史数据)
// shopIds 连锁子店Id
// startTime 查询起始日期
// endTime 查询结束日期
func (chain *Chain) GetChainRestaurantSaleDetail(shopIds_ interface{}, startTime_ string, endTime_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopIds"] = shopIds_
	params["startTime"] = startTime_
	params["endTime"] = endTime_
	return APIInterface(chain.config, "eleme.data.chain.getChainRestaurantSaleDetail", params)
}

// 获取连锁店今日实时的营业数据汇总
// shopIds 连锁子店Id
func (chain *Chain) GetChainRealTimeSaleDetail(shopIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopIds"] = shopIds_
	return APIInterface(chain.config, "eleme.data.chain.getChainRealTimeSaleDetail", params)
}

// 查询指定时间段内连锁店相关营业数据指标增长率
// shopIds 连锁子店Id
// startTime 查询起始日期
// endTime 查询结束日期
func (chain *Chain) GetChainRestaurantSaleRatio(shopIds_ interface{}, startTime_ string, endTime_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopIds"] = shopIds_
	params["startTime"] = startTime_
	params["endTime"] = endTime_
	return APIInterface(chain.config, "eleme.data.chain.getChainRestaurantSaleRatio", params)
}

