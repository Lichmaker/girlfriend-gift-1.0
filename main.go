package main

import (
	"fmt"
	v1 "lichmaker/girlfriend-gift-1/app/http/controllers/v1"
	"lichmaker/girlfriend-gift-1/app/oss_scanner"
	"lichmaker/girlfriend-gift-1/bootstrap"
	"lichmaker/girlfriend-gift-1/config"
	configpkg "lichmaker/girlfriend-gift-1/pkg/config"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	config.Initialize()
	bootstrap.SetupDB()

	// a := []string{"1", "2"}
	// println(len(a))
	// b := "abc"
	// if len(b) > 0 {
	// 	println(len(b))
	// 	fmt.Println(b[0])
	// }

	// fmt.Println(config.Get("database.mysql.host"))

	go func() {
		var Api = v1.ApiGroupApp.IndexApiGroup
		port := configpkg.Viper.GetString("app.port")
		r := gin.Default()
		r.GET("/get", Api.IndexApi.Get)
		r.Run(":" + port)
	}()

	// 写一个简易的ticker，每24个小时执行一次
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "执行oss_scanner.Run", configpkg.Viper.GetString("OSS_SCAN_DIR"))
	oss_scanner.Run(configpkg.Viper.GetString("OSS_SCAN_DIR"))
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for range ticker.C {
			fmt.Println(time.Now().Format("2006-01-02 15:04:05"), "执行oss_scanner.Run", configpkg.Viper.GetString("OSS_SCAN_DIR"))
			oss_scanner.Run(configpkg.Viper.GetString("OSS_SCAN_DIR"))
		}
	}()

	select {}
}
