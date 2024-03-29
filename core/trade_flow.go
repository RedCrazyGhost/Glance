/**
 @author: RedCrazyGhost
 @date: 2023/2/17

**/

package core

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

// TradeFlow 交易流
// Left 最左节点，Right 最右节点
type TradeFlow struct {
	Left  *TradeNode
	Right *TradeNode
	len   int
}

// NewDefaultTradeFlow 创建默认交易流
func NewDefaultTradeFlow() *TradeFlow {
	return &TradeFlow{nil, nil, 0}
}

func (f *TradeFlow) CreateFile() {
	filePath := SystemConfig.Shell.FilePath[:strings.LastIndex(SystemConfig.Shell.FilePath, "/")] + "trade_flow.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	if f.Left == nil {
		return
	}
	node := f.Left
	for {
		write.WriteString(node.String() + "\n")
		if node.Next != nil {
			node = node.Next
		} else {
			return
		}
	}
	write.Flush()
}

func (f *TradeFlow) Show() {
	if f.Left == nil {
		return
	}
	node := f.Left
	for {
		fmt.Println(node)
		if node.Next != nil {
			node = node.Next
		} else {
			return
		}
	}
}

// insert 添加TradeNode
// direction 大于等于(>=)0向右添加，小于(<)0向左添加
func (f *TradeFlow) insert(node *TradeNode, direction int) error {
	if f.Left == nil && f.Right == nil {
		f.Left = node
		f.Right = node
		err := node.setCashByMeta(nil)
		if err != nil {
			return err
		}
	} else if direction >= 0 {
		temp := f.Right
		temp.Next = node
		node.Last = temp
		f.Right = node
		err := f.Right.countCash()
		if err != nil {
			return err
		}
	} else if direction < 0 {
		temp := f.Left
		temp.Last = node
		node.Next = temp
		f.Left = node
		err := f.Left.countCash()
		if err != nil {
			return err
		}
	} else {
		return errors.New("交易流添加节点失败！")
	}
	return nil
}

// Verify 流计算检测
// todo 存在精度丢失的问题
func (f *TradeFlow) Verify() bool {
	node := f.Left
	var flag bool
	for {
		if node != nil {
			if node.getTradeType(Increase) {
				if node.Cash-node.TradeAmount == node.LBalance {
					flag = true
				} else {
					flag = false
				}
			} else {
				if node.Cash+node.TradeAmount == node.LBalance {
					flag = true
				} else {
					flag = false
				}
			}
			if !flag {
				return flag
			}
			node = node.Next
		} else {
			break
		}
	}
	return flag
}

// Option 原始交易交易链 生成新的交易链
func (f TradeFlow) Option() *TradeFlow {
	return nil

}

func (f TradeFlow) String() string {
	return fmt.Sprintf("%v\n", f.Left)
}
