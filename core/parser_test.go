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

	/** 运行结果：
	{Id  string  0 [0] false}
	{LBalance  float64  16 [1] false}
	{Cash  float64  24 [2] false}
	{Datetime  time.Time  32 [3] false}
	{Target  string  56 [4] false}
	{TradeType  core.TradeType  72 [5] false}
	{TradeAmount  float64  80 [6] false}
	{TradeChannels  []core.TradeChannel  88 [7] false}
	{TradeTags  []core.TradeTag  112 [8] false}
	{Annotation  string  136 [9] false}
	{Parser  core.Parser  152 [10] false}
	{Meta  []string  168 [11] false}
	{Last  *core.TradeNode  192 [12] false}
	{Next  *core.TradeNode  200 [13] false}
	**/
}
