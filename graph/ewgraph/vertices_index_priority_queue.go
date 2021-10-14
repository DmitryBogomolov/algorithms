package ewgraph

import "container/heap"

type _VerticesIndexPriorityQueue struct {
	size          int
	vertexToIndex []int
	indexToVertex []int
	weights       []float64
}

func (pq *_VerticesIndexPriorityQueue) Len() int {
	return pq.size
}
func (pq *_VerticesIndexPriorityQueue) Less(i, j int) bool {
	vi, vj := pq.indexToVertex[i], pq.indexToVertex[j]
	return pq.weights[vi] < pq.weights[vj]
}
func (pq *_VerticesIndexPriorityQueue) Swap(i, j int) {
	vi, vj := pq.indexToVertex[j], pq.indexToVertex[i]
	pq.indexToVertex[i], pq.indexToVertex[j] = vi, vj
	pq.vertexToIndex[vi], pq.vertexToIndex[vj] = i, j
}

func (pq *_VerticesIndexPriorityQueue) Push(val interface{}) {
	// updateVertex is used instead.
}
func (pq *_VerticesIndexPriorityQueue) Pop() interface{} {
	// popVertex is used instead.
	return nil
}

func (pq *_VerticesIndexPriorityQueue) updateVertex(vertex int, weight float64) {
	index := pq.vertexToIndex[vertex]
	if index == -1 {
		pq.indexToVertex[pq.size] = vertex
		pq.vertexToIndex[vertex] = pq.size
		pq.size++
		pq.weights[vertex] = weight
		heap.Push(pq, nil)
	} else {
		pq.weights[vertex] = weight
		heap.Fix(pq, index)
	}

}
func (pq *_VerticesIndexPriorityQueue) popVertex() int {
	ret := pq.indexToVertex[0]
	heap.Pop(pq)
	pq.indexToVertex[pq.size] = -1
	pq.vertexToIndex[ret] = -1
	pq.size--
	return ret
}

func newVerticesIndexPriorityQueue(numVertices int) *_VerticesIndexPriorityQueue {
	pq := _VerticesIndexPriorityQueue{
		vertexToIndex: make([]int, numVertices),
		indexToVertex: make([]int, numVertices),
		weights:       make([]float64, numVertices),
	}
	for i := 0; i < numVertices; i++ {
		pq.vertexToIndex[i] = -1
		pq.indexToVertex[i] = -1
	}
	return &pq
}
