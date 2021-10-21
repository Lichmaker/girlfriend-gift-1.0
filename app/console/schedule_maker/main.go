package main

import (
	"fmt"
	"lichmaker/girlfriend-gift-1/app/models/schedule"
	"lichmaker/girlfriend-gift-1/bootstrap"
	"lichmaker/girlfriend-gift-1/config"
	"lichmaker/girlfriend-gift-1/pkg/model"
	"time"
)

type result struct {
	Md5      string
	Oss_path string
	Mod_date string
}

func main() {
	config.Initialize()
	bootstrap.SetupDB()

	// 每天执行一次，每次执行都更新后天的数据
	timestamp := time.Now().Unix() + 86400*2
	dateString := time.Unix(timestamp, 0).Format("2006-01-02")

	// 最多跑10次，找没重复使用的
	var md5 string
	var result result
	for i := 0; i < 10; i++ {
		sql := "SELECT md5, oss_path, mod_date FROM pool ORDER BY rand() LIMIT 1"
		model.DB.Raw(sql).Scan(&result)
		
		md5 = result.Md5
		if schedule.Exists(md5) {
			continue
		} else {
			break
		}
	}

	var model schedule.Schedule
	model.Md5 = result.Md5
	model.Date = dateString
	model.Path = result.Oss_path
	model.ModDate = result.Mod_date
	err := model.Create()
	if err != nil {
		fmt.Printf("发生错误:%s",err)
	}
}
