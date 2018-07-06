package elemeOpenApi

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const fileMaxSize = 200 * 1024 * 1024

var fileExtList = []string{"mp4", "mov"}

type VideoClient struct {
	config Config
}

func NewVideoClient(conf Config) VideoClient {
	videoClient := VideoClient{}
	videoClient.SetConfig(conf)
	return videoClient
}

// 设置配置
func (videoClient *VideoClient) SetConfig(conf Config) {
	videoClient.config = conf
}

func (videoClient *VideoClient) UploadVideoClient(filePath string, title string, desc string, videoType string, shopId int64) (float64, error) {

	fileInfo, err1 := os.Stat(filePath)

	if err1 != nil {
		panic(err1)
	}

	if fileInfo.Size() > fileMaxSize {
		panic("视频大小不能超过200M")
	}

	fileExt := filepath.Ext(fileInfo.Name())[1:]

	if Index(fileExtList, strings.ToLower(fileExt)) == -1 {
		panic("只支持mp4和mov格式的视频")
	}

	eleme := NewAPIClient(videoClient.config)
	efsConfigGet, err2 := eleme.Content.GetEfsConfig(videoType)

	if err2 != nil {
		panic(err1)
	}
	efsConfigRes := efsConfigGet.(map[string]interface{})

	efsConfigCred := efsConfigRes["credentials"].(map[string]interface{})

	var efsConfig EfsConfig
	efsConfig.SetEfsAddress(efsConfigRes["efsAddress"].(string))
	efsConfig.SetBucketName(efsConfigRes["bucketName"].(string))
	efsConfig.SetAccessKeyId(efsConfigCred["accessKeyId"].(string))
	efsConfig.SetSecretAccessKey(efsConfigCred["secretAccessKey"].(string))
	efsConfig.SetSessionToken(efsConfigCred["sessionToken"].(string))
	efsConfig.SetExpiration(efsConfigCred["expiration"].(float64))

	videoFile, err3 := os.Open(filePath)

	if err3 != nil {
		panic(err3)
	}
	efsClient := NewEfsClient(efsConfig)

	t := time.Now()
	key := strconv.FormatInt(shopId, 10) + "_" + strconv.FormatInt(t.Unix(), 10)

	versionId, err4 := efsClient.putObject(efsConfig.bucketName, key, videoFile)
	if err4 != nil {
		panic(err4)
	}

	var videoInfo = map[string]string{
		"videoKey":    key,
		"versionId":   versionId,
		"sizeInByte":  strconv.FormatInt(fileInfo.Size(), 10),
		"title":       title,
		"description": desc}

	videoId, err5 := eleme.Content.UploadVideo(videoInfo, shopId, videoType)
	return videoId.(float64), err5
}

func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}
