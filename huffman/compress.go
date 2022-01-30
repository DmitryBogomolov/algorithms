package huffman

import (
	"bytes"

	"github.com/DmitryBogomolov/algorithms/bits"
	"github.com/DmitryBogomolov/algorithms/priorityqueue"
)

type _ByteCode []byte
type _ByteCodeTable []_ByteCode

func buildTree(data []byte) *_Node {
	frequencies := make([]int, 256)
	for _, item := range data {
		frequencies[item]++
	}
	queue := priorityqueue.New(func(lhs, rhs interface{}) bool {
		lNode := lhs.(*_Node)
		rNode := rhs.(*_Node)
		return lNode.frequency < rNode.frequency
	})
	for i, freq := range frequencies {
		if freq > 0 {
			queue.Insert(&_Node{item: byte(i), frequency: freq})
		}
	}
	for queue.Size() > 1 {
		lNode := queue.Remove().(*_Node)
		rNode := queue.Remove().(*_Node)
		n := &_Node{
			frequency: lNode.frequency + rNode.frequency,
			lNode:     lNode,
			rNode:     rNode,
		}
		queue.Insert(n)
	}
	return queue.Remove().(*_Node)
}

func extendCode(code _ByteCode, bit byte) _ByteCode {
	tmp := append(_ByteCode{}, code...)
	return append(tmp, bit)
}

func buildTableCore(node *_Node, table _ByteCodeTable, code []byte) {
	if node.isLeaf() {
		table[node.item] = code
	} else {
		buildTableCore(node.lNode, table, extendCode(code, 0))
		buildTableCore(node.rNode, table, extendCode(code, 1))
	}
}

func buildTable(root *_Node) _ByteCodeTable {
	table := make(_ByteCodeTable, 256)
	buildTableCore(root, table, nil)
	return table
}

func compressTreeCore(node *_Node, writer *bits.BitWriter) error {
	var err error
	if node.isLeaf() {
		err = writer.WriteBit(1)
		if err == nil {
			writer.WriteUint8(node.item)
		}
	} else {
		err = writer.WriteBit(0)
		if err == nil {
			compressTreeCore(node.lNode, writer)
		}
		if err == nil {
			compressTreeCore(node.rNode, writer)
		}
	}
	return err
}

func compressTree(root *_Node, writer *bits.BitWriter) error {
	err := compressTreeCore(root, writer)
	if err == nil {
		err = writer.Flush()
	}
	return err
}

func compressByteCode(code _ByteCode, writer *bits.BitWriter) error {
	var err error
	for _, bit := range code {
		if err = writer.WriteBit(bit); err != nil {
			break
		}
	}
	return err
}

func compressData(data []byte, table _ByteCodeTable, writer *bits.BitWriter) error {
	var err error
	for _, item := range data {
		code := table[item]
		if err = compressByteCode(code, writer); err != nil {
			break
		}
	}
	if err == nil {
		err = writer.Flush()
	}
	return err
}

// Compress compresses *data*.
// https://algs4.cs.princeton.edu/55compression/Huffman.java.html
func Compress(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, ErrEmptyData
	}
	root := buildTree(data)
	table := buildTable(root)
	var buffer bytes.Buffer
	writer := bits.NewBitWriter(&buffer)
	var err error
	if err = compressTree(root, writer); err != nil {
		return nil, CompressError{err}
	}
	if err = writer.WriteUint32(uint32(len(data))); err != nil {
		return nil, CompressError{err}
	}
	if err = compressData(data, table, writer); err != nil {
		return nil, CompressError{err}
	}
	return buffer.Bytes(), nil
}
