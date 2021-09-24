package config

import (
	"github.com/spf13/viper"
)

var Viper *viper.Viper

type ConfigMap map[string]interface{}

func init() {
	// 1. 初始化 Viper 库
	Viper = viper.New()
	// 2. 设置文件名称
	Viper.SetConfigName(".env")
	// 3. 配置类型，支持 "json", "toml", "yaml", "yml", "properties",
	//             "props", "prop", "env", "dotenv"
	Viper.SetConfigType("env")
	// 4. 环境变量配置文件查找的路径，相对于 main.go
	Viper.AddConfigPath(".")

	// 5. 开始读根目录下的 .env 文件，读不到会报错
	err := Viper.ReadInConfig()
	if err != nil {
		panic(err)
	}

	// 6. 设置环境变量前缀，用以区分 Go 的系统环境变量
	Viper.SetEnvPrefix("appenv")
	// 7. Viper.Get() 时，优先读取环境变量
	Viper.AutomaticEnv()
}

// Get 获取配置项，允许使用点式获取，如：app.name
func Get(path string, defaultValue ...interface{}) interface{} {
	// 不存在的情况
	if !Viper.IsSet(path) {
		if len(defaultValue) > 0 {
			return defaultValue[0]
		}
		return nil
	}
	return Viper.Get(path)
}
