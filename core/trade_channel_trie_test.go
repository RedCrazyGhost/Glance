/**
 @author: RedCrazyGhost
 @date: 2023/2/13

**/

package core

import (
	"fmt"
	"testing"
)

func TestNewTrie(t *testing.T) {
	trie := NewChannelTrie()
	for key, value := range channelMap {
		trie.insert(value, key)
	}
	channels := trie.find("财付支付宝0134微信支付-支付宝-中国建设银行")
	fmt.Printf("%v\n", channels)
}
