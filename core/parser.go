/**
 @author: RedCrazyGhost
 @date: 2023/2/15

**/

package core

type Parser interface {
	layout() string
	matchingIndex() map[string]interface{}
}

type AlipayParser struct {
}

// layout 时间解析样本
func (p AlipayParser) layout() string {
	return "2006/1/2 15:04"
}

// matchingIndex 匹配数组所在的index位置
func (p AlipayParser) matchingIndex() map[string]interface{} {
	m := make(map[string]interface{})
	m["Annotation"] = 3
	m["DateTime"] = 4
	m["Target"] = 5
	m["TradeChannel"] = 9
	m["TradeAmount"] = []int{6, 7}
	return m
}

// todo 通过反射执行不同的方法完成赋值
