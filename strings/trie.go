package strings

type trieNode struct {
	value int
	nodes []*trieNode
}

// Trie is search struct.
// https://algs4.cs.princeton.edu/52trie
type Trie struct {
	root     *trieNode
	alphabet Alphabet
}

// NoValue is placeholder for empty nodes in a tree.
const NoValue = -1

func newNode(alphabet Alphabet) *trieNode {
	size := alphabet.Size()
	return &trieNode{
		value: NoValue,
		nodes: make([]*trieNode, size, size),
	}
}

// NewTrie constructs trie instance.
func NewTrie(alphabet Alphabet) *Trie {
	trie := Trie{
		root:     nil,
		alphabet: alphabet,
	}
	return &trie
}

func (trie *Trie) size(node *trieNode) int {
	if node == nil {
		return 0
	}
	count := 0
	if node.value != NoValue {
		count++
	}
	for i := 0; i < trie.alphabet.Size(); i++ {
		count += trie.size(node.nodes[i])
	}
	return count
}

// Size returns amount of elements.
func (trie *Trie) Size() int {
	return trie.size(trie.root)
}

func (trie *Trie) get(node *trieNode, key []rune, symbolIdx int) *trieNode {
	if node == nil {
		return nil
	}
	if symbolIdx == len(key) {
		return node
	}
	nodeIdx := trie.alphabet.ToIndex(key[symbolIdx])
	return trie.get(node.nodes[nodeIdx], key, symbolIdx+1)
}

// Get finds value for a key.
func (trie *Trie) Get(key string) int {
	node := trie.get(trie.root, []rune(key), 0)
	if node == nil {
		return NoValue
	}
	return node.value
}

func getNodeIdx(alphabet Alphabet, key []rune, idx int) int {
	return alphabet.ToIndex(key[idx])
}

func (trie *Trie) put(node *trieNode, key []rune, symbolIdx int, val int) *trieNode {
	if node == nil {
		node = newNode(trie.alphabet)
	}
	if symbolIdx == len(key) {
		node.value = val
		return node
	}
	nodeIdx := getNodeIdx(trie.alphabet, key, symbolIdx)
	node.nodes[nodeIdx] = trie.put(node.nodes[nodeIdx], key, symbolIdx+1, val)
	return node
}

// Put add key-value pair.
func (trie *Trie) Put(key string, val int) {
	trie.root = trie.put(trie.root, []rune(key), 0, val)
}

func (trie *Trie) del(node *trieNode, key []rune, symbolIdx int) *trieNode {
	if node == nil {
		return nil
	}
	if symbolIdx == len(key) {
		node.value = NoValue
	} else {
		nodeIdx := getNodeIdx(trie.alphabet, key, symbolIdx)
		node.nodes[nodeIdx] = trie.del(node.nodes[nodeIdx], key, symbolIdx+1)
	}
	if node.value != NoValue {
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
	trie.root = trie.del(trie.root, []rune(key), 0)
}
