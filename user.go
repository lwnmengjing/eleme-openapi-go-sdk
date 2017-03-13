// 商户服务 
package elemeOpenApi

// 获取商户账号信息
func (user *User) GetUser() (interface{}, error) {
	params := make(map[string]interface{})
	return APIInterface(user.config, "eleme.user.getUser", params)
}

