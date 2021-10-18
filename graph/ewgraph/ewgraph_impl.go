package ewgraph

import "github.com/DmitryBogomolov/algorithms/graph/graph"

// ImplEdgeWeightedGraph is an implementaion of EdgeWeightedGraph.
type ImplEdgeWeightedGraph struct {
	graph.Graph
	weights [][]float64
}

// AdjacentWeights returns weights of edges adjacent to the vertex.
func (wgr *ImplEdgeWeightedGraph) AdjacentWeights(vertexID int) []float64 {
	return wgr.weights[vertexID]
}

// NewImplEdgeWeightedGraph creates instance of EdgeWeightedGraph.
func NewImplEdgeWeightedGraph(
	numVertices int, numEdges int, adjacency [][]int, weights [][]float64,
) *ImplEdgeWeightedGraph {
	return &ImplEdgeWeightedGraph{
		Graph:   graph.NewImplGraph(numVertices, numEdges, adjacency),
		weights: weights,
	}
}
