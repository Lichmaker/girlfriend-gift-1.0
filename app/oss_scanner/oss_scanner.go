package oss_scanner

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"lichmaker/girlfriend-gift-1/app/models/pool"
	"lichmaker/girlfriend-gift-1/pkg/oss_helper"
	"regexp"
	"strings"
	"time"
)

func Run(prefix string) {
	// 扫描oss中目标目录中的文件
	keysArr := oss_helper.ListFiles(prefix)

	// 遍历文件
	counter := 0
	for _, keyItem := range keysArr {
		keyFileName := strings.TrimPrefix(keyItem, prefix+"/")
		if len(keyFileName) == 0 {
			continue
		}

		dateStr, ok := parseFileName(keyFileName)
		if !ok {
			continue
		}

		fileData := oss_helper.DownloadFile(keyItem)
		if len(fileData) == 0 {
			continue
		}
		md5 := md5.Sum(fileData)
		md5Str := hex.EncodeToString(md5[:])
		exists := pool.CheckMd5Exists(md5Str)
		if exists {
			continue
		}
		pool.Create(md5Str, keyItem, keyItem, dateStr)
		counter++
	}

	fmt.Println("从OSS中记录到新图片：", counter)
}

// 解析文件名是否符合规范，需要是包含日期的.jpg文件或者.jpeg
// 日期格式规定为 2006-01-02
func parseFileName(fileName string) (string, bool) {
	// 检查文件后缀名
	if !strings.HasSuffix(fileName, ".jpg") && !strings.HasSuffix(fileName, ".jpeg") {
		return "", false
	}

	// 使用正则表达式匹配日期格式 2006-01-02
	pattern := `\d{4}-\d{2}-\d{2}`
	reg := regexp.MustCompile(pattern)
	dateStr := reg.FindString(fileName)

	if dateStr == "" {
		return "", false
	}

	// 验证日期格式是否正确
	_, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		return "", false
	}
	return dateStr, true
}
