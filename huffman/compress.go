package huffman

func collectFrequencies(data []byte) map[byte]int {
	frequencies := make(map[byte]int)
	for _, item := range data {
		frequencies[item]++
	}
	return frequencies
}

func buildTableCore(node *node, table byteCodeTable, code *bitBlock) {
	if node.isLeaf() {
		table.set(node.item, code)
	} else {
		lCode := code.clone()
		lCode.appendBit(false)
		rCode := code.clone()
		rCode.appendBit(true)
		buildTableCore(node.lNode, table, lCode)
		buildTableCore(node.rNode, table, rCode)
	}
}

func buildTable(root *node) byteCodeTable {
	table := newByteCodeTable()
	buildTableCore(root, table, newBitBlock())
	return table
}

func compressTrieCore(node *node, block *bitBlock) {
	if node.isLeaf() {
		block.appendBit(true)
		block.appendByte(node.item)
	} else {
		block.appendBit(false)
		compressTrieCore(node.lNode, block)
		compressTrieCore(node.rNode, block)
	}
}

func compressTrie(root *node, block *bitBlock) {
	compressTrieCore(root, block)
	block.align()
}

func compressLength(length int, block *bitBlock) {
	block.appendByte(byte(length))
	block.appendByte(byte(length >> 8))
	block.appendByte(byte(length >> 16))
	block.appendByte(byte(length >> 24))
}

func compressData(data []byte, table byteCodeTable, block *bitBlock) {
	for _, item := range data {
		code := table.get(item)
		block.append(code)
	}
	block.align()
}

// Compress compresses *data*.
// https://algs4.cs.princeton.edu/55compression/Huffman.java.html
func Compress(data []byte) []byte {
	frequencies := collectFrequencies(data)
	root := buildTrie(frequencies)
	table := buildTable(root)
	block := newBitBlock()
	compressTrie(root, block)
	compressLength(len(data), block)
	compressData(data, table, block)
	return block.buffer
}
