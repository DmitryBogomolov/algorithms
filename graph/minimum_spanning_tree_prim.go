package graph

import "container/heap"

type verticesPQ struct {
	size          int
	vertexToIndex []int
	indexToVertex []int
	weights       []float64
}

func (pq *verticesPQ) Len() int {
	return pq.size
}
func (pq *verticesPQ) Less(i, j int) bool {
	vi, vj := pq.indexToVertex[i], pq.indexToVertex[j]
	return pq.weights[vi] < pq.weights[vj]
}
func (pq *verticesPQ) Swap(i, j int) {
	vi, vj := pq.indexToVertex[j], pq.indexToVertex[i]
	pq.indexToVertex[i], pq.indexToVertex[j] = vi, vj
	pq.vertexToIndex[vi], pq.vertexToIndex[vj] = i, j
}

func (pq *verticesPQ) Push(val interface{}) {
}
func (pq *verticesPQ) Pop() interface{} {
	return nil
}

func (pq *verticesPQ) update(vertex int, weight float64) {
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
func (pq *verticesPQ) pop() int {
	ret := pq.indexToVertex[0]
	heap.Pop(pq)
	pq.indexToVertex[pq.size] = -1
	pq.vertexToIndex[ret] = -1
	pq.size--
	return ret
}

func newVerticesPQ(numVertices int) *verticesPQ {
	pq := verticesPQ{
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
