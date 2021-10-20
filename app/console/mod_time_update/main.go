package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

func main() {
	// 获取命令行参数
	args := os.Args
	if len(args) <= 1 {
		panic("无法获取参数")
	}
	p := args[1]
	if len(p) == 0 {
		panic("无法获取参数")
	}
	fileinfo, err := os.Stat(p)
	if err != nil {
		panic("传入路径错误")
	}
	if !fileinfo.IsDir() {
		panic("传入路径不是一个folder")
	}
	loc,_ := time.LoadLocation("PRC")

	_ = filepath.Walk(p, func(p2 string, info os.FileInfo, err1 error) error {
		if info.IsDir() {
			return nil
		}
		dateString := strings.TrimSuffix(info.Name(),path.Ext(info.Name()))
		dateParse, err2 := time.ParseInLocation("2006-01-02", dateString, loc)
		if err2 != nil || dateParse.Unix() <= 0 {
			fmt.Println("faild "+p2+" 无法解析日期，忽略")
			return nil
		}
		err3 := os.Chtimes(p2, dateParse, dateParse)
		if err3 != nil {
			fmt.Printf("%s修改失败：%v", p2, err3)
			return nil
		}
		fmt.Println("success "+p2+" 修改成功")
		return nil
	})

}
