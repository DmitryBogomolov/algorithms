package huffman

import "container/heap"

type node struct {
	item         byte
	frequency    int
	lNode, rNode *node
}

func (n node) isLeaf() bool {
	return n.lNode == nil && n.rNode == nil
}

type nodesPriorityQueue []*node

func (pq nodesPriorityQueue) Len() int {
	return len(pq)
}

func (pq nodesPriorityQueue) Less(i, j int) bool {
	return pq[i].frequency < pq[j].frequency
}

func (pq nodesPriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *nodesPriorityQueue) Push(item interface{}) {
	nodes := *pq
	node := item.(*node)
	*pq = append(nodes, node)
}

func (pq *nodesPriorityQueue) Pop() interface{} {
	nodes := *pq
	len := len(nodes)
	node := nodes[len-1]
	nodes[len-1] = nil
	*pq = nodes[0 : len-1]
	return node
}

func buildTrie(frequencies map[byte]int) *node {
	nodes := make([]*node, len(frequencies))
	i := 0
	for item, frequency := range frequencies {
		nodes[i] = &node{item: item, frequency: frequency}
		i++
	}

	queue := nodesPriorityQueue(nodes)
	heap.Init(&queue)
	for queue.Len() > 1 {
		lNode := heap.Pop(&queue).(*node)
		rNode := heap.Pop(&queue).(*node)
		node := node{
			frequency: lNode.frequency + rNode.frequency,
			lNode:     lNode,
			rNode:     rNode,
		}
		heap.Push(&queue, &node)
	}
	return heap.Pop(&queue).(*node)
}
