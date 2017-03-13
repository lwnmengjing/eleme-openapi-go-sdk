// 消息服务 
package elemeOpenApi

// 获取未到达的推送消息
// appId 应用ID
func (message *Message) GetNonReachedMessages(appId_ int) (interface{}, error) {
	params := make(map[string]interface{})
	params["appId"] = appId_
	return APIInterface(message.config, "eleme.message.getNonReachedMessages", params)
}

// 获取未到达的推送消息实体
// appId 应用ID
func (message *Message) GetNonReachedOMessages(appId_ int) (interface{}, error) {
	params := make(map[string]interface{})
	params["appId"] = appId_
	return APIInterface(message.config, "eleme.message.getNonReachedOMessages", params)
}

