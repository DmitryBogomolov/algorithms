package huffman

import (
	"encoding/binary"
	"errors"
)

func collectFrequencies(data []byte) map[byte]int {
	frequencies := make(map[byte]int)
	for _, item := range data {
		frequencies[item]++
	}
	return frequencies
}

func buildTableCore(node *_Node, table byteCodeTable, code *bitBlock, treeBits *int, dataBits *int) {
	*treeBits++
	if node.isLeaf() {
		*dataBits += node.frequency * code.size
		*treeBits += 8
		table.set(node.item, code)
	} else {
		lCode := code.clone()
		lCode.appendBit(false)
		rCode := code.clone()
		rCode.appendBit(true)
		buildTableCore(node.lNode, table, lCode, treeBits, dataBits)
		buildTableCore(node.rNode, table, rCode, treeBits, dataBits)
	}
}

func buildTable(root *_Node) (byteCodeTable, int) {
	table := newByteCodeTable()
	treeBits, dataBits := 0, 0
	buildTableCore(root, table, newBitBlock(0), &treeBits, &dataBits)
	// 32 - is for length, 14 - is a worst case of two alignments.
	return table, treeBits + 32 + 14 + dataBits
}

func compressTreeCore(node *_Node, block *bitBlock) {
	if node.isLeaf() {
		block.appendBit(true)
		block.appendByte(node.item)
	} else {
		block.appendBit(false)
		compressTreeCore(node.lNode, block)
		compressTreeCore(node.rNode, block)
	}
}

func compressTree(root *_Node, block *bitBlock) {
	compressTreeCore(root, block)
	block.align()
}

func compressLength(length int, block *bitBlock) {
	var bytes [4]byte
	binary.LittleEndian.PutUint32(bytes[:], uint32(length))
	block.appendByte(bytes[0])
	block.appendByte(bytes[1])
	block.appendByte(bytes[2])
	block.appendByte(bytes[3])
}

func compressData(data []byte, table byteCodeTable, block *bitBlock) {
	for _, item := range data {
		code := table.get(item)
		block.append(code)
	}
	block.align()
}

// ErrEmptyData tells that data is nil or empty.
var ErrEmptyData = errors.New("data is nil or empty")

// Compress compresses *data*.
// https://algs4.cs.princeton.edu/55compression/Huffman.java.html
func Compress(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, ErrEmptyData
	}
	frequencies := collectFrequencies(data)
	root := buildTree(frequencies)
	table, blockSize := buildTable(root)
	block := newBitBlock(blockSize)
	compressTree(root, block)
	compressLength(len(data), block)
	compressData(data, table, block)
	return block.getBuffer(), nil
}
