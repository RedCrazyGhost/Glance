/**
 @author: RedCrazyGhost
 @date: 2023/2/13

**/

package core

type TradeType uint8

const (
	Decrease TradeType = iota
	Increase
)

func (t TradeType) String() string {
	switch t {
	case Increase:
		return "增加(+)"
	case Decrease:
		return "减少(-)"
	default:
		return "未知错误！请联系作者：RedCrazyGhost->redcrazyghost@163.com"
	}
}
