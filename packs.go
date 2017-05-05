// 签约服务 
package elemeOpenApi

// 查询店铺当前生效合同类型
// shopId 店铺id
func (packs *Packs) GetEffectServicePackContract(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(packs.config, "eleme.packs.getEffectServicePackContract", params)
}

