// 商品服务
package elemeOpenApi

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

// 获取上传图片的url地址(新版)
// hash 图片hash值
func (file *File) GetImageUrl(hash_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["hash"] = hash_
	return APIInterface(file.config, "eleme.file.getImageUrl", params)
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

// 批量上架商品(新版)
// itemIds 商品ID列表
func (item *Item) BatchListItems(itemIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["itemIds"] = itemIds_
	return APIInterface(item.config, "eleme.product.item.batchListItems", params)
}

// 批量下架商品
// specIds 商品及商品规格的列表
func (item *Item) BatchOffShelf(specIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["specIds"] = specIds_
	return APIInterface(item.config, "eleme.product.item.batchOffShelf", params)
}

// 批量下架商品(新版)
// itemIds 商品ID列表
func (item *Item) BatchDelistItems(itemIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["itemIds"] = itemIds_
	return APIInterface(item.config, "eleme.product.item.batchDelistItems", params)
}

// 删除商品
// itemId 商品Id
func (item *Item) RemoveItem(itemId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["itemId"] = itemId_
	return APIInterface(item.config, "eleme.product.item.removeItem", params)
}

// 删除商品(新版)
// itemId 商品Id
func (item *Item) InvalidItem(itemId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["itemId"] = itemId_
	return APIInterface(item.config, "eleme.product.item.invalidItem", params)
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

// 批量更新商品库存(新版)
// stockMap 商品规格ID和库存设值的映射
func (item *Item) BatchUpdateStock(stockMap_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["stockMap"] = stockMap_
	return APIInterface(item.config, "eleme.product.item.batchUpdateStock", params)
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

// 根据商品扩展码获取商品
// shopId 店铺Id
// extendCode 商品扩展码
func (item *Item) GetItemByShopIdAndExtendCode(shopId_ int64, extendCode_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["extendCode"] = extendCode_
	return APIInterface(item.config, "eleme.product.item.getItemByShopIdAndExtendCode", params)
}

// 根据商品条形码获取商品
// shopId 店铺Id
// barCode 商品条形码
func (item *Item) GetItemsByShopIdAndBarCode(shopId_ int64, barCode_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["barCode"] = barCode_
	return APIInterface(item.config, "eleme.product.item.getItemsByShopIdAndBarCode", params)
}

// 批量修改商品价格
// shopId 店铺Id
// specPrices 商品Id及其下SkuId和价格对应Map(限制最多50个)
func (item *Item) BatchUpdatePrices(shopId_ int64, specPrices_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["specPrices"] = specPrices_
	return APIInterface(item.config, "eleme.product.item.batchUpdatePrices", params)
}

// 查询活动商品
// shopId 店铺Id
func (item *Item) GetItemIdsHasActivityByShopId(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(item.config, "eleme.product.item.getItemIdsHasActivityByShopId", params)
}

// 查询店铺活动商品(新版)
// shopId 店铺Id
func (item *Item) GetShopSalesItems(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(item.config, "eleme.product.item.getShopSalesItems", params)
}

// 设置订单餐盒费
// shopId  店铺ID
// status 是否按照订单设置餐盒费
// packingFee 订单餐盒费费用
func (item *Item) SetOrderPackingFee(shopId_ int64, status_ bool, packingFee_ float64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["status"] = status_
	params["packingFee"] = packingFee_
	return APIInterface(item.config, "eleme.product.item.setOrderPackingFee", params)
}

// 分页获取店铺下的商品
// queryPage 分页查询参数
func (item *Item) QueryItemByPage(queryPage_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["queryPage"] = queryPage_
	return APIInterface(item.config, "eleme.product.item.queryItemByPage", params)
}

// 获取原材料树
// shopId 店铺ID
func (item *Item) GetMaterialTree(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(item.config, "eleme.product.item.getMaterialTree", params)
}

// 主料关联配料
// shopId 店铺ID
// mainItemId 主料ID（商品ID）
// ingredientGroup  商品配料分组
func (item *Item) SetIngredient(shopId_ int64, mainItemId_ int64, ingredientGroup_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["mainItemId"] = mainItemId_
	params["ingredientGroup"] = ingredientGroup_
	return APIInterface(item.config, "eleme.product.item.setIngredient", params)
}

// 删除配料
// shopId 店铺ID
// mainItemId 主料ID（商品ID）
func (item *Item) RemoveIngredient(shopId_ int64, mainItemId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["mainItemId"] = mainItemId_
	return APIInterface(item.config, "eleme.product.item.removeIngredient", params)
}

// 针对主菜itemId设置菜品推荐
// shopId 店铺ID
// itemId 商品ID
// relatedItemIds 关联的商品ID
func (item *Item) SetRelatedItemIds(shopId_ int64, itemId_ int64, relatedItemIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["itemId"] = itemId_
	params["relatedItemIds"] = relatedItemIds_
	return APIInterface(item.config, "eleme.product.item.setRelatedItemIds", params)
}

// 对主菜itemId设置是否开启菜品推荐
// shopId 店铺ID
// itemId 商品ID
// display 是否展示
func (item *Item) DisplayRelatedItemIds(shopId_ int64, itemId_ int64, display_ bool) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["itemId"] = itemId_
	params["display"] = display_
	return APIInterface(item.config, "eleme.product.item.displayRelatedItemIds", params)
}

// 针对主菜itemId查询菜品推荐
// shopId 店铺ID
// itemId 商品ID
func (item *Item) GetRelatedItemIds(shopId_ int64, itemId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["itemId"] = itemId_
	return APIInterface(item.config, "eleme.product.item.getRelatedItemIds", params)
}

// 添加多规格商品
// categoryId 商品分类Id
// properties 商品属性
func (item *Item) CreateMultiSpecItem(categoryId_ int64, properties_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["categoryId"] = categoryId_
	params["properties"] = properties_
	return APIInterface(item.config, "eleme.product.item.createMultiSpecItem", params)
}

// 批量添加多规格商品
// categoryId 商品分类Id
// items 商品属性的列表
func (item *Item) BatchCreateMultiSpecItem(categoryId_ int64, items_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["categoryId"] = categoryId_
	params["items"] = items_
	return APIInterface(item.config, "eleme.product.item.batchCreateMultiSpecItem", params)
}

// 更新多规格商品
// itemId 商品Id
// categoryId 商品分类Id
// properties 商品属性
func (item *Item) UpdateMultiSpecItem(itemId_ int64, categoryId_ int64, properties_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["itemId"] = itemId_
	params["categoryId"] = categoryId_
	params["properties"] = properties_
	return APIInterface(item.config, "eleme.product.item.updateMultiSpecItem", params)
}

// 查询店铺商品分类
// shopId 店铺Id
func (category *Category) GetShopCategories(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(category.config, "eleme.product.category.getShopCategories", params)
}

// 查询店铺商品分类，包含二级分类
// shopId 店铺Id
func (category *Category) GetShopCategoriesWithChildren(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(category.config, "eleme.product.category.getShopCategoriesWithChildren", params)
}

// 查询商品分类详情
// categoryId 商品分类Id
func (category *Category) GetCategory(categoryId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["categoryId"] = categoryId_
	return APIInterface(category.config, "eleme.product.category.getCategory", params)
}

// 查询商品分类详情，包含二级分类
// categoryId 商品分类Id
func (category *Category) GetCategoryWithChildren(categoryId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["categoryId"] = categoryId_
	return APIInterface(category.config, "eleme.product.category.getCategoryWithChildren", params)
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

// 添加商品分类，支持二级分类
// shopId 店铺Id
// name 商品分类名称，长度需在50字以内
// parentId 父分类ID，如果没有可以填0
// description 商品分类描述，长度需在50字以内
func (category *Category) CreateCategoryWithChildren(shopId_ int64, name_ string, parentId_ int64, description_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["name"] = name_
	params["parentId"] = parentId_
	params["description"] = description_
	return APIInterface(category.config, "eleme.product.category.createCategoryWithChildren", params)
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

// 更新商品分类，包含二级分类
// categoryId 商品分类Id
// name 商品分类名称，长度需在50字以内
// parentId 父分类ID，如果没有可以填0
// description 商品分类描述，长度需在50字以内
func (category *Category) UpdateCategoryWithChildren(categoryId_ int64, name_ string, parentId_ int64, description_ string) (interface{}, error) {
	params := make(map[string]interface{})
	params["categoryId"] = categoryId_
	params["name"] = name_
	params["parentId"] = parentId_
	params["description"] = description_
	return APIInterface(category.config, "eleme.product.category.updateCategoryWithChildren", params)
}

// 删除商品分类
// categoryId 商品分类Id
func (category *Category) RemoveCategory(categoryId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["categoryId"] = categoryId_
	return APIInterface(category.config, "eleme.product.category.removeCategory", params)
}

// 删除商品分类(新版)
// categoryId 商品分类Id
func (category *Category) InvalidCategory(categoryId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["categoryId"] = categoryId_
	return APIInterface(category.config, "eleme.product.category.invalidCategory", params)
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

// 设置分类排序(新版)
// shopId 饿了么店铺Id
// categoryIds 需要排序的全部一级分类Id
func (category *Category) SetCategorySequence(shopId_ int64, categoryIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["categoryIds"] = categoryIds_
	return APIInterface(category.config, "eleme.product.category.setCategorySequence", params)
}

// 设置二级分类排序
// shopId 饿了么店铺Id
// categoryWithChildrenIds 需要排序的父分类Id，及其下属的二级分类ID
func (category *Category) SetCategoryPositionsWithChildren(shopId_ int64, categoryWithChildrenIds_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["categoryWithChildrenIds"] = categoryWithChildrenIds_
	return APIInterface(category.config, "eleme.product.category.setCategoryPositionsWithChildren", params)
}

// 查询商品后台类目
// shopId 店铺Id
func (category *Category) GetBackCategory(shopId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	return APIInterface(category.config, "eleme.product.category.getBackCategory", params)
}

// 设置分类类型
// shopId 店铺Id
// categoryId 商品分类Id
// categoryType 分类类型
func (category *Category) SetCategoryType(shopId_ int64, categoryId_ int64, categoryType_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["categoryId"] = categoryId_
	params["categoryType"] = categoryType_
	return APIInterface(category.config, "eleme.product.category.setCategoryType", params)
}

// 设置分组分时段置顶
// shopId 店铺Id
// categoryId 商品分类Id
// dayPartingStick 置顶时间设置
func (category *Category) SetDayPartingStickTime(shopId_ int64, categoryId_ int64, dayPartingStick_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["categoryId"] = categoryId_
	params["dayPartingStick"] = dayPartingStick_
	return APIInterface(category.config, "eleme.product.category.setDayPartingStickTime", params)
}

// 删除分组的分时置顶功能
// shopId 店铺Id
// categoryId 商品分类Id
func (category *Category) RemoveDayPartingStickTime(shopId_ int64, categoryId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["shopId"] = shopId_
	params["categoryId"] = categoryId_
	return APIInterface(category.config, "eleme.product.category.removeDayPartingStickTime", params)
}

// 添加套餐
// categoryId 分类Id
// oPackage 套餐属性
func (package_ *Package) CreatePackage(categoryId_ int64, oPackage_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["categoryId"] = categoryId_
	params["oPackage"] = oPackage_
	return APIInterface(package_.config, "eleme.product.package.createPackage", params)
}

// 修改套餐基本信息
// itemId 新套餐id即OItem中的商品Id
// categoryId 分类Id即OCategory中的分类Id
// update 套餐基本信息
func (package_ *Package) UpdatePackageContent(itemId_ int64, categoryId_ int64, update_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["itemId"] = itemId_
	params["categoryId"] = categoryId_
	params["update"] = update_
	return APIInterface(package_.config, "eleme.product.package.updatePackageContent", params)
}

// 修改套餐和主料的关联关系
// itemId 新套餐id即OItem中的商品Id
// packages 套餐关系
func (package_ *Package) UpdatePackageRelation(itemId_ int64, packages_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["itemId"] = itemId_
	params["packages"] = packages_
	return APIInterface(package_.config, "eleme.product.package.updatePackageRelation", params)
}

// 删除套餐
// itemId 套餐Id
func (package_ *Package) RemovePackage(itemId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["itemId"] = itemId_
	return APIInterface(package_.config, "eleme.product.package.removePackage", params)
}
