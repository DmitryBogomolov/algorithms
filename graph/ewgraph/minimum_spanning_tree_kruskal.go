package ewgraph

import (
	"algorithms/graph/graph"
	"algorithms/unionfind"
)

// MinimumSpanningTreeKruskal computes minimum spanning tree using Kruskal's algorithm.
// https://algs4.cs.princeton.edu/43mst/KruskalMST.java.html
func MinimumSpanningTreeKruskal(ewgraph EdgeWeightedGraph) EdgeWeightedGraph {
	edgesPriorityQueue := newEdgesPriorityQueue()
	allWeights := AllGraphWeights(ewgraph)
	for i, edge := range graph.AllGraphEdges(ewgraph) {
		edgesPriorityQueue.pushEdge(edge, allWeights[i])
	}
	numVertices := ewgraph.NumVertices()
	uf := unionfind.New(numVertices)
	adjacency := make([][]int, numVertices)
	weights := make([][]float64, numVertices)
	numEdges := 0
	for edgesPriorityQueue.Len() > 0 {
		edge, weight := edgesPriorityQueue.popEdge()
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
