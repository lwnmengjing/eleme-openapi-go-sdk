// 门店装修服务 
package elemeOpenApi

// 连锁店总店创建橱窗
// window 新增的橱窗信息和其关联连锁店子店I和子店商品
func (windows *Windows) CreateWindow(window_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["window"] = window_
	return APIInterface(windows.config, "eleme.decoration.windows.createWindow", params)
}

// 连锁店总店修改橱窗
// window 修改的橱窗信息和其关联连锁店子店ID和子店商品
func (windows *Windows) UpdateWindow(window_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["window"] = window_
	return APIInterface(windows.config, "eleme.decoration.windows.updateWindow", params)
}

// 连锁店总店删除橱窗
// window 删除橱窗信息
func (windows *Windows) DeleteWindow(window_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["window"] = window_
	return APIInterface(windows.config, "eleme.decoration.windows.deleteWindow", params)
}

// 连锁店总店对多个橱窗进行排序
// window 橱窗排序信息
func (windows *Windows) OrderWindow(window_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["window"] = window_
	return APIInterface(windows.config, "eleme.decoration.windows.orderWindow", params)
}

// 连锁店总店根据橱窗ID获取橱窗详情
// burstWindowId 橱窗ID
func (windows *Windows) GetWindowById(burstWindowId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["burstWindowId"] = burstWindowId_
	return APIInterface(windows.config, "eleme.decoration.windows.getWindowById", params)
}

// 连锁店总店获取橱窗信息集合
// operateShopId 连锁店总店ID
func (windows *Windows) GetWindowByShopId(operateShopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["operateShopId"] = operateShopId_
	return APIInterface(windows.config, "eleme.decoration.windows.getWindowByShopId", params)
}

// 连锁店总店创建招贴
// sign 招贴信息和其关联连锁店子店ID
func (sign *Sign) CreateSign(sign_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["sign"] = sign_
	return APIInterface(sign.config, "eleme.decoration.sign.createSign", params)
}

// 连锁店总店修改招贴
// signId 招贴ID
// sign 招贴信息和其关联连锁店子店ID
func (sign *Sign) UpdateSign(signId_ int64, sign_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["signId"] = signId_
	params["sign"] = sign_
	return APIInterface(sign.config, "eleme.decoration.sign.updateSign", params)
}

// 连锁店总店作废招贴
// signId 招贴ID
// operateShopId 连锁店总店ID
func (sign *Sign) InvalidSign(signId_ int64, operateShopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["signId"] = signId_
	params["operateShopId"] = operateShopId_
	return APIInterface(sign.config, "eleme.decoration.sign.invalidSign", params)
}

// 连锁店总店获取历史上传过的招贴图片
// sign 查询条件
func (sign *Sign) GetSignHistoryImage(sign_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["sign"] = sign_
	return APIInterface(sign.config, "eleme.decoration.sign.getSignHistoryImage", params)
}

// 连锁店总店分页查询店铺招贴
// sign 查询条件
func (sign *Sign) QuerySign(sign_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["sign"] = sign_
	return APIInterface(sign.config, "eleme.decoration.sign.querySign", params)
}

// 连锁店总店根据招贴ID查询店铺招贴详情
// signId 招贴ID
func (sign *Sign) GetSignById(signId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["signId"] = signId_
	return APIInterface(sign.config, "eleme.decoration.sign.getSignById", params)
}

// 连锁店总店创建海报接口
// poster 海报信息和其关联连锁店子店I和子店商品
func (poster *Poster) CreatePoster(poster_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["poster"] = poster_
	return APIInterface(poster.config, "eleme.decoration.poster.createPoster", params)
}

// 连锁店总店修改海报
// posterId 海报ID
// poster 海报信息和其关联连锁店子店ID和子店商品
func (poster *Poster) UpdatePoster(posterId_ int64, poster_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["posterId"] = posterId_
	params["poster"] = poster_
	return APIInterface(poster.config, "eleme.decoration.poster.updatePoster", params)
}

// 连锁店总店作废海报
// poster 作废海报信息
func (poster *Poster) InvalidPoster(poster_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["poster"] = poster_
	return APIInterface(poster.config, "eleme.decoration.poster.invalidPoster", params)
}

// 连锁店总店根据海报ID获取海报详情
// posterId 海报ID
// operateShopId 连锁店总店ID
func (poster *Poster) GetPosterDetailById(posterId_ int64, operateShopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["posterId"] = posterId_
	params["operateShopId"] = operateShopId_
	return APIInterface(poster.config, "eleme.decoration.poster.getPosterDetailById", params)
}

// 连锁店总店根据条件查询海报信息集合
// poster 查询海报条件参数
func (poster *Poster) QueryPoster(poster_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["poster"] = poster_
	return APIInterface(poster.config, "eleme.decoration.poster.queryPoster", params)
}

// 连锁店总店获取历史上传过的海报图片
// operateShopId 连锁店总店ID
func (poster *Poster) GetPosterHistoryImage(operateShopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["operateShopId"] = operateShopId_
	return APIInterface(poster.config, "eleme.decoration.poster.getPosterHistoryImage", params)
}

// 连锁店总店新增品牌故事
// story 品牌故事信息和其关联连锁店子店ID
func (story *Story) CreateBrandStory(story_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["story"] = story_
	return APIInterface(story.config, "eleme.decoration.story.createBrandStory", params)
}

// 连锁店总店更新品牌故事
// brandStoryId 品牌故事ID
// story 品牌故事信息和其关联连锁店子店ID
func (story *Story) UpdateBrandStory(brandStoryId_ int64, story_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["brandStoryId"] = brandStoryId_
	params["story"] = story_
	return APIInterface(story.config, "eleme.decoration.story.updateBrandStory", params)
}

// 连锁店总店删除品牌故事
// brandStoryId 品牌故事ID
// operateShopId 连锁店总店店铺ID
func (story *Story) DeleteBrandStory(brandStoryId_ int64, operateShopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["brandStoryId"] = brandStoryId_
	params["operateShopId"] = operateShopId_
	return APIInterface(story.config, "eleme.decoration.story.deleteBrandStory", params)
}

// 连锁店总店查询当前设置的品牌故事信息
// brandStoryId 品牌故事ID
func (story *Story) GetBrandStoryById(brandStoryId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["brandStoryId"] = brandStoryId_
	return APIInterface(story.config, "eleme.decoration.story.getBrandStoryById", params)
}

// 上传图片
// image 文件内容base64编码值
func (image *Image) Upload(image_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["image"] = image_
	return APIInterface(image.config, "eleme.decoration.image.upload", params)
}

// 根据图片HASH值获取图片信息
// hash 图片HASH值
func (image *Image) GetImage(hash_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["hash"] = hash_
	return APIInterface(image.config, "eleme.decoration.image.getImage", params)
}

