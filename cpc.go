// CPC竞价服务 
package elemeOpenApi

// 查询余额
// shopId 店铺ID
func (cpc *Cpc) GetAllBalance(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(cpc.config, "eleme.cpc.getAllBalance", params)
}

// 确认店铺两证是否齐全
// shopId 店铺ID
func (cpc *Cpc) CheckShopCertification(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(cpc.config, "eleme.cpc.checkShopCertification", params)
}

// 查询推广信息
// shopId 店铺ID
func (cpc *Cpc) GetWagerInformation(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(cpc.config, "eleme.cpc.getWagerInformation", params)
}

// 根据出价查询预估信息
// shopId 店铺ID
// bid CPC出价
func (cpc *Cpc) GetWagerEstimate(shopId_ int64, bid_ float64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["bid"] = bid_
	return APIInterface(cpc.config, "eleme.cpc.getWagerEstimate", params)
}

// 查询推荐价格、预估信息
// shopId 店铺ID
func (cpc *Cpc) GetSuggestWagerInfo(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(cpc.config, "eleme.cpc.getSuggestWagerInfo", params)
}

// 查询推广修改剩余次数
// shopId 店铺ID
func (cpc *Cpc) GetResidueDegree(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(cpc.config, "eleme.cpc.getResidueDegree", params)
}

// 设置推广状态
// shopId 店铺ID
// status 推广状态
func (cpc *Cpc) UpdateWagerStatus(shopId_ int64, status_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["status"] = status_
	return APIInterface(cpc.config, "eleme.cpc.updateWagerStatus", params)
}

// 设置推广出价
// shopId 店铺ID
// bid 要设置的出价(元)
func (cpc *Cpc) SetWagerGrade(shopId_ int64, bid_ float64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["bid"] = bid_
	return APIInterface(cpc.config, "eleme.cpc.setWagerGrade", params)
}

// 设置预算
// shopId 店铺ID
// budget 要设置的预算(元)
func (cpc *Cpc) SetWagerBudget(shopId_ int64, budget_ float64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["budget"] = budget_
	return APIInterface(cpc.config, "eleme.cpc.setWagerBudget", params)
}

// 更新自动投放状态
// shopId 店铺ID
// autoStatus 操作状态
// launchHours 小时集合
func (cpc *Cpc) UpdateAutoStatus(shopId_ int64, autoStatus_ interface{}, launchHours_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["autoStatus"] = autoStatus_
	params["launchHours"] = launchHours_
	return APIInterface(cpc.config, "eleme.cpc.updateAutoStatus", params)
}

// 设置推广速率
// shopId 店铺ID
// wagerSpeedMode 速率类型
func (cpc *Cpc) SetWagerSpeed(shopId_ int64, wagerSpeedMode_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["wagerSpeedMode"] = wagerSpeedMode_
	return APIInterface(cpc.config, "eleme.cpc.setWagerSpeed", params)
}

// 获取竞价推广实时排名
// shopId 店铺ID
func (cpc *Cpc) GetActualRanking(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(cpc.config, "eleme.cpc.getActualRanking", params)
}

// 查询推广效果数据
// shopId 店铺ID
// beginDate 开始时间
// endDate 结束时间
func (cpc *Cpc) GetUVSummary(shopId_ int64, beginDate_ string, endDate_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["beginDate"] = beginDate_
	params["endDate"] = endDate_
	return APIInterface(cpc.config, "eleme.cpc.getUVSummary", params)
}

// 查询推广点击分布信息
// shopId 店铺ID
// beginDate 开始时间
// endDate 结束时间
func (cpc *Cpc) GetRankAndCostInfo(shopId_ int64, beginDate_ string, endDate_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["beginDate"] = beginDate_
	params["endDate"] = endDate_
	return APIInterface(cpc.config, "eleme.cpc.getRankAndCostInfo", params)
}

// 获取推广活跃顾客的点击结构
// shopId 店铺ID
// date 时间
func (cpc *Cpc) GetUserDistribution(shopId_ int64, date_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["date"] = date_
	return APIInterface(cpc.config, "eleme.cpc.getUserDistribution", params)
}

