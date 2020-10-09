package huffman

import (
	"container/heap"
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

func buildTableCore(node *node, table byteCodeTable, code bitBlock) {
	if node.isLeaf() {
		table.set(node.ch, code)
	} else {
		leftCode := code.clone()
		leftCode.appendBit(false)
		rightCode := code.clone()
		rightCode.appendBit(true)
		buildTableCore(node.left, table, leftCode)
		buildTableCore(node.right, table, rightCode)
	}
}

func buildTable(root *node) byteCodeTable {
	table := newByteCodeTable()
	buildTableCore(root, table, bitBlock{})
	return table
}

func compressTrieCore(node *node, block *bitBlock) {
	if node.isLeaf() {
		block.appendBit(true)
		block.appendByte(node.ch)
	} else {
		block.appendBit(false)
		compressTrieCore(node.left, block)
		compressTrieCore(node.right, block)
	}
}

func compressTrie(root *node) bitBlock {
	var block bitBlock
	compressTrieCore(root, &block)
	return block
}

func compressLength(length int) bitBlock {
	var block bitBlock
	// TODO: Use 4 bytes.
	block.appendByte(byte(length))
	return block
}

func compressData(data []byte, table byteCodeTable) bitBlock {
	var block bitBlock
	for _, ch := range data {
		code := table.get(ch)
		block.append(code)
	}
	return block
}

// Compress compresses *data*.
// https://algs4.cs.princeton.edu/55compression/Huffman.java.html
func Compress(data []byte) []byte {
	frequencies := collectFrequencies(data)
	root := buildTrie(frequencies)
	table := buildTable(root)
	trieBlock := compressTrie(root)
	lengthBlock := compressLength(len(data))
	dataBlock := compressData(data, table)
	var block bitBlock
	block.append(trieBlock)
	block.align()
	block.append(lengthBlock)
	block.align()
	block.append(dataBlock)
	return block.buffer
}

func expandTrie(scanner *bitScanner) *node {
	bit := scanner.readBit()
	if bit {
		ch := scanner.readByte()
		return &node{ch: ch}
	}
	node := &node{}
	left := expandTrie(scanner)
	right := expandTrie(scanner)
	node.left = left
	node.right = right
	return node
}

func expandLength(scanner *bitScanner) int {
	// TODO: Use 4 bytes.
	length := scanner.readByte()
	return int(length)
}

func expandData(scanner *bitScanner, length int, root *node) []byte {
	buffer := make([]byte, length)
	i := 0
	idx := 0
	for i < length {
		node := root
		for !node.isLeaf() {
			bit := scanner.readBit()
			idx++
			if bit {
				node = node.right
			} else {
				node = node.left
			}
		}
		buffer[i] = node.ch
		i++
	}
	return buffer
}

// Expand expands *data*.
func Expand(data []byte) []byte {
	scanner := newBitScanner(data)
	root := expandTrie(scanner)
	scanner.align()
	length := expandLength(scanner)
	scanner.align()
	buffer := expandData(scanner, length, root)
	return buffer
}
