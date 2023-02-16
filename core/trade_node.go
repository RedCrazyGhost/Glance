/**
 @author: RedCrazyGhost
 @date: 2023/2/13

**/

package core

import (
	"crypto/md5"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"
)

// TradeFlow 交易流
type TradeFlow struct {
	Head *TradeNode
	End  *TradeNode
}

// NewDefaultTradeFlow 创建默认交易流
func NewDefaultTradeFlow() *TradeFlow {
	return &TradeFlow{nil, nil}
}

// TradeNode 节点
// 用于财产流水记录的节点
type TradeNode struct {
	Id            string         // 唯一标识 ✅
	LBalance      float64        // 交易前余额 ✅
	Cash          float64        // 交易后余额 ✅
	Datetime      time.Time      // 交易时间 ✅
	Target        string         // 交易目标 ✅
	TradeType     TradeType      // 交易类型 ✅
	TradeAmount   float64        // 交易金额 ✅
	TradeChannels []TradeChannel // 交易渠道 ✅
	TradeTags     []TradeTag     // 交易标签
	Annotation    string         // 注解 ✅

	Parser Parser   // 解析模版
	Meta   []string // 元数据

	Last *TradeNode
	Next *TradeNode
}

// NewTradeNode 创建交易节点
func NewTradeNode(parse Parser, metadata ...string) *TradeNode {
	node := &TradeNode{
		Parser: parse,
		Meta:   metadata,
	}
	err := node.parser()
	if err != nil {
		FailLog.Println("交易节点解析错误", "原因->"+err.Error())
	}
	return node
}

//setTradeChannelByMeta 设置交易渠道
func (n *TradeNode) setTradeChannelByMeta(index int) error {
	tradeChannelstr := n.Meta[index]
	if len(tradeChannelstr) > 0 {
		channels := TradeChannelTrie.find(tradeChannelstr)
		n.TradeChannels = channels
		return nil
	} else {
		return errors.New("交易渠道数据不存在！")
	}
}

// seTradeType 设置交易类型
func (n *TradeNode) setTradeType(tradeAmount float64) {
	if tradeAmount >= 0.0 {
		n.TradeType = Increase
	} else {
		n.TradeType = Decrease
	}
}

// setAnnotationByMeta 设置交易对象
func (n *TradeNode) setAnnotationByMeta(index int) error {
	annotation := n.Meta[index]
	if len(annotation) > 0 {
		n.Annotation = annotation
		return nil
	} else {
		return errors.New("交易注解数据不存在！")
	}
}

// setLBalanceByMeta 设置上个节点余额
func (n *TradeNode) setLBalanceByMeta(index int) error {
	balacncestr := n.Meta[index]
	if len(balacncestr) > 0 {
		// 处理科学计数法
		balance, err := strconv.ParseFloat(strings.ReplaceAll(balacncestr, ",", ""), 64)
		if err != nil {
			return errors.New("交易前余额数据解析错误！")
		}
		n.LBalance = balance
		return nil
	} else {
		return errors.New("交易前余额不存在！")
	}
}

// setTradeAmountByMeta 设置交易金额
// 用于建立链表的第一个节点
func (n *TradeNode) setTradeAmountByMeta(index int) error {
	tradeAmountstr := n.Meta[index]
	if len(tradeAmountstr) > 0 {
		// 处理科学计数法
		tradeAmount, err := strconv.ParseFloat(strings.ReplaceAll(tradeAmountstr, ",", ""), 64)
		if err != nil {
			return errors.New("交易后金额数据解析错误！")
		}
		n.TradeAmount = tradeAmount
		n.setTradeType(n.TradeAmount)
		return nil
	} else {
		return errors.New("交易后金额数据不存在！")
	}
}

// setCashByMeta 设置现金
// 用于建立链表的第一个节点
func (n *TradeNode) setCashByMeta(index int) error {
	cashstr := n.Meta[index]
	if len(cashstr) > 0 {
		// 处理科学计数法
		cash, err := strconv.ParseFloat(strings.ReplaceAll(cashstr, ",", ""), 64)
		if err != nil {
			return errors.New("交易后金额数据解析错误！")
		}
		n.Cash = cash
		return nil
	} else {
		return errors.New("交易后金额数据不存在！")
	}
}

// setTime 设置时间
// 解析时间字符串用于TradeNode
func (n *TradeNode) setDateTimeByMeta(index int) error {
	datetimestr := n.Meta[index]
	if len(datetimestr) > 0 {
		datetime, err := time.Parse(n.Parser.layout(), datetimestr)
		if err != nil {
			return errors.New("日期时间数据解析错误！")
		}
		n.Datetime = datetime
		return nil
	} else {
		return errors.New("日期时间数据不存在！")
	}

}

// setTargetByMeta 设置交易对象
func (n *TradeNode) setTargetByMeta(index int) error {
	target := n.Meta[index]
	if len(target) > 0 {
		n.Target = target
		return nil
	} else {
		return errors.New("交易对象数据不存在！")
	}
}

// setId 设置唯一标识
// 使用hash算法
func (n *TradeNode) setId() error {
	bytes := md5.Sum([]byte(fmt.Sprint(n.Meta)))
	n.Id = fmt.Sprintf("%v", bytes)
	return nil
}

// parser 对元数据进行解析
func (n *TradeNode) parser() error {
	if n.Meta == nil || len(n.Meta) == 0 {
		return errors.New("元数据不存在！")
	}
	if n.Parser == nil {
		return errors.New("解析器不存在！")
	}

	m := n.Parser.matchingIndex()
	err := n.setId()
	if err != nil {
		return err
	}

	for key, value := range m {
		switch key {
		case "LBalance":
			err = n.setLBalanceByMeta(value)
			if err != nil {
				return err
			}
			break
		case "Cash":
			err = n.setCashByMeta(value)
			if err != nil {
				return err
			}
			break
		case "DateTime":
			err = n.setDateTimeByMeta(value)
			if err != nil {
				return err
			}
			break
		case "Target":
			err = n.setTargetByMeta(value)
			if err != nil {
				return err
			}
			break
		case "TradeAmount":
			err = n.setTradeAmountByMeta(value)
			if err != nil {
				return err
			}
			break
		case "TradeChannel":
			err = n.setTradeChannelByMeta(value)
			if err != nil {
				return err
			}
			break
		case "Annotation":
			err = n.setAnnotationByMeta(value)
			if err != nil {
				return err
			}
			break
		default:
			break
		}
	}

	//err = n.setLBalanceByMeta(m["LBalance"])
	//if err != nil {
	//	return err
	//}
	//err = n.setCash(m["Cash"])
	//if err != nil {
	//	return err
	//}
	//err = n.setDateTimeByMeta(m["DateTime"])
	//if err != nil {
	//	return err
	//}
	//err = n.setTargetByMeta(m["Target"])
	//if err != nil {
	//	return err
	//}
	//err = n.setTradeAmountByMeta(m["TradeAmount"])
	//if err != nil {
	//	return err
	//}
	//err = n.setTradeChannelByMeta(m["TradeChannel"])
	//if err != nil {
	//	return err
	//}
	//err = n.setAnnotationByMeta(m["Annotation"])
	//if err != nil {
	//	return err
	//}

	return nil
}

func (n TradeNode) String() string {
	return fmt.Sprintf("唯一标识：%v\n交易前余额：%v\n交易后余额：%v\n交易日期：%v\n交易对象：%v\n交易类型：%v\n交易金额：%v\n交易渠道：%v\n交易标签：%v\n备注：%v\n",
		n.Id, n.LBalance, n.Cash, n.Datetime, n.Target, n.TradeType, n.TradeAmount, n.TradeChannels, n.TradeTags, n.Annotation)
}
