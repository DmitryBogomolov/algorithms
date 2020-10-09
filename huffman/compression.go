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
	nodes := make([]*node, len(frequencies))
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
	block.align()
	return block
}

func compressLength(length int) bitBlock {
	var block bitBlock
	block.appendByte(byte(length))
	block.appendByte(byte(length >> 8))
	block.appendByte(byte(length >> 16))
	block.appendByte(byte(length >> 24))
	return block
}

func compressData(data []byte, table byteCodeTable) bitBlock {
	var block bitBlock
	for _, ch := range data {
		code := table.get(ch)
		block.append(code)
	}
	block.align()
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
	block.append(lengthBlock)
	block.append(dataBlock)
	return block.buffer
}

func expandTrieCore(scanner *bitScanner) *node {
	var n node
	if scanner.readBit() {
		n.ch = scanner.readByte()
	} else {
		n.left = expandTrieCore(scanner)
		n.right = expandTrieCore(scanner)
	}
	return &n
}

func expandTrie(scanner *bitScanner) *node {
	root := expandTrieCore(scanner)
	scanner.align()
	return root
}

func expandLength(scanner *bitScanner) int {
	var length int
	length |= int(scanner.readByte())
	length |= int(scanner.readByte()) << 8
	length |= int(scanner.readByte()) << 16
	length |= int(scanner.readByte()) << 24
	return length
}

func expandData(scanner *bitScanner, length int, root *node) []byte {
	buffer := make([]byte, length)
	for i := 0; i < length; i++ {
		node := root
		for !node.isLeaf() {
			if scanner.readBit() {
				node = node.right
			} else {
				node = node.left
			}
		}
		buffer[i] = node.ch
	}
	scanner.align()
	return buffer
}

// Expand expands *data*.
func Expand(data []byte) []byte {
	scanner := newBitScanner(data)
	root := expandTrie(scanner)
	length := expandLength(scanner)
	buffer := expandData(scanner, length, root)
	return buffer
}
