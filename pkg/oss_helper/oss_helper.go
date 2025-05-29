package oss_helper

import (
	"fmt"
	"io"
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

func ListFiles(prefix string) []string {
	lsRes, err := bucketClient.ListObjectsV2(oss.Prefix(prefix))
	if err != nil {
		fmt.Printf("获取文件列表失败：%s\n", err)
		return nil
	}
	result := []string{}

	for _, v := range lsRes.Objects {
		// val := strings.TrimPrefix(v.Key, prefix+"/")
		// if len(val) > 0 {
		// 	result = append(result, val)
		// }
		result = append(result, v.Key)
	}
	return result
}

func DownloadFile(bucketKey string) []byte {
	body, err := bucketClient.GetObject(bucketKey)
	if err != nil {
		fmt.Printf("下载文件失败：%s\n", err)
		return nil
	}
	defer body.Close()
	data, err := io.ReadAll(body)
	if err != nil {
		fmt.Printf("读取文件内容失败：%s\n", err)
		return nil
	}
	return data
}
