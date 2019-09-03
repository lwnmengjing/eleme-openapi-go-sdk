// 经营体检 
package elemeOpenApi

// 根据商户ID查询商户经营体检信息
// shopId 店铺ID
// date 体检日期(最多查到7天内的体检数据)
func (diagnosis *Diagnosis) GetShopDiagnosis(shopId_ int64, date_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["date"] = date_
	return APIInterface(diagnosis.config, "eleme.diagnosis.getShopDiagnosis", params)
}

// 根据多个商户ID批量查询商户经营体检信息
// shopIds 店铺ID集合
// date 体检日期(最多查到7天内的体检数据)
func (diagnosis *Diagnosis) GetShopDiagnosisList(shopIds_ interface{}, date_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopIds"] = shopIds_
	params["date"] = date_
	return APIInterface(diagnosis.config, "eleme.diagnosis.getShopDiagnosisList", params)
}

