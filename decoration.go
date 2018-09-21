// 门店装修服务 
package elemeOpenApi

// 创建招贴
// sign 招贴信息和其关联门店ID集合
func (sign *Sign) CreateSign(sign_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["sign"] = sign_
	return APIInterface(sign.config, "eleme.decoration.sign.createSign", params)
}

// 修改招贴
// signId 招贴ID
// sign 招贴信息和其关联门店ID
func (sign *Sign) UpdateSign(signId_ int64, sign_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["signId"] = signId_
	params["sign"] = sign_
	return APIInterface(sign.config, "eleme.decoration.sign.updateSign", params)
}

// 作废招贴
// signId 招贴ID
func (sign *Sign) InvalidSign(signId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["signId"] = signId_
	return APIInterface(sign.config, "eleme.decoration.sign.invalidSign", params)
}

// 获取历史上传过的招贴图片
// sign 查询条件
func (sign *Sign) GetSignHistoryImage(sign_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["sign"] = sign_
	return APIInterface(sign.config, "eleme.decoration.sign.getSignHistoryImage", params)
}

// 查询有效招贴集合
func (sign *Sign) QuerySign() (interface{}, error) {
	params := make(map[string]interface{})
	return APIInterface(sign.config, "eleme.decoration.sign.querySign", params)
}

// 根据招贴ID查询店铺招贴详情
// signId 招贴ID
func (sign *Sign) GetSignById(signId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["signId"] = signId_
	return APIInterface(sign.config, "eleme.decoration.sign.getSignById", params)
}

// 新增品牌故事
// story 品牌故事信息和其关联连锁店子店ID
func (story *Story) CreateBrandStory(story_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["story"] = story_
	return APIInterface(story.config, "eleme.decoration.story.createBrandStory", params)
}

// 更新品牌故事
// brandStoryId 品牌故事ID
// story 品牌故事信息和其关联连锁店子店ID
func (story *Story) UpdateBrandStory(brandStoryId_ int64, story_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["brandStoryId"] = brandStoryId_
	params["story"] = story_
	return APIInterface(story.config, "eleme.decoration.story.updateBrandStory", params)
}

// 删除品牌故事
// brandStoryId 品牌故事ID
func (story *Story) DeleteBrandStory(brandStoryId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["brandStoryId"] = brandStoryId_
	return APIInterface(story.config, "eleme.decoration.story.deleteBrandStory", params)
}

// 查询品牌故事列表
func (story *Story) QueryBrandStoryList() (interface{}, error) {
	params := make(map[string]interface{})
	return APIInterface(story.config, "eleme.decoration.story.queryBrandStoryList", params)
}

// 查询当前设置的品牌故事信息
// brandStoryId 品牌故事ID
func (story *Story) GetBrandStoryById(brandStoryId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["brandStoryId"] = brandStoryId_
	return APIInterface(story.config, "eleme.decoration.story.getBrandStoryById", params)
}

// 保存精准分类
// category 精准分类信息
func (accuratecategory *Accuratecategory) SaveCategory(category_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["category"] = category_
	return APIInterface(accuratecategory.config, "eleme.decoration.accurateCategory.saveCategory", params)
}

// 根据门店ID获取精准分类
// category 查询参数
func (accuratecategory *Accuratecategory) GetAccurateCategory(category_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["category"] = category_
	return APIInterface(accuratecategory.config, "eleme.decoration.accurateCategory.getAccurateCategory", params)
}

// 查询精准分类
// category 查询参数
func (accuratecategory *Accuratecategory) QueryAccurateCategoryList(category_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["category"] = category_
	return APIInterface(accuratecategory.config, "eleme.decoration.accurateCategory.queryAccurateCategoryList", params)
}

// 创建多橱窗
// window 新增的橱窗信息和其关联门店ID和关联商品
func (windows *Windows) CreateWindow(window_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["window"] = window_
	return APIInterface(windows.config, "eleme.decoration.windows.createWindow", params)
}

// 修改橱窗
// window 修改的橱窗信息和其关联门店ID和门店商品
func (windows *Windows) UpdateWindow(window_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["window"] = window_
	return APIInterface(windows.config, "eleme.decoration.windows.updateWindow", params)
}

// 删除橱窗
// window 删除橱窗信息
func (windows *Windows) DeleteWindow(window_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["window"] = window_
	return APIInterface(windows.config, "eleme.decoration.windows.deleteWindow", params)
}

// 对多个橱窗进行排序
// window 橱窗排序信息
func (windows *Windows) OrderWindow(window_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["window"] = window_
	return APIInterface(windows.config, "eleme.decoration.windows.orderWindow", params)
}

// 根据橱窗ID获取橱窗详情
// burstWindowId 橱窗ID
func (windows *Windows) GetWindowById(burstWindowId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["burstWindowId"] = burstWindowId_
	return APIInterface(windows.config, "eleme.decoration.windows.getWindowById", params)
}

// 获取可见的橱窗信息集合
func (windows *Windows) GetWindows() (interface{}, error) {
	params := make(map[string]interface{})
	return APIInterface(windows.config, "eleme.decoration.windows.getWindows", params)
}

// 创建海报
// poster 海报信息和其关联门店ID和门店商品
func (poster *Poster) CreatePoster(poster_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["poster"] = poster_
	return APIInterface(poster.config, "eleme.decoration.poster.createPoster", params)
}

// 修改海报
// posterId 海报ID
// poster 海报信息和其关联门店ID和门店商品
func (poster *Poster) UpdatePoster(posterId_ int64, poster_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["posterId"] = posterId_
	params["poster"] = poster_
	return APIInterface(poster.config, "eleme.decoration.poster.updatePoster", params)
}

// 作废海报
// poster 作废海报信息
func (poster *Poster) InvalidPoster(poster_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["poster"] = poster_
	return APIInterface(poster.config, "eleme.decoration.poster.invalidPoster", params)
}

// 根据海报ID获取海报详情
// posterId 海报ID
func (poster *Poster) GetPosterDetailById(posterId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["posterId"] = posterId_
	return APIInterface(poster.config, "eleme.decoration.poster.getPosterDetailById", params)
}

// 查询有效的海报信息集合
func (poster *Poster) QueryEffectivePosters() (interface{}, error) {
	params := make(map[string]interface{})
	return APIInterface(poster.config, "eleme.decoration.poster.queryEffectivePosters", params)
}

// 获取历史上传过的海报图片
func (poster *Poster) GetPosterHistoryImage() (interface{}, error) {
	params := make(map[string]interface{})
	return APIInterface(poster.config, "eleme.decoration.poster.getPosterHistoryImage", params)
}

// 保存爆款橱窗
// burstWindow 爆款橱窗信息
func (burstwindow *Burstwindow) SaveBurstWindow(burstWindow_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["burstWindow"] = burstWindow_
	return APIInterface(burstwindow.config, "eleme.decoration.burstWindow.saveBurstWindow", params)
}

// 根据门店ID关闭店铺爆款橱窗
// shopId 门店ID
func (burstwindow *Burstwindow) CloseBurstWindowByShopId(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(burstwindow.config, "eleme.decoration.burstWindow.closeBurstWindowByShopId", params)
}

// 根据店铺ID查询该店铺的爆款橱窗信息
// shopId 店铺ID
func (burstwindow *Burstwindow) GetBurstWindowByShopId(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(burstwindow.config, "eleme.decoration.burstWindow.getBurstWindowByShopId", params)
}

// 根据门店ID集合查询店铺爆款橱窗信息集合
// shopIds 查询条件
func (burstwindow *Burstwindow) QueryBurstWindowList(shopIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopIds"] = shopIds_
	return APIInterface(burstwindow.config, "eleme.decoration.burstWindow.queryBurstWindowList", params)
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

