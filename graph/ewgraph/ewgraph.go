package ewgraph

import (
	"algorithms/graph/graph"
	"algorithms/graph/internals"
)

// EdgeWeightedGraph is a graph where each edge has an associated weight.
type EdgeWeightedGraph interface {
	graph.Graph
	// AdjacentWeights returns weights of edges adjacent to the vertex.
	AdjacentWeights(vertexID int) []float64
}

// AllGraphWeights returns all edge weights of an edge-weighted graph.
func AllGraphWeights(graph EdgeWeightedGraph) []float64 {
	var list []float64
	for vertexID := 0; vertexID < graph.NumVertices(); vertexID++ {
		weights := graph.AdjacentWeights(vertexID)
		for i, otherVertexID := range graph.AdjacentVertices(vertexID) {
			if otherVertexID > vertexID {
				list = append(list, weights[i])
			}
		}
	}
	return list
}

// TotalGraphWeight returns total weight of a graph.
func TotalGraphWeight(graph EdgeWeightedGraph) float64 {
	return internals.SumList(AllGraphWeights(graph))
}
