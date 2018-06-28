// 视频服务 
package elemeOpenApi

// 上传视频
// oVideoInfo 视频信息
// shopId 店铺Id
// videoType 视频类型
func (content *Content) UploadVideo(oVideoInfo_ interface{}, shopId_ int64, videoType_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["oVideoInfo"] = oVideoInfo_
	params["shopId"] = shopId_
	params["videoType"] = videoType_
	return APIInterface(content.config, "eleme.content.uploadVideo", params)
}

// 获取efs配置
// videoType 视频类型
func (content *Content) GetEfsConfig(videoType_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["videoType"] = videoType_
	return APIInterface(content.config, "eleme.content.getEfsConfig", params)
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

