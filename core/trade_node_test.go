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
	node := TradeNode{Parser: AlipayParser{}, Meta: []string{"1**********1", "2**************9", "4**********5", "收钱码收款", "1949/10/1 00:00", "RedCrazyGhost(redcrazyghost@163.com)", "0", "0", "0", "支付宝", "在线支付", ""}}
	err := node.setId()
	if err != nil {
		fmt.Println("setId错误")
	}
	fmt.Println(node)
}
