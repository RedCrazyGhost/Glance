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
	alipay := NewParser("AlipayParser")
	titles := []string{"账务流水号", "业务流水号", "商户订单号", "商品名称", "发生时间", "对方账号", "收入金额（+元）", "支出金额（-元）", "账户余额（元）", "交易渠道", "业务类型", "备注"}
	node0, err := NewTradeNode(alipay, titles, []string{"1**********1", "2**************9", "4**********5", "收钱码收款0", "1949/10/1 00:00", "RedCrazyGhost(redcrazyghost@163.com)", "1", "0", "1", "支付宝", "在线支付", ""}...)
	if err != nil {
		return
	}
	node1, err := NewTradeNode(alipay, titles, []string{"1**********1", "2**************9", "4**********5", "收钱码收款1", "1949/10/1 00:00", "RedCrazyGhost(redcrazyghost@163.com)", "2", "0", "3", "支付宝", "在线支付", ""}...)
	if err != nil {
		return
	}
	node2, err := NewTradeNode(alipay, titles, []string{"1**********1", "2**************9", "4**********5", "收钱码收款2", "1949/10/1 00:00", "RedCrazyGhost(redcrazyghost@163.com)", "0", "-3", "0", "支付宝", "在线支付", ""}...)
	if err != nil {
		return
	}

	err = flow.insert(node0, 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = flow.insert(node1, 1)
	if err != nil {
		fmt.Println(err.Error())
	}
	err = flow.insert(node2, -1)
	if err != nil {
		fmt.Println(err.Error())
	}

	flow.Show()
}
