/**
 @author: RedCrazyGhost
 @date: 2023/2/18

**/

package core

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port       string
	SpacedText string
}

var SystemConfig *Config

// initConfig 配置初始化
func InitConfig() {

	GloabalLog.Println("开始配置初始化！")
	SystemViper := viper.New()
	SystemViper.SetConfigFile("./config.yaml")
	err := SystemViper.ReadInConfig()
	if err != nil {
		SystemLogPool.Println("配置文件获取错误：", err, "正确路径：./config.yaml")
	}
	if err := SystemViper.Unmarshal(&SystemConfig); err != nil {
		SystemLogPool.Println("yaml配置解析失败：", err)
	}

	GloabalLog.Println("配置初始化完毕！")
}
