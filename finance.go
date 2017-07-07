// 金融服务 
package elemeOpenApi

// 查询商户余额,返回可用余额和总余额
// shopId 饿了么店铺Id
func (finance *Finance) QueryBalance(shopId_ int) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(finance.config, "eleme.finance.queryBalance", params)
}

// 查询余额流水,有流水改动的交易
// request 查询条件
func (finance *Finance) QueryBalanceLog(request_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["request"] = request_
	return APIInterface(finance.config, "eleme.finance.queryBalanceLog", params)
}

