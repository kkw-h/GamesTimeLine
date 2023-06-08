//Package config
/*
@Title: config.go
@Description
@Author: kkw 2022/12/30 17:04
*/
package config

import (
	"encoding/json"
	"fmt"
	viperlib "github.com/spf13/viper"
	"go.kkw.top/gamesTimeLine/pkg"
)

var Config *viperlib.Viper

func LoadConfig(path string) {
	viper := viperlib.New()
	viper.SetConfigName("api")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viperlib.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
			panic("找不到配置文件")
		} else {
			// Config file was found but another error was produced
			fmt.Println(err)
			panic("配置文件格式错误")
		}
	}
	Config = viper
	prettyJSON, err := json.MarshalIndent(Config.AllSettings(), "", "  ")
	if err != nil {
		fmt.Println("Failed to marshal map:", err)
		return
	}
	pkg.KwLogger.Info("Config", string(prettyJSON))
}

func Get(key string) string {
	return Config.GetString(key)
}
