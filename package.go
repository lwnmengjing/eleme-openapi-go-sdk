// 店铺服务包服务 
package elemeOpenApi

// 查询店铺当前生效合同类型
// shopId 店铺id
func (package *Package) GetEffectServicePackContract(shopId_ int) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(package.config, "eleme.package.getEffectServicePackContract", params)
}

