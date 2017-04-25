package elemeOpenApi

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

type OAuthClient struct {
	config Config
}

func NewAuthClient(conf Config) OAuthClient {
    auth := OAuthClient{}
    auth.SetConfig(conf)
    return auth
}

// 设置配置
func (oauth *OAuthClient) SetConfig(conf Config) {
	oauth.config = conf
}

// 获取授权 URL
func (oauth *OAuthClient) GenerateAuthUrl(state string, scope string, redirectUri string) string {
	data := url.Values{}
	responseType := "code"
	data.Set("state", state)
	data.Set("scope", scope)
	data.Set("redirect_uri", redirectUri)
	data.Set("response_type", responseType)
	data.Set("client_id", oauth.config.strKey)
	host := oauth.config.GetAPIHost()
	baseUrl := getAuthorizationURL(host)
	authUrl := baseUrl + "?" + data.Encode()
	return authUrl
}

// 获取 token 客户端模式
func (oauth *OAuthClient) GetClientToken() Token {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	byteData, _ := json.Marshal(data)
	oauth.config.info(string(byteData))
	host := oauth.config.GetAPIHost()
	tokenURL := getTokenURL(host)
	oauth.config.info(tokenURL)
	auth := oauth.generateAuth()
	req := genRequest(tokenURL, data, auth)
	body := oauth.request(req)
	var token Token
	json.Unmarshal(body, &token)
	return token
}

// 通过 auth code 获取 token
func (oauth *OAuthClient) GetTokenByAuthCode(code string, redirectURL string) Token {
	data := url.Values{}
	data.Set("code", code)
	data.Set("grant_type", "authorization_code")
	data.Set("redirect_uri", redirectURL)
	data.Set("client_id", oauth.config.strKey)
	byteData, _ := json.Marshal(data)
	oauth.config.info(string(byteData))
	host := oauth.config.GetAPIHost()
	tokenURL := getTokenURL(host)
	oauth.config.info(tokenURL)
	auth := oauth.generateAuth()
	req := genRequest(tokenURL, data, auth)
	body := oauth.request(req)
	var token Token
	json.Unmarshal(body, &token)
	return token
}

// 通过刷新 token 获取 token
func (oauth *OAuthClient) GetTokenByRefreshToken(refreshToken string, scope string) Token {
	data := url.Values{}
	data.Set("scope", scope)
	data.Set("grant_type", "refresh_token")
	data.Set("refresh_token", refreshToken)
	byteData, _ := json.Marshal(data)
	oauth.config.info(string(byteData))
	host := oauth.config.GetAPIHost()
	tokenURL := getTokenURL(host)
	oauth.config.info(tokenURL)
	auth := oauth.generateAuth()
	req := genRequest(tokenURL, data, auth)
	body := oauth.request(req)
	var token Token
	json.Unmarshal(body, &token)
	return token
}

// generate the authorization value
func (oauth *OAuthClient) generateAuth() string {
	auth := oauth.config.strKey + ":" + oauth.config.strSecret
	return "Basic " + base64Encoder(auth)
}

func getAuthorizationURL(host string) string {
	urlSuffix := "/authorize"
	return host + urlSuffix
}

func getTokenURL(host string) string {
	urlSuffix := "/token"
	return host + urlSuffix
}

func getAPIURL(host string) string {
	urlSuffix := "/api/v1"
	return host + urlSuffix
}

// encode string to base64 for authrization
func base64Encoder(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// do the request job and return its response
func (oauth *OAuthClient) request(req *http.Request) []byte {
	client := &http.Client{}
	header, _ := json.Marshal(req.Header)
	oauth.config.info(string(header))
	resp, err := client.Do(req)
	if err != nil {
		oauth.config.error(err.Error())
		return nil
	}
	defer resp.Body.Close()
	var reader io.ReadCloser
	reader, err = gzip.NewReader(resp.Body)
	if err != nil {
		oauth.config.error(err.Error())
		return nil
	}
	defer reader.Close()
	buf := new(bytes.Buffer)
	buf.ReadFrom(reader)
	oauth.config.info(buf.String())
	return buf.Bytes()
}

// generate oauth request structure
func genRequest(reqURL string, data url.Values, auth string) *http.Request {
	req, _ := http.NewRequest("POST", reqURL, bytes.NewBufferString(data.Encode()))
	req.Header.Set("authorization", auth)
	req.Header.Set("accept-encoding", "gzip")
	req.Header.Set("user-agent", "eleme-openapi-go-sdk")
	req.Header.Set("content-type", "application/x-www-form-urlencoded; charset=utf-8")
	return req
}
