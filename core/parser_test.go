/**
 @author: RedCrazyGhost
 @date: 2023/2/16

**/

package core

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAlipayParse(t *testing.T) {
	typeOfTradeNode := reflect.TypeOf(TradeNode{})

	for i := 0; i < typeOfTradeNode.NumField(); i++ {
		fmt.Println(typeOfTradeNode.Field(i))
	}

}

// TestNewParser  日志会导致无法地址报错
func TestNewParser(t *testing.T) {
	parser := NewParser("AlipayParser")
	fmt.Println(parser)
}
