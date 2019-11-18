// 商户信用分服务 
package elemeOpenApi

// 连锁店根据商户ID集合批量查询商户信用分信息
// shopIds 商户ID集合
func (chain *Chain) BatchQueryShopCreditScores(shopIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopIds"] = shopIds_
	return APIInterface(chain.config, "eleme.shopCreditScore.chain.batchQueryShopCreditScores", params)
}

// 连锁店根据商户ID集合批量查询店铺权益规则
// shopIds 商户ID集合
func (chain *Chain) BatchQueryShopEquityRules(shopIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopIds"] = shopIds_
	return APIInterface(chain.config, "eleme.shopCreditScore.chain.batchQueryShopEquityRules", params)
}

// 连锁店根据商户ID集合批量查询店铺扣罚规则
// shopIds 商户ID集合
func (chain *Chain) BatchQueryShopPunishRules(shopIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopIds"] = shopIds_
	return APIInterface(chain.config, "eleme.shopCreditScore.chain.batchQueryShopPunishRules", params)
}

// 连锁店根据商户ID集合批量查询查询商户信用分变更记录
// shopIds 商户ID集合
func (chain *Chain) BatchQueryShopCreditScoreRecords(shopIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopIds"] = shopIds_
	return APIInterface(chain.config, "eleme.shopCreditScore.chain.batchQueryShopCreditScoreRecords", params)
}

// 根据商户ID查询商户信用分信息
// shopId 商户ID
func (single *Single) GetShopCreditScore(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(single.config, "eleme.shopCreditScore.single.getShopCreditScore", params)
}

// 根据商户ID查询店铺权益规则
// shopId 商户ID
func (single *Single) GetShopEquityRules(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(single.config, "eleme.shopCreditScore.single.getShopEquityRules", params)
}

// 根据商户ID查询店铺扣罚规则
// shopId 商户ID
func (single *Single) GetShopPunishRules(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(single.config, "eleme.shopCreditScore.single.getShopPunishRules", params)
}

// 查询商户信用分变更记录
// shopId 商户ID
func (single *Single) GetShopCreditScoreRecord(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(single.config, "eleme.shopCreditScore.single.getShopCreditScoreRecord", params)
}

