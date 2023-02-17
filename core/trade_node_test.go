/**
 @author: RedCrazyGhost
 @date: 2023/2/16

**/

package core

import (
	"fmt"
	"testing"
)

func TestNewDefaultTradeFlow(t *testing.T) {
	flow := NewDefaultTradeFlow()
	node0 := TradeNode{Id: "node0", Parser: AlipayParser{}, MetaData: []string{"1**********1", "2**************9", "4**********5", "收钱码收款", "1949/10/1 00:00", "RedCrazyGhost(redcrazyghost@163.com)", "0", "0", "0", "支付宝", "在线支付", ""}}
	node1 := TradeNode{Id: "node1", Parser: AlipayParser{}, MetaData: []string{"1**********1", "2**************9", "4**********5", "收钱码收款", "1949/10/1 00:00", "RedCrazyGhost(redcrazyghost@163.com)", "0", "0", "0", "支付宝", "在线支付", ""}}
	node2 := TradeNode{Id: "node2", Parser: AlipayParser{}, MetaData: []string{"1**********1", "2**************9", "4**********5", "收钱码收款", "1949/10/1 00:00", "RedCrazyGhost(redcrazyghost@163.com)", "0", "0", "0", "支付宝", "在线支付", ""}}

	err := flow.insert(&node0, 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = flow.insert(&node1, 1)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = flow.insert(&node2, -1)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(flow)
}
