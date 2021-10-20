package ewgraph

import (
	"github.com/DmitryBogomolov/algorithms/graph/graph"
	"github.com/DmitryBogomolov/algorithms/graph/internal/utils"
)

// EdgeWeightedGraph is a graph where each edge has an associated weight.
type EdgeWeightedGraph interface {
	graph.Graph
	// AdjacentWeights returns weights of edges adjacent to the vertex.
	AdjacentWeights(vertexID int) []float64
}

// AllGraphWeights returns all edge weights of an edge-weighted graph.
func AllGraphWeights(wgr EdgeWeightedGraph) []float64 {
	var list []float64
	for vertexID := 0; vertexID < wgr.NumVertices(); vertexID++ {
		weights := wgr.AdjacentWeights(vertexID)
		for i, adjacentVertexID := range wgr.AdjacentVertices(vertexID) {
			if adjacentVertexID > vertexID {
				list = append(list, weights[i])
			}
		}
	}
	return list
}

// TotalGraphWeight returns total weight of a grapgrh.
func TotalGraphWeight(wgr EdgeWeightedGraph) float64 {
	return utils.SumList(AllGraphWeights(wgr))
}
