// 短信服务 
package elemeOpenApi

// 发送短信
// request 短信发送请求
func (sms *Sms) SendMessage(request_ interface{}) (interface{}, error) {
	params := make(map[string]interface{})
	params["request"] = request_
	return APIInterface(sms.config, "eleme.sms.sendMessage", params)
}

