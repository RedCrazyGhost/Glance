/**
 @author: RedCrazyGhost
 @date: 2023/2/15

**/

package core

import (
	"github.com/spf13/viper"
)

// Parser 解析模版接口
// layout 时间解析样本
// matchingIndex 匹配数组所在的index位置
// todo 工厂模式 ？ config
type Parser struct {
	Layout   string
	Relation map[string]interface{}
}

func NewParser(fileName string) *Parser {
	parserViper := viper.New()
	parserViper.SetConfigFile("./parser/" + fileName + ".yaml")
	err := parserViper.ReadInConfig()
	if err != nil {
		SystemLogPool.Println("配置文件获取错误：", err, "正确路径："+fileName)
		return nil
	}
	parser := &Parser{}
	if err := parserViper.Unmarshal(&parser); err != nil {
		SystemLogPool.Println(fileName, "yaml配置解析失败：", err)
		return nil
	}
	GloabalLog.Println(fileName, "解析配置解析完毕！")

	return parser
}

// setRelation 设置匹配对应关系
func (p *Parser) setRelation(m map[string]interface{}) {
	p.Relation = m
}

// getLayout 获取日期时间模版
func (p *Parser) getLayout() string {
	return p.Layout
}

// getRelation 获取匹配对应关系
func (p *Parser) getRelation() map[string]interface{} {
	return p.Relation
}
