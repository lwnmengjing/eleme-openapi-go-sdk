# Golang SDK 接入指南

## 接入指南

  1. Go version == 1.8
  2. 通过 go get 命令安装 SDK
  3. 创建 Config 配置类，填入 key，secret 和 sandbox 参数
  4. 使用 SDK 提供的接口进行开发调试
  5. 上线前调用 Config.SetSandbox 为 false 以及填入正式环境的 key 和 secret
 

## 安装

```go
    go get git.coding.net/napos_openapi/eleme-openapi-go-sdk.git
```

## 基本用法

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

## Token获取
企业应用与个人应用的 token 获取方法略有不同。

实际使用过程中，在 token 获取成功后，该 token 可以使用较长一段时间，需要缓存起来，请勿每次请求都重新获取 token。

### 企业应用


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


### 个人应用

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

### 非授权接口

```go
    import openapi "git.coding.net/napos_openapi/eleme-openapi-go-sdk.git"

    // 新建一个配置实例
    config := openapi.NewConfig(true, app_key, app_secret)

    // 新建一个空 Token
	t := openapi.Token{}
	t.Access_token = ""
	config.SetToken(t)

    // 设置 token
    config.SetToken(token)

    // 创建 openapi client
    eleme := elemeOpenApi.NewAPIClient(config)

    // 完成调用
    res, _ := eleme.Market.SyncMarketMessages(1512057600000, 1516185234000, 0, 10)

    ...

```


## Demo使用方法

该 demo 主要用来演示企业应用的授权流程和展示应用信息

1. 在开发者中心创建企业应用，记下沙箱环境店铺的账号和密码，并在沙箱环境中填入回调地址（该地址需要https）

2. 在 demo 的同一目录创建 config.json 并配置沙箱环境，否则无法运行 demo。

```json
    {"key":"yourkey","secret":"yoursecret","callbackUrl":"yourhost","userId":"","acessToken":"","refreshToken":""}
```

4. 运行 demo。

5. 打开SDK生成的授权URL，使用沙箱店铺的账号和密码进行授权，成功后调转回调接口，输出页面，展示店铺信息

6. 使用沙箱店铺的账号密码在 napos 客户端登陆，会发现刚刚授权的应用已安装，并能够打开应用跳转回调页，展示店铺信息


## Change Log
### [1.32.19]
	Release Date : 2021-01-12
- [Feature] 活动服务增加接口；订单服务增加接口

### [1.32.18]
	Release Date : 2020-11-03
- [Feature] 金融服务增加接口；订单服务更新接口

### [1.32.17]
	Release Date : 2020-04-24
- [Feature] 新增招聘市场服务、企业订单服务；订单服务更新接口
	
### [1.32.16]
	Release Date : 2019-12-27

- [Feature] 商品服务更新、新增接口

### [1.32.15]
	Release Date : 2019-12-27

- [Feature] 金融服务新增接口

### [1.32.14]
	Release Date : 2019-12-23

- [Feature] 店铺服务新增接口

### [1.32.13]
	Release Date : 2019-11-18

- [Feature] 新增商户信用分服务，订单服务更新接口

### [1.32.12]
	Release Date : 2019-9-26

- [Feature] 活动服务下线接口

### [1.32.11]
	Release Date : 2019-9-03

- [Feature] 新增经营体检服务

### [1.32.10]
	Release Date : 2019-7-19

- [Feature] 服务市场服务新增接口

### [1.32.9]
	Release Date : 2019-7-09

- [Feature] 消息服务新增接口

### [1.32.8]
	Release Date : 2019-6-28

- [Feature] 内容服务下线接口

### [1.32.7]
	Release Date : 2019-5-24

- [Feature] 店铺服务新增接口

### [1.32.6]
	Release Date : 2019-5-10

- [Feature] 活动服务下线接口

### [1.32.5]
	Release Date : 2019-4-15

- [Feature] 活动服务新增接口

### [1.32.4]
	Release Date : 2019-4-12

- [Feature] 商户会员卡服务新增接口

### [1.32.3]
	Release Date : 2019-3-5

- [Feature] 商品服务新增接口

### [1.32.2]
	Release Date : 2019-2-22

- [Feature] 餐厅入口流量服务下线接口

### [1.32.1]
	Release Date : 2019-1-3

- [Feature] 商品服务新增更新接口

### [1.32.0]
	Release Date : 2018-12-28

- [Feature] 商品服务新增接口
- [Feature] 服务市场服务新增接口

### [1.31.3]
	Release Date : 2018-12-21

- [Feature] 订单评价服务新增更新接口
- [Feature] 门店装修服务下线接口

### [1.31.2]
	Release Date : 2018-12-7

- [Feature] 订单服务更新接口

### [1.31.1]
	Release Date : 2018-11-30

- [Feature] 店铺服务新增接口
- [Feature] requestId新增时间戳后缀

### [1.31.0]
	Release Date : 2018-11-15

- [Feature] 活动服务新增接口
- [Feature] 店铺服务新增接口
- [Feature] 商户会员卡服务更新接口

### [1.30.5]
	Release Date : 2018-10-17

- [Feature] 活动服务新增接口
- [Feature] 订单评论服务更新接口
- [Feature] 订单服务新增更新接口

### [1.30.4]
	Release Date : 2018-09-29

- [Feature] 新增商户数据服务
- [Feature] 店铺服务更新接口

### [1.29.4]
	Release Date : 2018-09-21

- [Feature] 商户会员卡服务更新接口
- [Feature] 店铺装修服务服务更新接口
- [Feature] 短信服务新增接口

### [1.28.4]
	Release Date : 2018-09-14

- [Feature] 订单服务新增接口

### [1.28.3]
	Release Date : 2018-09-07

- [Feature] 活动服务新增接口
- [Feature] 店铺服务新增接口
- [Feature] 商品服务新增接口

### [1.27.3]
	Release Date : 2018-08-17

- [Feature] 活动服务新增接口
- [Feature] 商户会员卡服务修复bug

### [1.27.2]
	Release Date : 2018-08-03

- [Feature] 商品服务更新接口
- [Feature] 新增商户会员卡服务
- [Feature] 新增CPC竞价服务

### [1.26.2]
	Release Date : 2018-7-27

- [Feature] 订单服务新增接口
- [Feature] 商品服务新增接口
- [Feature] 店铺服务新增接口

### [1.25.2]
	Release Date : 2018-7-13

- [Feature] 订单评论服务新增接口

### [1.25.1]
	Release Date : 2018-7-6

- [Feature] 上传视频接口封装
- [Feature] 商品服务新增接口
- [Feature] 内容服务新增接口

### [1.24.1]
	Release Date : 2018-6-28

- [Feature] 内容服务新增接口
- [Feature] 商品服务新增接口
- [Feature] 订单服务新增接口
- [Feature] 店铺装修服务新增接口
- [Feature] 店铺服务新增接口

### [1.23.1]
	Release Date : 2018-6-22

- [Feature] 更新了店铺服务

### [1.23.0]
	Release Date : 2018-6-08

- [Feature] 更新了活动服务
- [Feature] 更新了订单服务

### [1.22.0]
	Release Date : 2018-5-18

- [Feature] 新增授权码换取OpenId接口
- [Feature] 活动服务新增若干接口
- [Feature] 商品服务新增接口

### [1.21.0]
	Release Date : 2018-3-23

- [Feature] 新增流量服务
- [Feature] 活动服务增加若干红包接口

### [1.20.0]
	Release Date : 2018-2-2

- [Feature] 订单评论服务增加若干接口
- [Feature] 活动服务增加定向送红包接口

### [1.19.0]
	Release Date : 2018-1-17

- [Feature] 增加服务市场服务

### [1.18.0]
	Release Date : 2018-1-5

- [Feature] 商品服务新增根据店铺 Id 查询商品接口

### [1.17.0]
	Release Date : 2017-12-29

- [Feature] 店铺服务新增设置是否支持预定单及预定天数接口

### [1.16.0]
	Release Date : 2017-12-24

- [Feature] 订单服务新增出餐和评价骑手接口
- [Feature] 订单评论服务新增新版回复评论接口

### [1.15.0]
	Release Date : 2017-12-1

- [Feature] 新增众包查询配送费接口

### [1.14.0]
	Release Date : 2017-10-27

- [Feature] 新增了代金券和零元试吃的活动接口

### [1.13.0]
	Release Date : 2017-09-01

- [Feature] 新增了限时抢购活动服务的接口
- [Feature] 新增了订单评论服务的接口

### [1.12.0]
	Release Date : 2017-08-18

- [Feature] 在订单服务中增加了三个关于索赔的接口

### [1.11.0]
	Release Date : 2017-08-03

- [Feature] 在金融服务中增加了五个用来查询账单的特权接口

### [1.10.0]
	Release Date : 2017-07-07

- [Feature] 新增了两个金融服务的特权接口 eleme.finance.queryBalance 和 eleme.finance.queryBalanceLog
- [Feature] 在商品服务中添加了查看活动商品的接口 eleme.product.item.getItemIdsHasActivityByShopId
- [Feature] 在商品服务中 OItem 实体类中 specs 里新增了一个 activityLevel 的属性
- [Feature] 订单服务中 shopId 由原来的 int 变为 int64
- [Feature] 订单服务中 OOrder 实体新增了纳税人识别号  taxpayerId 属性

### [1.9.0]
	Release Date : 2017-05-27

- [Feature] 在订单服务中增加了若干订单操作的轻量接口
- [Feature] 在用户服务中增加了 eleme.user.getPhoneNumber 获取当前授权帐号的手机号的接口
- [Feature] 在店铺服务中增加了 eleme.shop.setOnlineRefund 设置是否支持在线退单

### [1.8.0]
	Release Date : 2017-05-18

- [Feature] 在商品服务中增加了eleme.product.item.getItemByShopIdAndExtendCode 根据商品扩展码获取商品和eleme.product.item.getItemsByShopIdAndBarCode 根据商品条形码获取商品这两个新接口
- [Feature] 在订单服务中增加了 eleme.order.getUnreplyReminders 获取店铺未回复的催单；eleme.order.getUnprocessOrders 查询店铺未处理订单；eleme.order.getCancelOrders 查询店铺未处理的取消单；eleme.order.getRefundOrders 查询店铺未处理的退单；eleme.order.getAllOrders 查询全部订单这五个新接口

### [1.7.0]
	Release Date : 2017-05-12

- [Feature] 在商品服务中增加了 eleme.product.item.batchUpdatePrices 批量修改商品价格的接口
- [Feature] 在订单服务中增加了 eleme.order.cancelDelivery 取消呼叫配送和 eleme.order.callDelivery 呼叫配送这两个接口
- [Feature] 在订单服务中修改了 OOrder 类的定义，增加了一个 List<OActivity> 的属性
- [Feature] 在商品服务中增加了 eleme.product.category.getShopCategoriesWithChildren 查询店铺商品分类，包含二级分类；eleme.product.category.getCategoryWithChildren 查询商品分类详情，包含二级分类；eleme.product.category.createCategoryWithChildren 添加商品分类，支持二级分类；eleme.product.category.updateCategoryWithChildren 更新商品分类，包含二级分类；eleme.product.category.setCategoryPositionsWithChildren 设置二级分类排序这五个接口

### [1.6.0]
	Release Date : 2017-05-8

- [Feature] 在订单服务中新增了 eleme.order.replyReminder eleme.order.getCommodities eleme.order.mgetCommodities eleme.order.getRefundOrder eleme.order.mgetRefundOrders 这五个接口

### [1.5.1]
	Release Date : 2017-05-5

- [Feature] 将签约服务 Package 重命名为 Packs

### [1.5.0]
	Release Date : 2017-04-25

- [Feature] 增加接口查询商品后台分类  eleme.product.category.getBackCategory

### [1.4.0]
	Release Date : 2017-04-21

- [Feature] 增加接口查询店铺当前生效合同类型 eleme.package.getEffectServicePackContract

### [1.3.1]
	Release Date : 2017-04-21

- [Feature] 增加 user-agent eleme-openapi-go-sdk 帮助 debug

### [1.3.0]
	Release Date : 2017-04-14

- [Feature] 增加接口设置送达时间 SetDeliveryTime （只针对特定商户开放）

### [1.2.0]
	Release Date : 2017-4-11

- [Feature] 增加接口批量沽清库存 ClearAndTimingMaxStock （只针对特定商户开放）

### [1.1.0]
	Release Date : 2017-04-7

- [Feature] 增加接口确认订单送达  ReceivedOrder

