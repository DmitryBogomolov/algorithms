package huffman

func collectFrequencies(data []byte) map[byte]int {
	frequencies := make(map[byte]int)
	for _, ch := range data {
		frequencies[ch]++
	}
	return frequencies
}

func buildTableCore(node *node, table byteCodeTable, code bitBlock) {
	if node.isLeaf() {
		table.set(node.ch, code)
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
	buildTableCore(root, table, bitBlock{})
	return table
}

func compressTrieCore(node *node, block *bitBlock) {
	if node.isLeaf() {
		block.appendBit(true)
		block.appendByte(node.ch)
	} else {
		block.appendBit(false)
		compressTrieCore(node.lNode, block)
		compressTrieCore(node.rNode, block)
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
