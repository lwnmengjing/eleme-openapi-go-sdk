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

