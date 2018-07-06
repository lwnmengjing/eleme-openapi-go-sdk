package elemeOpenApi

type EfsConfig struct {
	efsAddress  string
	bucketName  string
	credentials Credentials
}

type Credentials struct {
	AccessKeyId     string
	SecretAccessKey string
	SessionToken    string
	Expiration      float64
}

func (efsConf *EfsConfig) SetEfsAddress(str string) {
	efsConf.efsAddress = str
}

func (efsConf *EfsConfig) SetBucketName(str string) {
	efsConf.bucketName = str
}

func (efsConf *EfsConfig) SetCredentials(credentials Credentials) {
	efsConf.credentials = credentials
}

func (efsConf *EfsConfig) SetAccessKeyId(str string) {
	efsConf.credentials.AccessKeyId = str
}

func (efsConf *EfsConfig) SetSecretAccessKey(str string) {
	efsConf.credentials.SecretAccessKey = str
}

func (efsConf *EfsConfig) SetSessionToken(str string) {
	efsConf.credentials.SessionToken = str
}

func (efsConf *EfsConfig) SetExpiration(str float64) {
	efsConf.credentials.Expiration = str
}
