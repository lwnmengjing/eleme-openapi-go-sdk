// 视频服务 
package elemeOpenApi

// 建立视频与相对应的业务的关联关系
// videoId 视频Id
// bizId 业务Id(如业务类型为GOOD，业务Id为商品Id)
// bindBizType 业务类型
func (content *Content) SetVideoBindRelation(videoId_ int64, bizId_ int64, bindBizType_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["videoId"] = videoId_
	params["bizId"] = bizId_
	params["bindBizType"] = bindBizType_
	return APIInterface(content.config, "eleme.content.setVideoBindRelation", params)
}

// 取消视频与对应业务的关联关系
// videoId 视频Id
// bizId 业务Id(如业务类型为GOOD，业务Id为商品Id)
// bindBizType 业务类型
func (content *Content) UnsetVideoBindRelation(videoId_ int64, bizId_ int64, bindBizType_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["videoId"] = videoId_
	params["bizId"] = bizId_
	params["bindBizType"] = bindBizType_
	return APIInterface(content.config, "eleme.content.unsetVideoBindRelation", params)
}

// 通过视频id查询视频信息
// videoId 视频Id
func (content *Content) GetVideoInfo(videoId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["videoId"] = videoId_
	return APIInterface(content.config, "eleme.content.getVideoInfo", params)
}

// 通过视频id获取所有相关联的业务关系
// videoId 视频Id
func (content *Content) GetVideoBindInfo(videoId_ int64) (interface{}, error) {
	params := make(map[string]interface{})
	params["videoId"] = videoId_
	return APIInterface(content.config, "eleme.content.getVideoBindInfo", params)
}

