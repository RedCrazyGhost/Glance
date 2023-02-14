/**
 @author: RedCrazyGhost
 @date: 2023/2/13

**/

package core

// TradeChannel 交易渠道
type TradeChannel uint8

const (
	unknown TradeChannel = iota
	Alipay
	WeChatpay
	Tenpay
	BOF
	CCB
)

// todo 思考 是否要常驻内存，完成建树后回收map内存
var channelMap = map[TradeChannel]string{
	Alipay:    "支付宝",
	WeChatpay: "微信支付",
	Tenpay:    "财付通",
	BOF:       "中国银行",
	CCB:       "中国建设银行",
}

func (t TradeChannel) String() string {
	return channelMap[t]
}
