package config

import "lichmaker/girlfriend-gift-1/pkg/config"

func init() {
	config.Viper.Set("app", map[string]interface{}{
		"port":	config.Get("PORT", "8080"),
		"albumPath": config.Get("ALBUM_PATH"),
		"debug": config.Get("DEBUG", "FALSE"),
	})
}
