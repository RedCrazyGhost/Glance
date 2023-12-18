/**
 @author: RedCrazyGhost
 @date: 2023/2/18

**/

package core

import (
	"github.com/spf13/viper"
)

type Config struct {
	Run    string
	Server ServerConfig
	Shell  ShellConfig
	Base   BaseConfig
}

type ShellConfig struct {
	FilePath   string
	ParserName string
}
type BaseConfig struct {
	SpacedText string
	CSV        CSVConfig
}

type ServerConfig struct {
	Port string
}

type CSVConfig struct {
	Comment          string
	LazyQuotes       bool
	ReuseRecord      bool
	TrimLeadingSpace bool
}

var SystemConfig *Config

// InitConfig 配置初始化
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
	GloabalLog.Println(SystemConfig)
	GloabalLog.Println("配置初始化完毕！")
}
