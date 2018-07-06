package elemeOpenApi

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type EfsClient struct {
	efsConfig EfsConfig
}

var s3Session *session.Session
var s3Client *s3.S3

func NewEfsClient(efsConf EfsConfig) EfsClient {
	efsClient := EfsClient{}
	efsClient.SetEfsConfig(efsConf)
	efsClient.init()
	return efsClient
}

// 设置配置
func (efsClient *EfsClient) SetEfsConfig(efsConf EfsConfig) {
	efsClient.efsConfig = efsConf
}

func (efsClient *EfsClient) init() {
	accessID := efsClient.efsConfig.credentials.AccessKeyId
	secretKey := efsClient.efsConfig.credentials.SecretAccessKey
	token := efsClient.efsConfig.credentials.SessionToken
	efsAddress := efsClient.efsConfig.efsAddress

	s3Config := &aws.Config{
		Credentials:      credentials.NewStaticCredentials(accessID, secretKey, token),
		Endpoint:         aws.String(efsAddress), //efs 服务地址
		Region:           aws.String("eleme"),    //使用默认值即可
		DisableSSL:       aws.Bool(true),         // 是否使用https访问
		S3ForcePathStyle: aws.Bool(true),         //是否强制请求使用路径模式寻址， ie:, 路径模式'http://efs.ele.me/BUCKET/KEY'. 默认情况下， s3 client使用bucket作为子域名方式寻址. ie, 'http://BUCKET.efs.ele.me/KEY'
	}
	// s3Session可用于并发的创建用于发送具体请求的client端
	var err error
	s3Session, err = session.NewSession(s3Config)
	if err != nil {
		fmt.Println("create session error: %v", err)
		return
	}
	// s3Client 用于发送具体请求
	s3Client = s3.New(s3Session)
}

func (efsClient *EfsClient) putObject(bucket, key string, video *os.File) (string, error) {

	putOutput, err := s3Client.PutObject(&s3.PutObjectInput{
		Body:   video,
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	return *putOutput.VersionId, err
}
