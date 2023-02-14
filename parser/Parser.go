/**
 @author: RedCrazyGhost
 @date: 2023/2/14

**/

package parser

import "Glance/core"

type Generator interface {
	parser() *core.TradeNode
}
