package ewdigraph

import (
	"github.com/DmitryBogomolov/algorithms/graph/ewgraph"
	"github.com/DmitryBogomolov/algorithms/graph/internals/utils"
)

// AllDigraphWeights returns all edges of an edge-weighted digraph.
func AllDigraphWeights(wdgr ewgraph.EdgeWeightedGraph) []float64 {
	var list []float64
	for vertexID := 0; vertexID < wdgr.NumVertices(); vertexID++ {
		weights := wdgr.AdjacentWeights(vertexID)
		for i := range wdgr.AdjacentVertices(vertexID) {
			list = append(list, weights[i])
		}
	}
	return list
}

// TotalDigraphWeight returns total weight of a digraph.
func TotalDigraphWeight(wdgr ewgraph.EdgeWeightedGraph) float64 {
	return utils.SumList(AllDigraphWeights(wdgr))
}
