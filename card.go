// 商户会员卡服务 
package elemeOpenApi

// 上传图片
// imageBase64 上传图片
func (card *Card) UploadImage(imageBase64_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["imageBase64"] = imageBase64_
	return APIInterface(card.config, "eleme.card.uploadImage", params)
}

// 创建模板
// templateInfo 模板信息
func (card *Card) CreateTemplate(templateInfo_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["templateInfo"] = templateInfo_
	return APIInterface(card.config, "eleme.card.createTemplate", params)
}

// 查询模板信息
// templateId 模板id列表
func (card *Card) MgetTemplateInfo(templateId_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["templateId"] = templateId_
	return APIInterface(card.config, "eleme.card.mgetTemplateInfo", params)
}

// 更新模板信息
// templateId 模板id
// templateInfo 模板更新信息
func (card *Card) UpdateTemplate(templateId_ string, templateInfo_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["templateId"] = templateId_
	params["templateInfo"] = templateInfo_
	return APIInterface(card.config, "eleme.card.updateTemplate", params)
}

// 查询模板应用的店铺
// templateId 模板id列表
func (card *Card) MgetShopIdsByTemplateIds(templateId_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["templateId"] = templateId_
	return APIInterface(card.config, "eleme.card.mgetShopIdsByTemplateIds", params)
}

// 应用模板
// templateId 模板id
// shopIds 店铺列表
func (card *Card) ApplyTemplate(templateId_ string, shopIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["templateId"] = templateId_
	params["shopIds"] = shopIds_
	return APIInterface(card.config, "eleme.card.applyTemplate", params)
}

// 开卡
// templateId 模板ID
// cardUserInfo 会员用户信息
// cardAccountInfo 会员账户信息
func (card *Card) OpenCard(templateId_ string, cardUserInfo_ interface{}, cardAccountInfo_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["templateId"] = templateId_
	params["cardUserInfo"] = cardUserInfo_
	params["cardAccountInfo"] = cardAccountInfo_
	return APIInterface(card.config, "eleme.card.openCard", params)
}

// 更新会员信息
// cardUserInfo 用户基本信息
// cardAccountInfo 用户账户信息
func (card *Card) UpdateUserInfo(cardUserInfo_ interface{}, cardAccountInfo_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["cardUserInfo"] = cardUserInfo_
	params["cardAccountInfo"] = cardAccountInfo_
	return APIInterface(card.config, "eleme.card.updateUserInfo", params)
}

// 根据userToken获取用户信息(该接口不再使用)
// userToken userToken有效期10分钟.饿了么app上跳转到外部H5页面https://www.abc.com?accessToken=c8cea843-1fb5-473f-bb10-a9d2aa239c39,其中accessToken为userToken,用其作为该接口的入参获取到用户信息
func (card *Card) GetUserByToken(userToken_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["userToken"] = userToken_
	return APIInterface(card.config, "eleme.card.getUserByToken", params)
}

