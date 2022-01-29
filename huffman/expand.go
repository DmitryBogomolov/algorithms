package huffman

import (
	"bytes"

	"github.com/DmitryBogomolov/algorithms/bits"
)

func expandTreeCore(reader *bits.BitReader) (*_Node, error) {
	var node _Node
	bit, err := reader.ReadBit()
	if err != nil {
		return nil, err
	}
	if bit > 0 {
		if node.item, err = reader.ReadUint8(); err != nil {
			return nil, err
		}
	} else {
		if node.lNode, err = expandTreeCore(reader); err != nil {
			return nil, err
		}
		if node.rNode, err = expandTreeCore(reader); err != nil {
			return nil, err
		}
	}
	return &node, nil
}

func expandTree(reader *bits.BitReader) (*_Node, error) {
	root, err := expandTreeCore(reader)
	reader.Flush()
	return root, err
}

func expandByteCode(reader *bits.BitReader, root *_Node) (byte, error) {
	node := root
	for !node.isLeaf() {
		bit, err := reader.ReadBit()
		if err != nil {
			return 0, err
		}
		if bit > 0 {
			node = node.rNode
		} else {
			node = node.lNode
		}
	}
	return node.item, nil
}

func expandData(reader *bits.BitReader, length int, root *_Node) ([]byte, error) {
	buffer := make([]byte, length)
	for i := 0; i < length; i++ {
		b, err := expandByteCode(reader, root)
		if err != nil {
			return nil, err
		}
		buffer[i] = b
	}
	reader.Flush()
	return buffer, nil
}

// Expand expands *data*.
// https://algs4.cs.princeton.edu/55compression/Huffman.java.html
func Expand(data []byte) (result []byte, err error) {
	if len(data) == 0 {
		return nil, ErrEmptyData
	}
	buffer := bytes.NewBuffer(data)
	reader := bits.NewBitReader(buffer)
	root, err := expandTree(reader)
	if err != nil {
		return nil, ErrDataCorrupted
	}
	length, err := reader.ReadUint32()
	if err != nil {
		return nil, ErrDataCorrupted
	}
	result, err = expandData(reader, int(length), root)
	if err != nil {
		return nil, ErrDataCorrupted
	}
	return result, nil
}
