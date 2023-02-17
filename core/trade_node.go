/**
 @author: RedCrazyGhost
 @date: 2023/2/13

**/

package core

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

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

	Parser     Parser         // 解析模版
	MetaTittle map[string]int // 元数据标题
	MetaData   []string       // 元数据

	Last *TradeNode
	Next *TradeNode
}

// NewTradeNode 创建交易节点
func NewTradeNode(parse Parser, metatitle []string, metadata ...string) (*TradeNode, error) {
	node := &TradeNode{
		Parser:   parse,
		MetaData: metadata,
	}
	err := node.setMetaTitle(metatitle...)
	if err != nil {
		FailLog.Println("交易节点ID设置错误", "原因->"+err.Error())
		return nil, err

	}
	err = node.parser()
	if err != nil {
		FailLog.Println(node.Id, "交易节点解析错误", "原因->"+err.Error())
		return nil, err

	}
	return node, nil
}

// countCash 计算当前余额
func (n *TradeNode) countCash() error {
	err := n.setLBalanceByMeta(nil)
	if err != nil {
		return err
	}
	err = n.setCashByMeta(nil)
	if err != nil {
		return err
	}
	node := n.Next
	for {
		if node != nil {
			err := node.setCashByMeta(nil)
			if err != nil {
				return err
			}
			if node.Next != nil {
				node = node.Next
			} else {
				break
			}
		} else {
			break
		}
	}
	return nil
}

// setMetaTitle 设置映射关系
func (n *TradeNode) setMetaTitle(metatitle ...string) error {
	if len(metatitle) > 0 {
		n.MetaTittle = make(map[string]int)
		for index, str := range metatitle {
			n.MetaTittle[str] = index
		}
		return nil
	} else {
		return errors.New("元数据表头所需的数据不存在！")
	}
}

// getMetaIndexByString 通过表头获取元数据角标
func (n *TradeNode) getMetaIndexByString(title string) int {
	if n.MetaTittle != nil && len(n.MetaTittle) > 0 {
		return n.MetaTittle[title]
	} else {
		return -1
	}
}

//setTradeChannelByMeta 设置交易渠道
func (n *TradeNode) setTradeChannelByMeta(index interface{}) error {
	tradeChannelstr, err := n.getDataString(index)
	if err != nil {
		return err
	}
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
func (n *TradeNode) setAnnotationByMeta(index interface{}) error {
	annotation, err := n.getDataString(index)
	if err != nil {
		return err
	}

	if len(annotation) > 0 {
		n.Annotation = annotation
		return nil
	} else {
		return errors.New("交易注解数据不存在！")
	}
}

// setLBalanceByMeta 设置上个节点余额
func (n *TradeNode) setLBalanceByMeta(index interface{}) error {
	if index == nil {
		if n.Last != nil {
			n.LBalance = n.Last.Cash
		}
		return nil
	}

	balacncestr, err := n.getDataString(index)
	if err != nil {
		return err
	}

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
func (n *TradeNode) setTradeAmountByMeta(index interface{}) error {
	v := 0.0
	str, err := n.getDataString(index)
	if err != nil {
		return err
	}
	strs := strings.Split(str, "\t")
	for _, str := range strs {
		tradeAmountstr := str
		if len(tradeAmountstr) > 0 {
			// 处理科学计数法
			tradeAmount, err := strconv.ParseFloat(strings.ReplaceAll(tradeAmountstr, ",", ""), 64)
			if err != nil {
				return errors.New("交易后金额数据解析错误！")
			}
			v += tradeAmount
		} else {
			return errors.New("交易后金额数据不存在！")
		}
	}
	n.TradeAmount = math.Abs(v)
	n.setTradeType(v)
	return nil
}

// setCashByMeta 设置现金
// 用于建立链表的第一个节点
func (n *TradeNode) setCashByMeta(index interface{}) error {
	if index == nil {
		if n.TradeType == Increase {
			n.Cash = n.LBalance + n.TradeAmount
		} else {
			n.Cash = n.LBalance - n.TradeAmount
		}
		return nil
	}

	cashstr, err := n.getDataString(index)
	if err != nil {
		return err
	}
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
func (n *TradeNode) setDateTimeByMeta(index interface{}) error {
	datetimestr, err := n.getDataString(index)
	if err != nil {
		return err
	}
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
func (n *TradeNode) setTargetByMeta(index interface{}) error {
	target, err := n.getDataString(index)
	if err != nil {
		return err
	}
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
	n.Id = MD5(nil, n.MetaData...)
	return nil
}

// parser 对元数据进行解析
func (n *TradeNode) parser() error {
	if n.MetaData == nil || len(n.MetaData) == 0 {
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

	return nil
}

// getData 获取对应的数据
// todo 修改参数名称
func (n *TradeNode) getDataString(index interface{}) (string, error) {
	var metastr string
	switch index.(type) {
	case int:
		metastr = n.MetaData[index.(int)]
		if len(metastr) <= 0 {
			return "", errors.New("元数据内容为空！")
		}
		break
	case []int:
		arr := make([]int, 0)
		arr = index.([]int)
		leng := len(arr)
		for i := 0; i < leng; i++ {
			metastr += n.MetaData[arr[i]]
			if len(metastr) <= 0 {
				return "", errors.New("元数据内容为空！")
			}
			if i != leng-1 {
				metastr += "\t"
			}
		}
		break
	case string:
		index := n.getMetaIndexByString(index.(string))
		if index >= 0 {
			metastr = n.MetaData[index]
			if len(metastr) <= 0 {
				return "", errors.New("元数据内容为空！")
			}
		} else {
			return "", errors.New("元数据表头角标数据不存在！")
		}
		break

	case []string:
		strs := make([]string, 0)
		strs = index.([]string)
		leng := len(strs)
		for i := 0; i < leng; i++ {
			index := n.getMetaIndexByString(strs[i])
			if index >= 0 {
				metastr += n.MetaData[index]
				if len(metastr) <= 0 {
					return "", errors.New("元数据内容为空！")
				}
				if i != leng-1 {
					metastr += "\t"
				}
			} else {
				return "", errors.New("元数据表头角标数据不存在！")
			}
		}
		break
	}

	return metastr, nil

}

func (n TradeNode) String() string {
	return fmt.Sprintf("唯一标识：%v\n交易前余额：%v\n交易后余额：%v\n交易日期：%v\n交易对象：%v\n交易类型：%v\n交易金额：%v\n交易渠道：%v\n交易标签：%v\n备注：%v\n",
		n.Id, n.LBalance, n.Cash, n.Datetime, n.Target, n.TradeType, n.TradeAmount, n.TradeChannels, n.TradeTags, n.Annotation)
}
