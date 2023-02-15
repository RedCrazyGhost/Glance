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
	Head *TradeNode
	End  *TradeNode
}

// NewDefaultTradeFlow 创建默认交易流
// todo 重写内容
func NewDefaultTradeFlow() *TradeFlow {
	tradeNode := NewTradeNode(0.0)
	return &TradeFlow{tradeNode, tradeNode}
}

// NewTradeFlow 根据cash创建初试交易流
// todo 重写内容
func NewTradeFlow(cash float64) *TradeFlow {
	tradeNode := NewTradeNode(cash)
	return &TradeFlow{tradeNode, tradeNode}
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

	Parser *Parser // 解析模版
	Meta   string  // 元数据

	Last *TradeNode
	Next *TradeNode
}

// NewTradeNode 创建交易节点
// todo 完成初始化函数
func NewTradeNode(Balance       float64 ,
Cash          float64        ,
Datetime      time.Time     ,
TradeType     TradeType     ,
TradeAmount   float64       ,
TradeChannels []TradeChannel ,
TradeTags     []TradeTag    ,
Annotation    string     ) *TradeNode {
	node := TradeNode{
		Cash: ,
	}
	node.setId()

	return &node
}

// setCash 设置现金
// 用于建立链表的第一个节点
func (n *TradeNode) setCash(cash float64) {
	n.Cash = cash
}

// setTime 设置时间
// 解析时间字符串用于TradeNode
func (n *TradeNode) setdateTime(layout string) error {
	return nil
}

// setId 设置唯一标识
// 使用hash算法
func (n *TradeNode) setId() {
	bytes := md5.Sum([]byte(n.Meta))
	n.Id = fmt.Sprint(bytes)
}

func (n *TradeNode) parser() {

}
