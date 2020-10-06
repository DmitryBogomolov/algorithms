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

func buildTableCore(node *node, table map[byte]string, value string) {
	if node.isLeaf() {
		table[node.ch] = value
	} else {
		buildTableCore(node.left, table, value+"0")
		buildTableCore(node.right, table, value+"1")
	}
}

func buildTable(root *node) map[byte]string {
	table := make(map[byte]string)
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

func compressCore(data []byte, table map[byte]string) string {
	result := ""
	for _, ch := range data {
		coded := table[ch]
		result += coded
	}
	return result
}

// Compress compresses *data*.
// https://algs4.cs.princeton.edu/55compression/Huffman.java.html
func Compress(data []byte) {
	frequencies := collectFrequencies(data)
	root := buildTrie(frequencies)
	table := buildTable(root)
	compressedTrie := saveTrie(root)
	compressedData := compressCore(data, table)
	fmt.Printf("trie: %s\n", compressedTrie)
	fmt.Printf("compressed: %d, original: %d\n", len(compressedData), len(data)*8)
}
