/**
 @author: RedCrazyGhost
 @date: 2023/2/15

**/

package core

import "reflect"

type Parser interface {
	parserArrayIndex() map[*reflect.StructField]int
}

type AlipayParser struct {
}

// todo 编写解析内容
func (p AlipayParser) parserArrayIndex() {
	typeOfTradeNode := reflect.TypeOf(TradeNode{})

	name, b := typeOfTradeNode.FieldByName("Annotation")
	name.Index

	//datatimestr := metastr[4]
	//targetstr := metastr[5]
	//incomestr := metastr[6]
	//expenditurestr := metastr[7]
	//blancestr := metastr[8]
	//channelstr := metastr[9]
	//tagstr := metastr[10]
	//annotationstr := metastr[11]

	NewTradeNode()
	return nil
}
