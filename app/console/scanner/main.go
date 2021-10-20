package main

import (
	"lichmaker/girlfriend-gift-1/bootstrap"
	"lichmaker/girlfriend-gift-1/config"
	configPkg "lichmaker/girlfriend-gift-1/pkg/config"
	"lichmaker/girlfriend-gift-1/pkg/uploader"
	"os"
	"path/filepath"
)

func main() {
	config.Initialize()
	bootstrap.SetupDB()

	root := configPkg.Viper.GetString("app.albumPath")
	if len(root) == 0 {
		panic("ALBUM_PATH 配置为空")
	}
	var pathArray []string
	pathData := make(map[string]string)
	_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		pathArray = append(pathArray, path)
		pathData[path] = info.ModTime().Format("2006-01-02")
		return nil
	})
	uploader.Do(pathData)
}
