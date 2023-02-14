/**
 @author: RedCrazyGhost
 @date: 2023/2/14

**/

package parser

import "Glance/core"

// todo Parser解析怎么做
type Generator interface {
	parser() *core.TradeNode
}
