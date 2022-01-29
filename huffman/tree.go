package huffman

import (
	"github.com/DmitryBogomolov/algorithms/priorityqueue"
)

type _Node struct {
	item         byte
	frequency    int
	lNode, rNode *_Node
}

func (n _Node) isLeaf() bool {
	return n.lNode == nil && n.rNode == nil
}

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
