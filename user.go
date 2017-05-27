// 商户服务 
package elemeOpenApi

// 获取商户账号信息
func (user *User) GetUser() (interface{}, error) {
	params := make(map[string]interface{})
	return APIInterface(user.config, "eleme.user.getUser", params)
}

// 获取当前授权帐号的手机号,特权接口仅部分帐号可以调用
func (user *User) GetPhoneNumber() (interface{}, error) {
	params := make(map[string]interface{})
	return APIInterface(user.config, "eleme.user.getPhoneNumber", params)
}

