/**
 @author: RedCrazyGhost
 @date: 2023/2/15

**/

package core

type Parser interface {
	layout() string
	matchingIndex() map[string]int
}

type AlipayParser struct {
}

// layout 时间解析样本
func (p AlipayParser) layout() string {
	return "2006/1/2 15:04"
}

// matchingIndex 匹配数组所在的index位置
func (p AlipayParser) matchingIndex() map[string]int {
	m := make(map[string]int)
	m["Annotation"] = 3
	m["DateTime"] = 4
	m["Target"] = 5
	m["TradeChannel"] = 9
	return m
}

// todo 通过反射执行不同的方法完成赋值
