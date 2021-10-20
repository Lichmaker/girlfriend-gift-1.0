package uploader

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"lichmaker/girlfriend-gift-1/app/models/pool"
	"lichmaker/girlfriend-gift-1/pkg/oss_helper"
	"os"
	"path/filepath"
)

func Do(pathData map[string]string) error {
	md5Map := make(map[string]string)
	var md5Array []string
	for path := range pathData {
		md5 := GetFileMd5(path)
		md5Map[md5] = path
		md5Array = append(md5Array, md5)
	}
	notIn := pool.QueryMd5NotIn(md5Array)
	for _, _md5 := range notIn {
		// 上传到 oss
		_, fileName := filepath.Split(md5Map[_md5])
		ossPath, err := oss_helper.Upload(md5Map[_md5], fileName)
		if err == nil {
			pool.Create(_md5, md5Map[_md5], ossPath, pathData[md5Map[_md5]])
		}
		fmt.Println(oss_helper.GetUrl(ossPath, int64(100)))
	}
	return nil
}

func GetFileMd5(path string) string {
	file, err := os.Open(path)
	if err != nil {
		fmt.Printf("打开文件失败,file=%s, err=%s", path, err)
		return ""
	}
	defer file.Close()
	md5h := md5.New()
	io.Copy(md5h, file)
	return hex.EncodeToString(md5h.Sum(nil))
}
