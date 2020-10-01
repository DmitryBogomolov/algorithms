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

func get(node *trieNode, key []rune, symbolIdx int, alphabet Alphabet) *trieNode {
	if node == nil {
		return nil
	}
	if symbolIdx == len(key) {
		return node
	}
	nodeIdx := alphabet.ToIndex(key[symbolIdx])
	return get(node.nodes[nodeIdx], key, symbolIdx+1, alphabet)
}

// Get finds value for a key.
func (trie *Trie) Get(key string) int {
	node := get(trie.root, []rune(key), 0, trie.alphabet)
	if node == nil {
		return NoValue
	}
	return node.value
}

func getNodeIdx(alphabet Alphabet, key []rune, idx int) int {
	return alphabet.ToIndex(key[idx])
}

func put(node *trieNode, key []rune, symbolIdx int, val int, alphabet Alphabet) *trieNode {
	if node == nil {
		node = newNode(alphabet)
	}
	if symbolIdx == len(key) {
		node.value = val
		return node
	}
	nodeIdx := getNodeIdx(alphabet, key, symbolIdx)
	node.nodes[nodeIdx] = put(node.nodes[nodeIdx], key, symbolIdx+1, val, alphabet)
	return node
}

// Put add key-value pair.
func (trie *Trie) Put(key string, val int) {
	trie.root = put(trie.root, []rune(key), 0, val, trie.alphabet)
}

func del(node *trieNode, key []rune, symbolIdx int, alphabet Alphabet) *trieNode {
	if node == nil {
		return nil
	}
	if symbolIdx == len(key) {
		node.value = NoValue
	} else {
		nodeIdx := getNodeIdx(alphabet, key, symbolIdx)
		node.nodes[nodeIdx] = del(node.nodes[nodeIdx], key, symbolIdx+1, alphabet)
	}
	if node.value != NoValue {
		return node
	}
	for i := 0; i < alphabet.Size(); i++ {
		if node.nodes[i] != nil {
			return node
		}
	}
	return nil
}

// Del removes a key.
func (trie *Trie) Del(key string) {
	trie.root = del(trie.root, []rune(key), 0, trie.alphabet)
}
