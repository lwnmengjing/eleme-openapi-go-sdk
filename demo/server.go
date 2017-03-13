package main

import (
	"crypto/md5"
	openapi "eleme-openapi-go-sdk"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
)

var conf openapi.Config
var token openapi.Token
var info GetInfoRes
var appConfig Config

var html []byte

type Res struct {
	Message string `json:"message"`
}

type Result struct {
	UserID          int64  `json:"userId"`
	UserName        string `json:"userName"`
	AuthorizedShops []struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
	} `json:"authorizedShops"`
}

type GetInfoRes struct {
	Result struct {
		UserID   string `json:"userId"`
		ShopName string `json:"shopName"`
		OAuthURL string `json:"oAuthUrl"`
	} `json:"result"`
	Error string `json:"error"`
}

type PushMsg struct {
	AppID     int64  `json:"appId"`
	RequestID string `json:"requestId"`
	Type      int64  `json:"type"`
	Message   string `json:"message"`
	ShopID    int64  `json:"shopId"`
	Timestamp int64  `json:"timestamp"`
	UserID    int64  `json:"userId"`
	Signature string `json:"signature"`
}

type Order struct {
	ID          string      `json:"id"`
	OrderID     string      `json:"orderId"`
	Address     string      `json:"address"`
	CreatedAt   string      `json:"createdAt"`
	ActiveAt    string      `json:"activeAt"`
	DeliverFee  float64     `json:"deliverFee"`
	DeliverTime interface{} `json:"deliverTime"`
	Description string      `json:"description"`
	Groups      []struct {
		Name  string `json:"name"`
		Type  string `json:"type"`
		Items []struct {
			ID         int           `json:"id"`
			SkuID      int           `json:"skuId"`
			Name       string        `json:"name"`
			CategoryID int           `json:"categoryId"`
			Price      float64       `json:"price"`
			Quantity   int           `json:"quantity"`
			Total      float64       `json:"total"`
			Additions  []interface{} `json:"additions"`
			NewSpecs   []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"newSpecs"`
			Attributes []struct {
				Name  string `json:"name"`
				Value string `json:"value"`
			} `json:"attributes"`
			ExtendCode string      `json:"extendCode"`
			BarCode    string      `json:"barCode"`
			Weight     interface{} `json:"weight"`
		} `json:"items"`
	} `json:"groups"`
	Invoice                string      `json:"invoice"`
	Book                   bool        `json:"book"`
	OnlinePaid             bool        `json:"onlinePaid"`
	RailwayAddress         interface{} `json:"railwayAddress"`
	PhoneList              []string    `json:"phoneList"`
	ShopID                 int         `json:"shopId"`
	ShopName               string      `json:"shopName"`
	DaySn                  int         `json:"daySn"`
	Status                 string      `json:"status"`
	RefundStatus           string      `json:"refundStatus"`
	UserID                 int         `json:"userId"`
	TotalPrice             float64     `json:"totalPrice"`
	OriginalPrice          float64     `json:"originalPrice"`
	Consignee              string      `json:"consignee"`
	DeliveryGeo            string      `json:"deliveryGeo"`
	DeliveryPoiAddress     string      `json:"deliveryPoiAddress"`
	Invoiced               bool        `json:"invoiced"`
	Income                 float64     `json:"income"`
	ServiceRate            float64     `json:"serviceRate"`
	ServiceFee             float64     `json:"serviceFee"`
	Hongbao                float64     `json:"hongbao"`
	PackageFee             float64     `json:"packageFee"`
	ActivityTotal          float64     `json:"activityTotal"`
	ShopPart               float64     `json:"shopPart"`
	ElemePart              float64     `json:"elemePart"`
	Downgraded             bool        `json:"downgraded"`
	VipDeliveryFeeDiscount float64     `json:"vipDeliveryFeeDiscount"`
	OpenID                 string      `json:"openId"`
}

type Config struct {
	Key          string `json:"key"`
	Secret       string `json:"secret"`
	CallbackURL  string `json:"callbackUrl"`
	UserID       int64  `json:"userId"`
	AcessToken   string `json:"acessToken"`
	RefreshToken string `json:"refreshToken"`
}

func readConfig() bool {
	configText, err := ioutil.ReadFile("./config.json")
	if err == nil {
		json.Unmarshal(configText, &appConfig)
		return true
	}
	return false
}

func writeConfig() {
	configText, _ := json.Marshal(appConfig)
	ioutil.WriteFile("./config.json", configText, 0644)
}

func main() {
	html, _ = ioutil.ReadFile("./index.html")
	bRead := readConfig()
	if !bRead {
		fmt.Println("run this demo after configuring the config.json")
		return
	}
	conf = openapi.NewConfig(true, appConfig.Key, appConfig.Secret)
	var logger openapi.SimpleLogger
	conf.SetLogger(&logger)
	http.HandleFunc("/push", pushHandler)
	http.HandleFunc("/callback", callbackHandler)
	http.HandleFunc("/getInfo", getInfoHandler)
	http.ListenAndServe("0.0.0.0:80", nil)
}

func pushHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in push handler")
	body, _ := ioutil.ReadAll(r.Body)
	var msg PushMsg
	json.Unmarshal(body, &msg)
	signiture := sign(msg)
	if signiture == msg.Signature {
		resp := Res{}
		resp.Message = "OK"
		strResp, _ := json.Marshal(resp)
		fmt.Fprintf(w, string(strResp))
		var order Order
		json.Unmarshal([]byte(msg.Message), &order)
		fmt.Println(order.OrderID)
		eleme := openapi.NewAPIClient(conf)
		eleme.Order.ConfirmOrder(order.OrderID)
		defer r.Body.Close()
	}
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("in callback handler")
	oauth := openapi.NewAuthClient(conf)
	callbackUrl := appConfig.CallbackURL
	r.ParseForm()
	defer r.Body.Close()
	if len(r.Form) == 0 {
		token.Access_token = appConfig.AcessToken
		conf.SetToken(token)
	} else {
		code := r.Form["code"][0]
		token = oauth.GetTokenByAuthCode(code, callbackUrl)
		conf.SetToken(token)
	}
	eleme := openapi.NewAPIClient(conf)
	user, err := eleme.User.GetUser()
	if err != nil {
		info.Result.OAuthURL = oauth.GenerateAuthUrl("ok", "all", callbackUrl)
		fmt.Fprintf(w, string(html))
	} else {
		byteRes, _ := json.Marshal(user)
		fmt.Println(string(byteRes))
		var res Result
		json.Unmarshal(byteRes, &res)
		info.Result.UserID = strconv.FormatInt(res.UserID, 10)
		info.Result.ShopName = res.AuthorizedShops[0].Name
		respHtml := fmt.Sprintf(string(html), info.Result.UserID, info.Result.ShopName)
		appConfig.AcessToken = token.Access_token
		appConfig.RefreshToken = token.Refresh_token
		appConfig.UserID = res.UserID
		writeConfig()
		fmt.Fprintf(w, respHtml)
	}
}

func getInfoHandler(w http.ResponseWriter, r *http.Request) {
	byteResp, _ := json.Marshal(info)
	println(string(byteResp))
	fmt.Fprintf(w, string(byteResp))
}

func combine(left string, right string) string {
	return fmt.Sprintf("%s=%s", left, right)
}

func sign(msg PushMsg) string {
	arrMsg := []string{
		combine("appId", strconv.FormatInt(msg.AppID, 10)),
		combine("requestId", msg.RequestID),
		combine("type", strconv.FormatInt(msg.Type, 10)),
		combine("message", msg.Message),
		combine("shopId", strconv.FormatInt(msg.ShopID, 10)),
		combine("timestamp", strconv.FormatInt(msg.Timestamp, 10)),
		combine("userId", strconv.FormatInt(msg.UserID, 10)),
	}
	sort.Strings(arrMsg)
	result := ""
	for index := range arrMsg {
		result += arrMsg[index]
	}
	signiture := checkSum(result)
	return strings.ToUpper(signiture)
}

func checkSum(str string) string {
	instance := md5.New()
	instance.Write([]byte(str))
	return hex.EncodeToString(instance.Sum(nil))
}
