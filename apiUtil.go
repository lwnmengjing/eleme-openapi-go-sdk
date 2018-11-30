package elemeOpenApi

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
	"strings"
	"time"
)

type ApiBody struct {
	Token     string                 `json:"token"`
	Nop       string                 `json:"nop"`
	Metas     map[string]interface{} `json:"metas"`
	Params    map[string]interface{} `json:"params"`
	Action    string                 `json:"action"`
	ID        string                 `json:"id"`
	Signature string                 `json:"signature"`
}

type apiResponse struct {
	ID     string
	Result interface{}
	Error  interface{}
}

type APIError struct {
	Code    string
	Message string
}

func (e APIError) Error() string {
	return e.Code + " " + e.Message
}

// api 通用接口
func APIInterface(config *Config, strAction string, params map[string]interface{}) (interface{}, error) {
	body := geneAPIRequestBody(config, strAction, params)
	req := geneAPIRequest(config, body)
	// could every class share a same http client?
	client := &http.Client{}
	response, _ := client.Do(req)
	defer response.Body.Close()
	responseBody, _ := ioutil.ReadAll(response.Body)
	config.info(string(responseBody))
	var apiResp apiResponse
	json.Unmarshal([]byte(responseBody), &apiResp)
	if apiResp.Error != nil {
		tmp, _ := json.Marshal(apiResp.Error)
		config.error(string(tmp))
		var apiErr APIError
		json.Unmarshal(tmp, &apiErr)
		return nil, apiErr
	}
	return apiResp.Result, nil

}

// 给 body 签名
func (body *ApiBody) Sign(secret string) {
	strConcatArr := concatArr(body.Metas, body.Params)
	strToSign := body.Action + body.Token + strConcatArr + secret
	strSign := checkSum(strToSign)
	body.Signature = strings.ToUpper(strSign)
}

func geneAPIRequest(config *Config, body ApiBody) *http.Request {
	strHost := config.GetAPIHost()
	strReqURL := getAPIURL(strHost)
	strMethod := "POST"
	jsonBody, _ := json.Marshal(body)
	req, _ := http.NewRequest(strMethod, strReqURL, strings.NewReader(string(jsonBody)))
	req.Header.Set("content-type", "application/json")
	req.Header.Set("user-agent", "eleme-openapi-go-sdk")
	header, _ := json.Marshal(req.Header)
	config.info(strReqURL)
	config.info(string(jsonBody))
	config.info(string(header))
	return req
}

func geneAPIRequestBody(config *Config, strAction string, params map[string]interface{}) ApiBody {
	NOP := "1.0.0"
	body := ApiBody{}
	body.Token = config.token.Access_token
	body.Nop = NOP
	body.Action = strAction
	body.ID = pseudoUUID4()
	metas := make(map[string]interface{})
	metas["app_key"] = config.strKey
	metas["timestamp"] = getTimestamp()
	body.Metas = metas
	body.Params = params
	body.Sign(config.strSecret)
	return body
}

func concatArr(metas map[string]interface{}, params map[string]interface{}) string {
	var arr []string
	arrMetas := mapToArray(metas)
	for indexMetas := range arrMetas {
		arr = append(arr, arrMetas[indexMetas])
	}
	arrParams := mapToArray(params)
	for indexParams := range arrParams {
		arr = append(arr, arrParams[indexParams])
	}
	sort.Strings(arr)
	result := ""
	for index := range arr {
		result += arr[index]
	}
	return result
}

func mapToArray(dict map[string]interface{}) []string {
	var arr []string
	var _type string
	var _value string
	var item string
	funcArr := formatFuncArr()
	for key := range dict {
		_type = getType(dict[key])
		_value = funcArr[_type](dict[key])
		item = key + "=" + _value
		arr = append(arr, item)
	}
	return arr
}

func formatFuncArr() map[string]func(interface{}) string {
	funcArr := make(map[string]func(interface{}) string)
	funcArr["int"] = intProcessor
	funcArr["int64"] = int64Processor
	funcArr["float64"] = float64Processor
	funcArr["bool"] = boolProcessor
	funcArr["string"] = stringProcessor
	funcArr["unknown"] = unknownProcess
	return funcArr
}

func getType(i interface{}) string {
	switch i.(type) {
	case string:
		return "string"
	case bool:
		return "bool"
	case int:
		return "int"
	case int64:
		return "int64"
	case float64:
		return "float64"
	default:
		return "unknown"
	}
}

func checkSum(str string) string {
	instance := md5.New()
	instance.Write([]byte(str))
	return hex.EncodeToString(instance.Sum(nil))
}

func pseudoUUID4() string {
	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
	}
	b[6] = (b[6] | 0x40) & 0x4F
	b[8] = (b[8] | 0x80) & 0xBF
	uuid := fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
	return strings.ToUpper(strings.Replace(uuid, "-", "", -1)) + "|" + strconv.FormatInt(getMilliTimestamp(), 10)
}

func getMilliTimestamp() int64 {
	timestamp := time.Now().UnixNano()
	return timestamp / int64(time.Millisecond)
}

func getTimestamp() int64 {
	timestamp := time.Now().Unix()
	return timestamp
}

func stringProcessor(i interface{}) string {
	return quote(i.(string))
}

func intProcessor(i interface{}) string {
	number := i.(int)
	return strconv.Itoa(number)
}

func int64Processor(i interface{}) string {
	number := i.(int64)
	return strconv.FormatInt(number, 10)
}

func float64Processor(i interface{}) string {
	number := i.(float64)
	return strconv.FormatFloat(number, 'f', -1, 64)
}

func boolProcessor(i interface{}) string {
	b := i.(bool)
	if b {
		return "true"
	}
	return "false"
}

func unknownProcess(i interface{}) string {
	unknown, _ := json.Marshal(i)
	return string(unknown)
}

func quote(str string) string {
	return "\"" + str + "\""
}
