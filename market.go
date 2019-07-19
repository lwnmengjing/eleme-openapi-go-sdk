// 服务市场服务 
package elemeOpenApi

// 同步某一段时间内的服务市场消息
// start 开始时间
// end 结束时间
// offset 消息偏移量
// limit 查询消息数
func (market *Market) SyncMarketMessages(start_ interface{}, end_ interface{}, offset_ int, limit_ int) (interface{}, error) {
	params := make(map[string]interface{})
	params["start"] = start_
	params["end"] = end_
	params["offset"] = offset_
	params["limit"] = limit_
	return APIInterface(market.config, "eleme.market.syncMarketMessages", params)
}

// 创建内购项目订单
// request 创建订单请求信息
func (market *Market) CreateOrder(request_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["request"] = request_
	return APIInterface(market.config, "eleme.market.createOrder", params)
}

// 查询服务市场订单
// orderNo 服务市场订单编号
func (market *Market) QueryOrder(orderNo_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderNo"] = orderNo_
	return APIInterface(market.config, "eleme.market.queryOrder", params)
}

// 服务市场确认订单
// orderNo 服务市场订单编号
func (market *Market) ConfirmOrder(orderNo_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderNo"] = orderNo_
	return APIInterface(market.config, "eleme.market.confirmOrder", params)
}

