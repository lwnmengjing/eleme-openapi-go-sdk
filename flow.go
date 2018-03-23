// 餐厅入口流量服务 
package elemeOpenApi

// 根据时间段获取餐厅流量入口数据
// request 餐厅入口流量查询条件
func (flow *Flow) GetEntryFlowStatsData(request_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["request"] = request_
	return APIInterface(flow.config, "eleme.flow.getEntryFlowStatsData", params)
}

