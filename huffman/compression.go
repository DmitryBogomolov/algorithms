package huffman

import (
	"container/heap"
	"fmt"
	"strings"
)

type node struct {
	ch          byte
	freq        int
	left, right *node
}

func (n node) isLeaf() bool {
	return n.left == nil && n.right == nil
}

type nodesPriorityQueue []*node

func (pq nodesPriorityQueue) Len() int {
	return len(pq)
}

func (pq nodesPriorityQueue) Less(i, j int) bool {
	return pq[i].freq < pq[j].freq
}

func (pq nodesPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *nodesPriorityQueue) Push(item interface{}) {
	nodes := *pq
	node := item.(*node)
	*pq = append(nodes, node)
}

func (pq *nodesPriorityQueue) Pop() interface{} {
	nodes := *pq
	len := len(nodes)
	node := nodes[len-1]
	nodes[len-1] = nil
	*pq = nodes[0 : len-1]
	return node
}

func collectFrequencies(data []byte) map[byte]int {
	frequencies := make(map[byte]int)
	for _, ch := range data {
		frequencies[ch]++
	}
	return frequencies
}

func buildTrie(frequencies map[byte]int) *node {
	nodes := make([]*node, len(frequencies), len(frequencies))
	i := 0
	for ch, freq := range frequencies {
		nodes[i] = &node{ch: ch, freq: freq}
		i++
	}

	queue := nodesPriorityQueue(nodes)
	heap.Init(&queue)
	for queue.Len() > 1 {
		left := heap.Pop(&queue).(*node)
		right := heap.Pop(&queue).(*node)
		node := node{freq: left.freq + right.freq, left: left, right: right}
		heap.Push(&queue, &node)
	}
	return heap.Pop(&queue).(*node)
}

func buildTableCore(node *node, table byteCodeTable, value string) {
	if node.isLeaf() {
		table.set(node.ch, value)
	} else {
		buildTableCore(node.left, table, value+"0")
		buildTableCore(node.right, table, value+"1")
	}
}

func buildTable(root *node) byteCodeTable {
	table := newByteCodeTable()
	buildTableCore(root, table, "")
	return table
}

func saveTrieCore(node *node, builder *strings.Builder) {
	if node.isLeaf() {
		builder.WriteString("1")
		builder.WriteByte(node.ch)
	} else {
		builder.WriteString("0")
		saveTrieCore(node.left, builder)
		saveTrieCore(node.right, builder)
	}
}

func saveTrie(root *node) string {
	builder := new(strings.Builder)
	saveTrieCore(root, builder)
	return builder.String()
}

func compressCore(data []byte, table byteCodeTable) string {
	result := ""
	for _, ch := range data {
		coded := table.get(ch)
		result += coded
	}
	return result
}

// Compress compresses *data*.
// https://algs4.cs.princeton.edu/55compression/Huffman.java.html
func Compress(data []byte) string {
	frequencies := collectFrequencies(data)
	root := buildTrie(frequencies)
	table := buildTable(root)
	compressedTrie := saveTrie(root)
	len := fmt.Sprintf("%08d", len(data))
	compressedData := compressCore(data, table)
	return compressedTrie + len + compressedData
}

func loadTrie(data string, idx int) (*node, int) {
	if data[idx] == byte('1') {
		ch := data[idx+1]
		return &node{ch: ch}, 2
	}
	node := &node{}
	left, leftCnt := loadTrie(data, idx+1)
	right, rightCng := loadTrie(data, idx+1+leftCnt)
	node.left = left
	node.right = right
	return node, 1 + leftCnt + rightCng
}

func decompressCore(data string, len int, root *node) []byte {
	ret := make([]byte, len, len)
	i := 0
	idx := 0
	for i < len {
		node := root
		for !node.isLeaf() {
			ch := data[idx]
			idx++
			if ch == byte('1') {
				node = node.right
			} else {
				node = node.left
			}
		}
		ret[i] = node.ch
		i++
	}
	return ret
}

// Decompress decompressed *data*.
func Decompress(data string) []byte {
	root, cnt := loadTrie(data, 0)
	if root == nil {
		return nil
	}
	var len int
	fmt.Sscanf(data[cnt:cnt+8], "%d", &len)
	ret := decompressCore(data[cnt+8:], len, root)
	return ret
}
