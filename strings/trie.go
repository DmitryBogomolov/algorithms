package strings

// Trie is a search struct.
// https://algs4.cs.princeton.edu/52trie/TrieST.java.html
type Trie struct {
	root     *_TrieNode
	alphabet Alphabet
	size     int
}

type _TrieNode struct {
	value interface{}
	nodes []*_TrieNode
}

func (trie *Trie) symbolsCount() int {
	return trie.alphabet.Size()
}

func (trie *Trie) newNode() *_TrieNode {
	return &_TrieNode{
		value: nil,
		nodes: make([]*_TrieNode, trie.symbolsCount()),
	}
}

// NewTrie constructs trie instance.
func NewTrie(alphabet Alphabet) *Trie {
	return &Trie{
		root:     nil,
		alphabet: alphabet,
		size:     -1,
	}
}

// NewTrieASCII constructs trie with ASCII alphabet.
func NewTrieASCII() *Trie {
	return NewTrie(ASCIIAlphabet)
}

func (trie *Trie) sizeCore(node *_TrieNode) int {
	if node == nil {
		return 0
	}
	count := 0
	if node.value != nil {
		count++
	}
	for i := 0; i < trie.symbolsCount(); i++ {
		count += trie.sizeCore(node.nodes[i])
	}
	return count
}

// Size returns amount of elements in a trie.
func (trie *Trie) Size() int {
	if trie.size == -1 {
		trie.size = trie.sizeCore(trie.root)
	}
	return trie.size
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
	trie.size = -1
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
	for i := 0; i < trie.symbolsCount(); i++ {
		if node.nodes[i] != nil {
			return node
		}
	}
	return nil
}

// Del removes a key.
func (trie *Trie) Del(key string) {
	trie.size = -1
	trie.root = trie.delCore(trie.root, []rune(key), 0)
}

func (trie *Trie) idxToSymbol(idx int) rune {
	return trie.alphabet.ToSymbol(idx)
}

func (trie *Trie) keysWithPrefixCore(node *_TrieNode, prefix []rune, collection *[]string) {
	if node == nil {
		return
	}
	if node.value != nil {
		*collection = append(*collection, string(prefix))
	}
	for i := 0; i < trie.symbolsCount(); i++ {
		trie.keysWithPrefixCore(node.nodes[i], append(prefix, trie.idxToSymbol(i)), collection)
	}
}

// KeysWithPrefix collects keys with *prefix*.
func (trie *Trie) KeysWithPrefix(prefix string) []string {
	var collection []string
	trie.keysWithPrefixCore(trie.getCore(trie.root, []rune(prefix), 0), []rune(prefix), &collection)
	return collection
}

// Keys returns all keys.
func (trie *Trie) Keys() []string {
	return trie.KeysWithPrefix("")
}

func (trie *Trie) keysThatMatchCore(node *_TrieNode, prefix []rune, pattern []rune, collection *[]string) {
	if node == nil {
		return
	}
	if len(prefix) == len(pattern) {
		if node.value != nil {
			*collection = append(*collection, string(prefix))
		}
		return
	}
	nextSymbol := pattern[len(prefix)]
	for i := 0; i < trie.symbolsCount(); i++ {
		if nextSymbol == '.' || nextSymbol == trie.idxToSymbol(i) {
			trie.keysThatMatchCore(node.nodes[i], append(prefix, trie.idxToSymbol(i)), pattern, collection)
		}
	}
}

// KeysThatMatch collects keys matching *pattern*.
func (trie *Trie) KeysThatMatch(pattern string) []string {
	var collection []string
	trie.keysThatMatchCore(trie.root, nil, []rune(pattern), &collection)
	return collection
}

func (trie *Trie) longestPrefixCore(node *_TrieNode, str []rune, symbolIdx int, length int) int {
	if node == nil {
		return length
	}
	if node.value != nil {
		length = symbolIdx
	}
	if symbolIdx == len(str) {
		return length
	}
	nodeIdx := trie.symbolToIdx(str, symbolIdx)
	return trie.longestPrefixCore(node.nodes[nodeIdx], str, symbolIdx+1, length)
}

// LongestPrefix returns longest key that is prefix for *str*.
func (trie *Trie) LongestPrefix(str string) string {
	len := trie.longestPrefixCore(trie.root, []rune(str), 0, 0)
	return str[:len]
}
