package ewgraph

import (
	"algorithms/graph/graph"
	"container/heap"
)

type _EdgesPriorityQueue struct {
	edges   []graph.Edge
	weights []float64
}

func (pq *_EdgesPriorityQueue) Len() int {
	return len(pq.edges)
}
func (pq *_EdgesPriorityQueue) Less(i, j int) bool {
	return pq.weights[i] < pq.weights[j]
}
func (pq *_EdgesPriorityQueue) Swap(i, j int) {
	pq.edges[i], pq.edges[j] = pq.edges[j], pq.edges[i]
	pq.weights[i], pq.weights[j] = pq.weights[j], pq.weights[i]
}

func (pq *_EdgesPriorityQueue) Push(val interface{}) {
	// pushEdge is used instead.
}
func (pq *_EdgesPriorityQueue) Pop() interface{} {
	// popEdge is used instead.
	return nil
}

func (pq *_EdgesPriorityQueue) pushEdge(edge graph.Edge, weight float64) {
	pq.edges = append(pq.edges, edge)
	pq.weights = append(pq.weights, weight)
	heap.Push(pq, nil)
}
func (pq *_EdgesPriorityQueue) popEdge() (graph.Edge, float64) {
	edge, weight := pq.edges[0], pq.weights[0]
	heap.Pop(pq)
	n := pq.Len() - 1
	pq.edges = pq.edges[:n]
	pq.weights = pq.weights[:n]
	return edge, weight
}

func newEdgesPriorityQueue() *_EdgesPriorityQueue {
	return &_EdgesPriorityQueue{}
}
