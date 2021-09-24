package main

import (
	v1 "lichmaker/girlfriend-gift-1/app/http/controllers/v1"
	"lichmaker/girlfriend-gift-1/bootstrap"
	"lichmaker/girlfriend-gift-1/config"
	configpkg "lichmaker/girlfriend-gift-1/pkg/config"

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

	var Api = v1.ApiGroupApp.IndexApiGroup
	port := configpkg.Viper.GetString("app.port")
	r := gin.Default()
	r.GET("/get", Api.IndexApi.Get)
	r.Run(":" + port)
}
