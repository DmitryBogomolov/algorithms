package ewgraph

import (
	"algorithms/graph/graph"
	"algorithms/unionfind"
	"container/heap"
)

type edgesPQ struct {
	edges   []graph.Edge
	weights []float64
}

func (pq *edgesPQ) Len() int {
	return len(pq.edges)
}
func (pq *edgesPQ) Less(i, j int) bool {
	return pq.weights[i] < pq.weights[j]
}
func (pq *edgesPQ) Swap(i, j int) {
	pq.edges[i], pq.edges[j] = pq.edges[j], pq.edges[i]
	pq.weights[i], pq.weights[j] = pq.weights[j], pq.weights[i]
}

func (pq *edgesPQ) Push(val interface{}) {
}
func (pq *edgesPQ) Pop() interface{} {
	return nil
}

func (pq *edgesPQ) push(edge graph.Edge, weight float64) {
	pq.edges = append(pq.edges, edge)
	pq.weights = append(pq.weights, weight)
	heap.Push(pq, nil)
}
func (pq *edgesPQ) pop() (graph.Edge, float64) {
	edge, weight := pq.edges[0], pq.weights[0]
	heap.Pop(pq)
	n := pq.Len() - 1
	pq.edges = pq.edges[:n]
	pq.weights = pq.weights[:n]
	return edge, weight
}

func newEdgesPQ() *edgesPQ {
	return &edgesPQ{}
}

// MinimumSpanningTreeKruskal computes minimum spanning tree using Kruskal's algorithm.
// https://algs4.cs.princeton.edu/43mst/KruskalMST.java.html
func MinimumSpanningTreeKruskal(ewgraph EdgeWeightedGraph) EdgeWeightedGraph {
	pq := newEdgesPQ()
	allWeights := AllGraphWeights(ewgraph)
	for i, edge := range graph.AllGraphEdges(ewgraph) {
		pq.push(edge, allWeights[i])
	}
	numVertices := ewgraph.NumVertices()
	uf := unionfind.New(numVertices)
	adjacency := make([][]int, numVertices)
	weights := make([][]float64, numVertices)
	numEdges := 0
	for pq.Len() > 0 {
		edge, weight := pq.pop()
		vertexID1, vertexID2 := edge.Vertex1(), edge.Vertex2()
		if !uf.Connected(vertexID1, vertexID2) {
			uf.Union(vertexID1, vertexID2)
			addWeightedEdge(adjacency, weights, vertexID1, vertexID2, weight)
			numEdges++
		}
	}
	return minimumSpanningTree{
		origin:    ewgraph,
		numEdges:  numEdges,
		adjacency: adjacency,
		weights:   weights,
	}
}
