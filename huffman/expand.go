package huffman

import (
	"encoding/binary"
)

func expandTreeCore(scanner *bitScanner) *_Node {
	var n _Node
	if scanner.readBit() {
		n.item = scanner.readByte()
	} else {
		n.lNode = expandTreeCore(scanner)
		n.rNode = expandTreeCore(scanner)
	}
	return &n
}

func expandTree(scanner *bitScanner) *_Node {
	root := expandTreeCore(scanner)
	scanner.align()
	return root
}

func expandLength(scanner *bitScanner) int {
	return int(binary.LittleEndian.Uint32([]byte{
		scanner.readByte(),
		scanner.readByte(),
		scanner.readByte(),
		scanner.readByte(),
	}))
}

func expandData(scanner *bitScanner, length int, root *_Node) []byte {
	buffer := make([]byte, length)
	for i := 0; i < length; i++ {
		node := root
		for !node.isLeaf() {
			if scanner.readBit() {
				node = node.rNode
			} else {
				node = node.lNode
			}
		}
		buffer[i] = node.item
	}
	scanner.align()
	return buffer
}

// Expand expands *data*.
// https://algs4.cs.princeton.edu/55compression/Huffman.java.html
func Expand(data []byte) (buffer []byte, err error) {
	if len(data) == 0 {
		err = ErrEmptyData
		return
	}
	defer func() {
		if innerErr := recover(); innerErr != nil {
			err = ErrDataCorrupted
		}
	}()
	scanner := newBitScanner(data)
	root := expandTree(scanner)
	length := expandLength(scanner)
	buffer = expandData(scanner, length, root)
	return
}
