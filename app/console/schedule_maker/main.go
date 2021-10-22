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

	// 优先找所有往年同一天的
	var queryData []result
	dateArray := getDateArray()
	for _, date := range dateArray {
		var queryTmp []result
		sql := "SELECT md5, oss_path, mod_date FROM pool WHERE mod_date = ?"
		model.DB.Raw(sql, date).Scan(&queryTmp)
		queryData = append(queryData, queryTmp...)
	}

	if len(queryData) > 0 {
		// 存在同日期数据， 就只写入同日期的数据
		write(dateString, queryData)
	} else {
		// 不存在同日期数据，就随机找1个图插进去
		// 最多跑10次，尽量找没重复使用的
		var md5 string
		var queryResult result
		for i := 0; i < 10; i++ {
			sql := "SELECT md5, oss_path, mod_date FROM pool ORDER BY rand() LIMIT 1"
			model.DB.Raw(sql).Scan(&queryResult)
			md5 = queryResult.Md5
			if schedule.Exists(md5) {
				continue
			} else {
				break
			}
		}
		write(dateString, []result{
			queryResult,
		})
	}
}

func getDateArray() []string {
	var data []string
	date := "1990-" + time.Now().Local().Format("01-02")
	// 直接简单粗暴从1990年起拿100年的日期
	for i := 0; i < 100; i++ {
		parse, _ := time.Parse("2006-01-02", date)
		parse = parse.AddDate(1, 0, 0)
		date = parse.Format("2006-01-02")
		data = append(data, date)
	}
	return data
}

func write(dateString string, data []result) {
	for _, r := range data {
		var model schedule.Schedule
		model.Md5 = r.Md5
		model.Date = dateString
		model.Path = r.Oss_path
		model.ModDate = r.Mod_date
		err := model.Create()
		if err != nil {
			fmt.Printf("发生错误:%s", err)
		}
	}
}
