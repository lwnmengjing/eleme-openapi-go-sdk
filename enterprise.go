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

// 更新企业订餐店铺在线点餐启用状态
// enableRequest 门店启用请求
func (enterprise *Enterprise) UpdateArrivalShopOnlineOrderEnable(enableRequest_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["enableRequest"] = enableRequest_
	return APIInterface(enterprise.config, "eleme.enterprise.updateArrivalShopOnlineOrderEnable", params)
}

// 获取饿了么企餐用户认证
// req 饿了么企餐用户
func (enterprise *Enterprise) GetUserAuthentication(req_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["req"] = req_
	return APIInterface(enterprise.config, "eleme.enterprise.getUserAuthentication", params)
}

// ISV创建饿了么订单，获取订单编号
// createReq 创单参数
func (enterprise *Enterprise) CreateOnlineOrder(createReq_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["createReq"] = createReq_
	return APIInterface(enterprise.config, "eleme.enterprise.createOnlineOrder", params)
}

// ISV订单详情同步
// detailReq 订单详情
func (enterprise *Enterprise) PushOrderDetail(detailReq_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["detailReq"] = detailReq_
	return APIInterface(enterprise.config, "eleme.enterprise.pushOrderDetail", params)
}

// 加载企业订餐买单页面
// req 买单参数
func (enterprise *Enterprise) LoadPaymentPage(req_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["req"] = req_
	return APIInterface(enterprise.config, "eleme.enterprise.loadPaymentPage", params)
}

// 企业付订单查询接口
// orderReqDto 订单请求
func (enterprise *Enterprise) GetEntArrivalOrderDetail(orderReqDto_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["orderReqDto"] = orderReqDto_
	return APIInterface(enterprise.config, "eleme.enterprise.getEntArrivalOrderDetail", params)
}

