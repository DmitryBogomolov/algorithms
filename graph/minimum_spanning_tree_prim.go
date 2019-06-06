package graph

import (
	"container/heap"
	"math"
)

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

type minimumSpanningTree struct {
	origin    EdgeWeightedGraph
	numEdges  int
	adjacency [][]int
	weights   [][]float64
}

func (t minimumSpanningTree) NumVertices() int {
	return t.origin.NumVertices()
}
func (t minimumSpanningTree) NumEdges() int {
	return t.numEdges
}
func (t minimumSpanningTree) AdjacentVertices(vertex int) []int {
	return t.adjacency[vertex]
}
func (t minimumSpanningTree) AdjacentWeights(vertex int) []float64 {
	return t.weights[vertex]
}

func scanMinimumSpanningTreeVertex(
	pq *verticesPQ, marked []bool, edgeTo []int, distTo []float64,
	graph EdgeWeightedGraph, current int,
) {
	marked[current] = true
	weights := graph.AdjacentWeights(current)
	for i, v := range graph.AdjacentVertices(current) {
		weight := weights[i]
		if !marked[v] && weight < distTo[v] {
			edgeTo[v] = current
			distTo[v] = weight
			pq.update(v, weight)
		}
	}
}

func processMinimumSpanningTree(
	pq *verticesPQ, marked []bool, edgeTo []int, distTo []float64,
	graph EdgeWeightedGraph, start int,
) {
	distTo[start] = 0
	pq.update(start, 0)
	for pq.Len() > 0 {
		v := pq.pop()
		scanMinimumSpanningTreeVertex(pq, marked, edgeTo, distTo, graph, v)
	}
}

// MinimumSpanningTreePrim computes minimum spanning tree for a graph using Prim's algorithm.
func MinimumSpanningTreePrim(graph EdgeWeightedGraph) EdgeWeightedGraph {
	numVertices := graph.NumVertices()
	marked := make([]bool, numVertices)
	edgeTo := make([]int, numVertices)
	distTo := make([]float64, numVertices)
	pq := newVerticesPQ(numVertices)
	for v := 0; v < numVertices; v++ {
		edgeTo[v] = -1
		distTo[v] = math.MaxFloat64
	}
	for v := 0; v < numVertices; v++ {
		if !marked[v] {
			processMinimumSpanningTree(pq, marked, edgeTo, distTo, graph, v)
		}
	}
	adjacency := make([][]int, numVertices)
	weights := make([][]float64, numVertices)
	numEdges := 0
	for v := 0; v < numVertices; v++ {
		w := edgeTo[v]
		if w != -1 {
			weight := distTo[v]
			adjacency[v] = append(adjacency[v], w)
			adjacency[w] = append(adjacency[w], v)
			weights[v] = append(weights[v], weight)
			weights[w] = append(weights[w], weight)
			numEdges++
		}
	}
	return minimumSpanningTree{
		origin:    graph,
		numEdges:  numEdges,
		adjacency: adjacency,
		weights:   weights,
	}
}
