package huffman

func collectFrequencies(data []byte) map[byte]int {
	frequencies := make(map[byte]int)
	for _, item := range data {
		frequencies[item]++
	}
	return frequencies
}

func buildTableCore(node *node, table byteCodeTable, code *bitBlock, trieBits *int, dataBits *int) {
	*trieBits++
	if node.isLeaf() {
		*dataBits += node.frequency * code.size
		*trieBits += 8
		table.set(node.item, code)
	} else {
		lCode := code.clone()
		lCode.appendBit(false)
		rCode := code.clone()
		rCode.appendBit(true)
		buildTableCore(node.lNode, table, lCode, trieBits, dataBits)
		buildTableCore(node.rNode, table, rCode, trieBits, dataBits)
	}
}

func buildTable(root *node) (byteCodeTable, int) {
	table := newByteCodeTable()
	trieBits, dataBits := 0, 0
	buildTableCore(root, table, newBitBlock(0), &trieBits, &dataBits)
	// 32 - is for length, 14 - is a worst case of two alignments.
	return table, trieBits + 32 + 14 + dataBits
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
	table, blockSize := buildTable(root)
	block := newBitBlock(blockSize)
	compressTrie(root, block)
	compressLength(len(data), block)
	compressData(data, table, block)
	return block.buffer
}
