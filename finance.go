// 金融服务 
package elemeOpenApi

// 查询商户余额,返回可用余额和总余额
// shopId 饿了么店铺id
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

// 查询总店账单(旧接口，即将下线)
// shopId 饿了么总店店铺id
// query 查询条件
func (finance *Finance) QueryHeadBills(shopId_ int64, query_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["query"] = query_
	return APIInterface(finance.config, "eleme.finance.queryHeadBills", params)
}

// 查询总店账单(新接口)
// shopId 饿了么总店店铺id
// query 查询条件
func (finance *Finance) QueryHeadBillsNew(shopId_ int64, query_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["query"] = query_
	return APIInterface(finance.config, "eleme.finance.queryHeadBillsNew", params)
}

// 查询总店订单(旧接口，即将下线)
// shopId 饿了么总店店铺id
// query 查询条件
func (finance *Finance) QueryHeadOrders(shopId_ int64, query_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["query"] = query_
	return APIInterface(finance.config, "eleme.finance.queryHeadOrders", params)
}

// 查询总店订单(新接口)
// shopId 饿了么总店店铺id
// query 查询条件
func (finance *Finance) QueryHeadOrdersNew(shopId_ int64, query_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["query"] = query_
	return APIInterface(finance.config, "eleme.finance.queryHeadOrdersNew", params)
}

// 查询分店账单(旧接口，即将下线)
// shopId 饿了么分店店铺id
// query 查询条件
func (finance *Finance) QueryBranchBills(shopId_ int64, query_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["query"] = query_
	return APIInterface(finance.config, "eleme.finance.queryBranchBills", params)
}

// 查询分店账单(新接口)
// shopId 饿了么分店店铺id
// query 查询条件
func (finance *Finance) QueryBranchBillsNew(shopId_ int64, query_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["query"] = query_
	return APIInterface(finance.config, "eleme.finance.queryBranchBillsNew", params)
}

// 查询分店订单(旧接口，即将下线)
// shopId 饿了么分店店铺id
// query 查询条件
func (finance *Finance) QueryBranchOrders(shopId_ int64, query_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["query"] = query_
	return APIInterface(finance.config, "eleme.finance.queryBranchOrders", params)
}

// 查询分店订单(新接口)
// shopId 饿了么分店店铺id
// query 查询条件
func (finance *Finance) QueryBranchOrdersNew(shopId_ int64, query_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["query"] = query_
	return APIInterface(finance.config, "eleme.finance.queryBranchOrdersNew", params)
}

// 查询订单(旧接口，即将下线)
// shopId 饿了么店铺id
// orderId 订单id
func (finance *Finance) GetOrder(shopId_ int64, orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["orderId"] = orderId_
	return APIInterface(finance.config, "eleme.finance.getOrder", params)
}

// 查询订单(新接口)
// shopId 饿了么店铺id
// orderId 订单id
func (finance *Finance) GetOrderNew(shopId_ int64, orderId_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["orderId"] = orderId_
	return APIInterface(finance.config, "eleme.finance.getOrderNew", params)
}

// 查询返现汇总信息账单
// shopId 饿了么分店、单店、总店店铺id
// query 查询条件
func (finance *Finance) QueryAllowanceBills(shopId_ int64, query_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["query"] = query_
	return APIInterface(finance.config, "eleme.finance.queryAllowanceBills", params)
}

// 查询返现每日详单
// shopId 饿了么分店、单店、总店店铺id
// query 查询条件
func (finance *Finance) QueryAllowanceBillDetail(shopId_ int64, query_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["query"] = query_
	return APIInterface(finance.config, "eleme.finance.queryAllowanceBillDetail", params)
}

