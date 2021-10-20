package oss_helper

import (
	"fmt"
	"lichmaker/girlfriend-gift-1/pkg/config"
	"strings"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

var (
	ossClient       *oss.Client
	bucketClient    *oss.Bucket
	endpoint        string
	accessKeyID     string
	accessKeySecret string
	bucket          string
	baseDir         string
)

func init() {
	var err error
	endpoint = config.Viper.GetString("OSS_ENDPOINT")
	accessKeyID = config.Viper.GetString("OSS_ACCESS_KEY_ID")
	accessKeySecret = config.Viper.GetString("OSS_ACCESS_KEY_SECRET")
	baseDir = config.Viper.GetString("OSS_BASE_DIR")
	baseDir = strings.TrimSuffix(baseDir, "/")
	bucket = config.Viper.GetString("OSS_BUCKET")
	ossClient, err = oss.New(endpoint, accessKeyID, accessKeySecret)
	if err != nil {
		panic(err)
	}
	bucketClient, err = ossClient.Bucket(bucket)
	if err != nil {
		panic(err)
	}
}

func Test() {
	lsRes, _ := ossClient.ListBuckets()

	for _, bucket := range lsRes.Buckets {
		fmt.Println("Buckets:", bucket.Name)
	}

	// bucketClient, err := client.Bucket("zhuzixuan-gift")
}

func Upload(localPath string, ossFileName string) (string, error) {
	bucketKey := fmt.Sprintf("%s/%s", baseDir, ossFileName)
	fmt.Printf("开始上传文件，本地路径 %s , oss路径 %s\n", localPath, bucketKey)
	err := bucketClient.PutObjectFromFile(bucketKey, localPath)
	if err != nil {
		fmt.Printf("上传文件失败：%s\n", err)
		return "", err
	}
	return bucketKey, nil
}

func GetUrl(bucketKey string, ttl int64) string {
	str, _ := bucketClient.SignURL(bucketKey, oss.HTTPGet, ttl)
	return str
}
