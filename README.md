# Golang SDK 接入指南

## 接入指南

  1. Go version == 1.8
  2. 通过 go get 命令安装 SDK
  3. 创建 Config 配置类，填入 key，secret 和 sandbox 参数
  4. 使用 SDK 提供的接口进行开发调试
  5. 上线前调用 Config.SetSandbox 为 false 以及填入正式环境的 key 和 secret
 

### 安装

```go
    go get git.coding.net/napos_openapi/eleme-openapi-go-sdk.git
```

### 基本用法

```go
    import openapi "git.coding.net/napos_openapi/eleme-openapi-go-sdk.git"

    // 新建一个配置实例
    config := openapi.NewConfig(false, app_key, app_secret)
    
    // 设置 token
    config.SetToken(token)
    
    // 新建一个 API 实例
    eleme := openapi.NewAPIClient(config)
    
    // 调用服务 API
    data := eleme.Shop.GetShop(shopId)

```

### Token获取
企业应用与个人应用的 token 获取方法略有不同。

实际使用过程中，在 token 获取成功后，该 token 可以使用较长一段时间，需要缓存起来，请勿每次请求都重新获取 token。

#### 企业应用


```go
    import openapi "git.coding.net/napos_openapi/eleme-openapi-go-sdk.git"
    
    // 新建一个配置实例
    config := openapi.NewConfig(false, app_key, app_secret)

    // 新建 oauth 客户端实例
    oauth := openapi.NewAuthClient(config)

    // 获取 token
    token := oauth.GetClientToken()
    
    // 根据 OAuth 2.0 中的对应 state，scope 和 callback_url，获取授权 URL
    authURL := oauth.GenerateAuthUrl(state, scope, redirecrUri)
```

商家打开授权URL，同意授权后，跳转到您的回调页面，并返回code

```go
    ...
    // 通过授权得到的 code，以及正确的 callback_url，获取token
    token := oauth.GetTokenByAuthCode(code, redirectURL)
    ...
```


#### 个人应用

```go
    import openapi "git.coding.net/napos_openapi/eleme-openapi-go-sdk.git"

    // 新建一个配置实例
    config := openapi.NewConfig(false, app_key, app_secret)

    // 新建 oauth 客户端实例
    oauth := openapi.NewAuthClient(config)

    // 获取 token
    token := oauth.GetClientToken()

    // 设置 token
    config.SetToken(token)

    ...

```


### Demo使用方法

该 demo 主要用来演示企业应用的授权流程和展示应用信息

1. 在开发者中心创建企业应用，记下沙箱环境店铺的账号和密码，并在沙箱环境中填入回调地址（该地址需要https）

2. 在 demo 的同一目录创建 config.json 并配置沙箱环境，否则无法运行 demo。

```json
    {"key":"yourkey","secret":"yoursecret","callbackUrl":"yourhost","userId":"","acessToken":"","refreshToken":""}
```

4. 运行 demo。

5. 打开SDK生成的授权URL，使用沙箱店铺的账号和密码进行授权，成功后调转回调接口，输出页面，展示店铺信息

6. 使用沙箱店铺的账号密码在 napos 客户端登陆，会发现刚刚授权的应用已安装，并能够打开应用跳转回调页，展示店铺信息


### Change Log

#### v1.1.0 
- 增加接口确认订单送达  ReceivedOrder

#### v1.2.0 
- 增加接口批量沽清库存 ClearAndTimingMaxStock （只针对特定商户开放）

#### v1.3.0
- 增加接口设置送达时间 SetDeliveryTime （只针对特定商户开放）

#### v1.3.1
- 增加 user-agent eleme-openapi-go-sdk 帮助 debug

#### v1.4.0
- 增加接口查询店铺当前生效合同类型 eleme.package.getEffectServicePackContract

#### v1.5.0
- 增加接口查询商品后台分类  eleme.product.category.getBackCategory

v1.5.1
- 将签约服务 Package 重命名为 Packs

#### v1.6.0
- 在订单服务中新增了 eleme.order.replyReminder eleme.order.getCommodities eleme.order.mgetCommodities eleme.order.getRefundOrder eleme.order.mgetRefundOrders 这五个接口

#### v1.7.0
- 在商品服务中增加了 eleme.product.item.batchUpdatePrices 批量修改商品价格的接口
- 在订单服务中增加了 eleme.order.cancelDelivery 取消呼叫配送和 eleme.order.callDelivery 呼叫配送这两个接口
- 在订单服务中修改了 OOrder 类的定义，增加了一个 List<OActivity> 的属性
- 在商品服务中增加了 eleme.product.category.getShopCategoriesWithChildren 查询店铺商品分类，包含二级分类；eleme.product.category.getCategoryWithChildren 查询商品分类详情，包含二级分类；eleme.product.category.createCategoryWithChildren 添加商品分类，支持二级分类；eleme.product.category.updateCategoryWithChildren 更新商品分类，包含二级分类；eleme.product.category.setCategoryPositionsWithChildren 设置二级分类排序这五个接口

#### v1.8.0
- 在商品服务中增加了eleme.product.item.getItemByShopIdAndExtendCode 根据商品扩展码获取商品和eleme.product.item.getItemsByShopIdAndBarCode 根据商品条形码获取商品这两个新接口
- 在订单服务中增加了 eleme.order.getUnreplyReminders 获取店铺未回复的催单；eleme.order.getUnprocessOrders 查询店铺未处理订单；eleme.order.getCancelOrders 查询店铺未处理的取消单；eleme.order.getRefundOrders 查询店铺未处理的退单；eleme.order.getAllOrders 查询全部订单这五个新接口

