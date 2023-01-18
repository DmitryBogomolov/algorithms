package huffman

import (
	"bufio"
	"io"
	"io/ioutil"

	"github.com/DmitryBogomolov/algorithms/bits"
	"github.com/DmitryBogomolov/algorithms/priorityqueue"
)

type _ByteCode []byte
type _ByteCodeTable []_ByteCode

const tableSize = 256

func buildTree(data []byte) *_Node {
	frequencies := make([]int, tableSize)
	for _, item := range data {
		frequencies[item]++
	}
	queue := priorityqueue.New(func(lhs, rhs *_Node) bool {
		return lhs.frequency < rhs.frequency
	})
	for i, freq := range frequencies {
		if freq > 0 {
			queue.Insert(&_Node{item: byte(i), frequency: freq})
		}
	}
	for queue.Size() > 1 {
		lNode := queue.Remove()
		rNode := queue.Remove()
		n := &_Node{
			frequency: lNode.frequency + rNode.frequency,
			lNode:     lNode,
			rNode:     rNode,
		}
		queue.Insert(n)
	}
	return queue.Remove()
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
	table := make(_ByteCodeTable, tableSize)
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
func Compress(input io.Reader, output io.Writer) error {
	var err error
	data, err := ioutil.ReadAll(input)
	if err != nil {
		return CompressError{err}
	}
	if len(data) == 0 {
		return CompressError{io.EOF}
	}
	root := buildTree(data)
	table := buildTable(root)
	outputWrapper := bufio.NewWriter(output)
	bitWriter := bits.NewBitWriter(outputWrapper)
	if err = compressTree(root, bitWriter); err != nil {
		return CompressError{err}
	}
	if err = bitWriter.WriteUint32(uint32(len(data))); err != nil {
		return CompressError{err}
	}
	if err = compressData(data, table, bitWriter); err != nil {
		return CompressError{err}
	}
	if err = outputWrapper.Flush(); err != nil {
		return CompressError{err}
	}
	return nil
}
