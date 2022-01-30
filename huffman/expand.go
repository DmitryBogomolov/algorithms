package huffman

import (
	"bufio"
	"io"

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
	if err == nil {
		_, err = reader.Flush()
	}
	return root, err
}

func expandByteCode(root *_Node, reader *bits.BitReader) (byte, error) {
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

func expandData(root *_Node, length int, reader *bits.BitReader, target io.ByteWriter) error {
	var err error
	for i := 0; i < length; i++ {
		b, err := expandByteCode(root, reader)
		if err != nil {
			return err
		}
		if err = target.WriteByte(b); err != nil {
			return err
		}
	}
	_, err = reader.Flush()
	return err
}

// Expand expands *data*.
// https://algs4.cs.princeton.edu/55compression/Huffman.java.html
func Expand(input io.Reader, output io.Writer) error {
	var err error
	inputWrapper := bufio.NewReader(input)
	bitReader := bits.NewBitReader(inputWrapper)
	root, err := expandTree(bitReader)
	if err != nil {
		return ExpandError{err}
	}
	length, err := bitReader.ReadUint32()
	if err != nil {
		return ExpandError{err}
	}
	outputWrapper := bufio.NewWriter(output)
	if err = expandData(root, int(length), bitReader, outputWrapper); err != nil {
		return ExpandError{err}
	}
	if err = outputWrapper.Flush(); err != nil {
		return ExpandError{err}
	}
	return nil
}
