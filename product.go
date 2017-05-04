// 商品服务 
package elemeOpenApi

// 查询店铺商品分类
// shopId 店铺Id
func (category *Category) GetShopCategories(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(category.config, "eleme.product.category.getShopCategories", params)
}

// 查询商品分类详情
// categoryId 商品分类Id
func (category *Category) GetCategory(categoryId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["categoryId"] = categoryId_
	return APIInterface(category.config, "eleme.product.category.getCategory", params)
}

// 添加商品分类
// shopId 店铺Id
// name 商品分类名称，长度需在50字以内
// description 商品分类描述，长度需在50字以内
func (category *Category) CreateCategory(shopId_ int64, name_ string, description_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["name"] = name_
	params["description"] = description_
	return APIInterface(category.config, "eleme.product.category.createCategory", params)
}

// 更新商品分类
// categoryId 商品分类Id
// name 商品分类名称，长度需在50字以内
// description 商品分类描述，长度需在50字以内
func (category *Category) UpdateCategory(categoryId_ int64, name_ string, description_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["categoryId"] = categoryId_
	params["name"] = name_
	params["description"] = description_
	return APIInterface(category.config, "eleme.product.category.updateCategory", params)
}

// 删除商品分类
// categoryId 商品分类Id
func (category *Category) RemoveCategory(categoryId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["categoryId"] = categoryId_
	return APIInterface(category.config, "eleme.product.category.removeCategory", params)
}

// 设置分类排序
// shopId 饿了么店铺Id
// categoryIds 需要排序的分类Id
func (category *Category) SetCategoryPositions(shopId_ int64, categoryIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["categoryIds"] = categoryIds_
	return APIInterface(category.config, "eleme.product.category.setCategoryPositions", params)
}

// 查询商品后台分类
// shopId 店铺Id
func (category *Category) GetBackCategory(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(category.config, "eleme.product.category.getBackCategory", params)
}

// 上传图片，返回图片的hash值
// image 文件内容base64编码值
func (file *File) UploadImage(image_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["image"] = image_
	return APIInterface(file.config, "eleme.file.uploadImage", params)
}

// 通过远程URL上传图片，返回图片的hash值
// url 远程Url地址
func (file *File) UploadImageWithRemoteUrl(url_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["url"] = url_
	return APIInterface(file.config, "eleme.file.uploadImageWithRemoteUrl", params)
}

// 获取上传文件的访问URL，返回文件的Url地址
// hash 图片hash值
func (file *File) GetUploadedUrl(hash_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["hash"] = hash_
	return APIInterface(file.config, "eleme.file.getUploadedUrl", params)
}

// 获取一个分类下的所有商品
// categoryId 商品分类Id
func (item *Item) GetItemsByCategoryId(categoryId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["categoryId"] = categoryId_
	return APIInterface(item.config, "eleme.product.item.getItemsByCategoryId", params)
}

// 查询商品详情
// itemId 商品Id
func (item *Item) GetItem(itemId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["itemId"] = itemId_
	return APIInterface(item.config, "eleme.product.item.getItem", params)
}

// 批量查询商品详情
// itemIds 商品Id的列表
func (item *Item) BatchGetItems(itemIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["itemIds"] = itemIds_
	return APIInterface(item.config, "eleme.product.item.batchGetItems", params)
}

// 添加商品
// categoryId 商品分类Id
// properties 商品属性
func (item *Item) CreateItem(categoryId_ int64, properties_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["categoryId"] = categoryId_
	params["properties"] = properties_
	return APIInterface(item.config, "eleme.product.item.createItem", params)
}

// 批量添加商品
// categoryId 商品分类Id
// items 商品属性的列表
func (item *Item) BatchCreateItems(categoryId_ int64, items_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["categoryId"] = categoryId_
	params["items"] = items_
	return APIInterface(item.config, "eleme.product.item.batchCreateItems", params)
}

// 更新商品
// itemId 商品Id
// categoryId 商品分类Id
// properties 商品属性
func (item *Item) UpdateItem(itemId_ int64, categoryId_ int64, properties_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["itemId"] = itemId_
	params["categoryId"] = categoryId_
	params["properties"] = properties_
	return APIInterface(item.config, "eleme.product.item.updateItem", params)
}

// 批量置满库存
// specIds 商品及商品规格的列表
func (item *Item) BatchFillStock(specIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["specIds"] = specIds_
	return APIInterface(item.config, "eleme.product.item.batchFillStock", params)
}

// 批量沽清库存
// specIds 商品及商品规格的列表
func (item *Item) BatchClearStock(specIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["specIds"] = specIds_
	return APIInterface(item.config, "eleme.product.item.batchClearStock", params)
}

// 批量上架商品
// specIds 商品及商品规格的列表
func (item *Item) BatchOnShelf(specIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["specIds"] = specIds_
	return APIInterface(item.config, "eleme.product.item.batchOnShelf", params)
}

// 批量下架商品
// specIds 商品及商品规格的列表
func (item *Item) BatchOffShelf(specIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["specIds"] = specIds_
	return APIInterface(item.config, "eleme.product.item.batchOffShelf", params)
}

// 删除商品
// itemId 商品Id
func (item *Item) RemoveItem(itemId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["itemId"] = itemId_
	return APIInterface(item.config, "eleme.product.item.removeItem", params)
}

// 批量删除商品
// itemIds 商品Id的列表
func (item *Item) BatchRemoveItems(itemIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["itemIds"] = itemIds_
	return APIInterface(item.config, "eleme.product.item.batchRemoveItems", params)
}

// 批量更新商品库存
// specStocks 商品以及规格库存列表
func (item *Item) BatchUpdateSpecStocks(specStocks_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["specStocks"] = specStocks_
	return APIInterface(item.config, "eleme.product.item.batchUpdateSpecStocks", params)
}

// 设置商品排序
// categoryId 商品分类Id
// itemIds 商品Id列表
func (item *Item) SetItemPositions(categoryId_ int64, itemIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["categoryId"] = categoryId_
	params["itemIds"] = itemIds_
	return APIInterface(item.config, "eleme.product.item.setItemPositions", params)
}

// 批量沽清库存并在次日2:00开始置满
// clearStocks 店铺Id及商品Id的列表
func (item *Item) ClearAndTimingMaxStock(clearStocks_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["clearStocks"] = clearStocks_
	return APIInterface(item.config, "eleme.product.item.clearAndTimingMaxStock", params)
}

