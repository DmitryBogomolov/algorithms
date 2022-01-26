package strings

// Trie is a search struct.
// https://algs4.cs.princeton.edu/52trie/TrieST.java.html
type Trie struct {
	root     *_TrieNode
	alphabet Alphabet
}

type _TrieNode struct {
	value interface{}
	nodes []*_TrieNode
}

func (trie *Trie) newNode() *_TrieNode {
	return &_TrieNode{
		value: nil,
		nodes: make([]*_TrieNode, trie.alphabet.Size()),
	}
}

// NewTrie constructs trie instance.
func NewTrie(alphabet Alphabet) *Trie {
	return &Trie{
		root:     nil,
		alphabet: alphabet,
	}
}

func (trie *Trie) sizeCore(node *_TrieNode) int {
	if node == nil {
		return 0
	}
	count := 0
	if node.value != nil {
		count++
	}
	for i := 0; i < trie.alphabet.Size(); i++ {
		count += trie.sizeCore(node.nodes[i])
	}
	return count
}

// Size returns amount of elements in a trie.
func (trie *Trie) Size() int {
	return trie.sizeCore(trie.root)
}

func (trie *Trie) symbolToIdx(key []rune, idx int) int {
	return trie.alphabet.ToIndex(key[idx])
}

func (trie *Trie) getCore(node *_TrieNode, key []rune, symbolIdx int) *_TrieNode {
	if node == nil {
		return nil
	}
	if symbolIdx == len(key) {
		return node
	}
	nodeIdx := trie.symbolToIdx(key, symbolIdx)
	return trie.getCore(node.nodes[nodeIdx], key, symbolIdx+1)
}

// Get finds a value for a key.
func (trie *Trie) Get(key string) interface{} {
	node := trie.getCore(trie.root, []rune(key), 0)
	if node == nil {
		return nil
	}
	return node.value
}

func (trie *Trie) putCore(node *_TrieNode, key []rune, symbolIdx int, val interface{}) *_TrieNode {
	if node == nil {
		node = trie.newNode()
	}
	if symbolIdx == len(key) {
		node.value = val
		return node
	}
	nodeIdx := trie.symbolToIdx(key, symbolIdx)
	node.nodes[nodeIdx] = trie.putCore(node.nodes[nodeIdx], key, symbolIdx+1, val)
	return node
}

// Put sets a value for a key.
func (trie *Trie) Put(key string, val interface{}) {
	trie.root = trie.putCore(trie.root, []rune(key), 0, val)
}

func (trie *Trie) delCore(node *_TrieNode, key []rune, symbolIdx int) *_TrieNode {
	if node == nil {
		return nil
	}
	if symbolIdx == len(key) {
		node.value = nil
	} else {
		nodeIdx := trie.symbolToIdx(key, symbolIdx)
		node.nodes[nodeIdx] = trie.delCore(node.nodes[nodeIdx], key, symbolIdx+1)
	}
	if node.value != nil {
		return node
	}
	for i := 0; i < trie.alphabet.Size(); i++ {
		if node.nodes[i] != nil {
			return node
		}
	}
	return nil
}

// Del removes a key.
func (trie *Trie) Del(key string) {
	trie.root = trie.delCore(trie.root, []rune(key), 0)
}

func (trie *Trie) idxToSymbol(idx int) rune {
	return trie.alphabet.ToSymbol(idx)
}

func (trie *Trie) keysWithPrefixCore(node *_TrieNode, prefix string, collection *[]string) {
	if node == nil {
		return
	}
	if node.value != nil {
		*collection = append(*collection, prefix)
	}
	for i := 0; i < trie.alphabet.Size(); i++ {
		trie.keysWithPrefixCore(node.nodes[i], prefix+string(trie.idxToSymbol(i)), collection)
	}
}

// KeysWithPrefix collects keys with *prefix*.
func (trie *Trie) KeysWithPrefix(prefix string) []string {
	var collection []string
	trie.keysWithPrefixCore(trie.getCore(trie.root, []rune(prefix), 0), prefix, &collection)
	return collection
}

// Keys returns all keys.
func (trie *Trie) Keys() []string {
	return trie.KeysWithPrefix("")
}

func (trie *Trie) keysThatMatchCore(node *_TrieNode, prefix string, pattern string, collection *[]string) {
	if node == nil {
		return
	}
	if len(prefix) == len(pattern) {
		if node.value != nil {
			*collection = append(*collection, prefix)
		}
		return
	}
	nextSymbol := rune(pattern[len(prefix)])
	for i := 0; i < trie.alphabet.Size(); i++ {
		if nextSymbol == '.' || nextSymbol == trie.idxToSymbol(i) {
			trie.keysThatMatchCore(node.nodes[i], prefix+string(trie.idxToSymbol(i)), pattern, collection)
		}
	}
}

// KeysThatMatch collects keys matching *pattern*.
func (trie *Trie) KeysThatMatch(pattern string) []string {
	var collection []string
	trie.keysThatMatchCore(trie.root, "", pattern, &collection)
	return collection
}

func (trie *Trie) longestPrefixCore(node *_TrieNode, str string, symbolIdx int, length int) int {
	if node == nil {
		return length
	}
	if node.value != nil {
		length = symbolIdx
	}
	if symbolIdx == len(str) {
		return length
	}
	nodeIdx := trie.symbolToIdx([]rune(str), symbolIdx)
	return trie.longestPrefixCore(node.nodes[nodeIdx], str, symbolIdx+1, length)
}

// LongestPrefix returns longest key that is prefix for *str*.
func (trie *Trie) LongestPrefix(str string) string {
	len := trie.longestPrefixCore(trie.root, str, 0, 0)
	return str[:len]
}
