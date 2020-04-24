// 企业订餐商户服务 
package elemeOpenApi

// 关联企业订餐到店订单
// relateReqDto 订单关联请求
func (enterprise *Enterprise) UpdateEntArrivalOrderRelate(relateReqDto_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["relateReqDto"] = relateReqDto_
	return APIInterface(enterprise.config, "eleme.enterprise.updateEntArrivalOrderRelate", params)
}

// 更新企业订餐店铺订单关联启用状态
// enableRequest 门店启用请求
func (enterprise *Enterprise) UpdateEntArrivalShopEnable(enableRequest_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["enableRequest"] = enableRequest_
	return APIInterface(enterprise.config, "eleme.enterprise.updateEntArrivalShopEnable", params)
}

