package config

import (
	"github.com/spf13/viper"
)

// Load 加载配置文件
func Load(filename string) error {
	viper.SetConfigFile(filename)
	return viper.ReadInConfig()
}

// Get 获取指定配置项
func Get(index string) interface{} {
	return viper.Get(index)
}

// GetString 获取指定配置项
func GetString(index string) string {
	return viper.GetString(index)
}

// GetStringSlice 获取指定配置项
func GetStringSlice(index string) []string {
	return viper.GetStringSlice(index)
}
