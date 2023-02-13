/**
 @author: RedCrazyGhost
 @date: 2023/2/13

**/

package core

import (
	"crypto/md5"
	"fmt"
	"time"
)

// TradeFlow 交易流
type TradeFlow struct {
	head *TradeNode
}

// NewDefaultTradeFlow 创建默认交易流
func NewDefaultTradeFlow() *Trade {
	tradeNode := NewTradeNode(0.0)
	return &TradeFlow{tradeNode}
}

// NewTradeFlow 根据cash创建初试交易流
func NewTradeFlow(cash float64) *Trade {
	tradeNode := NewTradeNode(cash)
	return &TradeFlow{tradeNode}
}

// TradeNode 节点
// 用于财产流水记录的节点
type TradeNode struct {
	Id            string         // 唯一标识
	Balance       float64        // 上个节点操作后的余额
	Cash          float64        // 当前节点余额
	Datetime      time.Time      // 交易时间
	TradeType     TradeType      // 交易类型
	TradeAmount   float64        // 交易金额
	TradeChannels []TradeChannel // 交易渠道
	TradeTags     []TradeTag     // 交易标签
	Annotation    string         // 注解
	Meta          string         // 元数据

	Last *TradeNode
	Next *TradeNode
}

// NewTradeNode 创建交易节点
// todo 完善交易节点的创建
func NewTradeNode(cash float64) *TradeNode {
	tradeNode := TradeNode{
		Cash: cash,
	}

}

// setCash 设置现金
// 用于建立链表的第一个节点
func (n *TradeNode) setCash(cash float64) {
	n.Cash = cash
}

// setTime 设置时间
// 解析时间字符串用于TradeNode
// todo 开发 日期时间文本解析成Time.time
func (n *TradeNode) setTime(datetimestr string) {

}

// setId 设置唯一标识
// 使用hash算法
func (n *TradeNode) setId() {
	bytes := md5.Sum([]byte(n.Meta + n.Datetime.String()))
	n.Id = fmt.Sprint(bytes)
}
