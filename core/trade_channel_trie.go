/**
 @author: RedCrazyGhost
 @date: 2023/2/13

**/

package core

var TradeChannelTrie *ChannelTrie

// ChannelTrie 支付渠道树
type ChannelTrie struct {
	root *ChannelTrieNode
}

// NewChannelTrie 创建字典树
func NewChannelTrie() *ChannelTrie {
	trieNode := NewChannelTrieNode('!', 0)
	return &ChannelTrie{trieNode}
}

// InitChannelTrie 初始化树
func InitChannelTrie() {
	TradeChannelTrie = NewChannelTrie()
	for key, value := range channelMap {
		TradeChannelTrie.insert(value, key)
	}
}

// ChannelTrieNode 字典树
type ChannelTrieNode struct {
	char         rune                      // 字符
	deep         uint8                     // 深度
	isEnding     bool                      // 结尾节点
	TradeChannel TradeChannel              // 支付渠道
	children     map[rune]*ChannelTrieNode // 子节点
}

// NewChannelTrieNode  创建字典树节点
func NewChannelTrieNode(char rune, deep uint8) *ChannelTrieNode {
	return &ChannelTrieNode{
		char:         char,
		isEnding:     false,
		deep:         deep,
		TradeChannel: unknown,
		children:     make(map[rune]*ChannelTrieNode),
	}
}

// insert 往字典树里添加单词
func (t *ChannelTrie) insert(word string, channel TradeChannel) {
	node := t.root
	for _, char := range word {
		value, ok := node.children[char]
		if !ok {
			value = NewChannelTrieNode(char, node.deep+1)
			node.children[char] = value
		}
		node = value
	}
	node.TradeChannel = channel
	node.isEnding = true
}

// find 寻找匹配名称的支付渠道名称
func (t *ChannelTrie) find(str string) []TradeChannel {
	runes := []rune(str)
	var channels []TradeChannel
	node := t.root
	for i := 0; i < len(runes); i++ {
		value, ok := node.children[runes[i]]
		if ok {
			node = value
			if node.isEnding {
				channels = append(channels, node.TradeChannel)
				node = t.root
			}
		}
	}
	return channels
}
