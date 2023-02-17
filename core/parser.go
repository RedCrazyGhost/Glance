/**
 @author: RedCrazyGhost
 @date: 2023/2/15

**/

package core

// Parser 解析模版接口
// layout 时间解析样本
// matchingIndex 匹配数组所在的index位置
type Parser interface {
	layout() string
	matchingIndex() map[string]interface{}
}

// AlipayParser 支付宝账单解析模版
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

// CCBParser 中国建设银行账单解析模版
type CCBParser struct {
}

// layout 时间解析样本
func (p CCBParser) layout() string {
	return "20060102"
}

// matchingIndex 匹配数组所在的index位置
func (p CCBParser) matchingIndex() map[string]interface{} {
	m := make(map[string]interface{})
	m["Annotation"] = []string{"序号", "对方账号与户名"}
	m["DateTime"] = "交易日期"
	m["Target"] = 8
	m["TradeChannel"] = []string{"交易地点/附言", "对方账号与户名"}
	m["TradeAmount"] = 5
	return m
}
