package ewdigraph

import (
	"algorithms/graph/ewgraph"
	"algorithms/graph/internals"
)

// AllDigraphWeights returns all edges of an edge-weighted digraph.
func AllDigraphWeights(digraph ewgraph.EdgeWeightedGraph) []float64 {
	var list []float64
	for vertexID := 0; vertexID < digraph.NumVertices(); vertexID++ {
		weights := digraph.AdjacentWeights(vertexID)
		for i := range digraph.AdjacentVertices(vertexID) {
			list = append(list, weights[i])
		}
	}
	return list
}

// TotalDigraphWeight returns total weight of a digraph.
func TotalDigraphWeight(digraph ewgraph.EdgeWeightedGraph) float64 {
	return internals.SumList(AllDigraphWeights(digraph))
}
