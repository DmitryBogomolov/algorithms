package huffman

import "errors"

func expandTrieCore(scanner *bitScanner) *node {
	var n node
	if scanner.readBit() {
		n.item = scanner.readByte()
	} else {
		n.lNode = expandTrieCore(scanner)
		n.rNode = expandTrieCore(scanner)
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

// ErrDataCorrupted tells that data is corrupted.
var ErrDataCorrupted = errors.New("data is corrupted")

// Expand expands *data*.
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
	root := expandTrie(scanner)
	length := expandLength(scanner)
	buffer = expandData(scanner, length, root)
	return
}
