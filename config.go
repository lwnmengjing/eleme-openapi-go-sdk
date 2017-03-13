package elemeOpenApi

type Config struct {
	bSandbox  bool
	bSetURL   bool
	strKey    string
	strSecret string
	strURL    string
	logger    ElemeSdkLogger
	bLogger   bool
	token     Token
}

type ElemeSdkLogger interface {
	Info(message string)
	Error(message string)
}

type Token struct {
	Access_token  string
	Token_type    string
	Expires_in    int
	Refresh_token string
}

// 构造一个新的 openapi 配置
func NewConfig(bSandbox bool, key string, secret string) Config {
	conf := Config{}
	conf.SetSandbox(bSandbox)
	conf.SetKey(key)
	conf.SetSecret(secret)
	conf.bSetURL = false
	return conf
}

// 设置是否为沙箱环境
func (conf *Config) SetSandbox(b bool) {
	conf.bSandbox = b
}

// 设置 openapi key
func (conf *Config) SetKey(key string) {
	conf.strKey = key
}

// 设置 openapi secret
func (conf *Config) SetSecret(secret string) {
	conf.strSecret = secret
}

// 设置 openapi URL
func (conf *Config) SetURL(url string) {
	conf.strURL = url
	conf.bSetURL = true
}

// 设置 token
func (conf *Config) SetToken(token Token) {
	conf.token = token
}

// 获取 openapi host
func (conf *Config) GetAPIHost() string {
	if conf.bSetURL {
		return conf.strURL
	}
	if conf.bSandbox {
		return sandboxURL()
	}
	return productURL()
}

// 设置日志接口
func (conf *Config) SetLogger(logger ElemeSdkLogger) {
	conf.logger = logger
	conf.bLogger = true
}

func (conf *Config) diagnosis() {
	if conf.logger == nil {
		conf.bLogger = false
	} else {
		conf.bLogger = true
	}
}

func (conf *Config) error(err string) {
	conf.diagnosis()
	if conf.bLogger {
		conf.logger.Error(err)
	}
}

func (conf *Config) info(info string) {
	conf.diagnosis()
	if conf.bLogger {
		conf.logger.Info(info)
	}
}

func productURL() string {
	return "https://open-api.shop.ele.me"
}

func sandboxURL() string {
	return "https://open-api-sandbox.shop.ele.me"
}
