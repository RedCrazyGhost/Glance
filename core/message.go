/**
 @author: RedCrazyGhost
 @date: 2023/2/15

**/

package core

import "fmt"

type Message struct {
	Mod string
	Obj interface{}
	Msg string
}

// NewMessage 消息模版
func NewMessage(Mod, Obj, Msg string) *Message {
	return &Message{Mod: Mod, Obj: Obj, Msg: Msg}
}

// NewLogMessage 日志消息模版
func NewLogMessage(Obj, Msg string) *Message {
	return NewMessage("日志/log", Obj, Msg)
}

// NewAppMessage 应用消息模版
func NewAppMessage(Obj, Msg string) *Message {
	return NewMessage("应用/Application", Obj, Msg)
}

// NewAppMessage 交易节点消息模版
func NewTradeNodeMessage(Obj, Msg string) *Message {
	return NewMessage("交易节点/TradeNode", Obj, Msg)
}

func (c Message) String() string {
	return fmt.Sprintf("Mod:%s Obj:%v Msg:%s", c.Mod, c.Obj, c.Msg)
}
