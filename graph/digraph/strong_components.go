package digraph

import (
	"algorithms/graph/graph"
	"algorithms/graph/internals/utils"
)

func findStrongComponentsCore(componentCount *int, componendIDs []int, marked []bool, digraph graph.Graph, vertexID int) {
	marked[vertexID] = true
	componendIDs[vertexID] = *componentCount
	for _, adjacentVertexID := range digraph.AdjacentVertices(vertexID) {
		if !marked[adjacentVertexID] {
			findStrongComponentsCore(componentCount, componendIDs, marked, digraph, adjacentVertexID)
		}
	}
}

// FindStrongComponents returns strongly connected components in a digraph.
// This implementation uses the Kosaraju-Sharir algorithm.
// In a digraph vertices are strongly connected if they are mutually reachable.
// https://algs4.cs.princeton.edu/42digraph/KosarajuSharirSCC.java.html
func FindStrongComponents(digraph graph.Graph) graph.ConnectedComponents {
	componentCount := 0
	componendIDs := make([]int, digraph.NumVertices())
	utils.ResetList(componendIDs)
	marked := make([]bool, digraph.NumVertices())
	reversedPostorder := getReversedPostorder(ReverseDigraph(digraph))
	for _, vertexID := range reversedPostorder {
		if !marked[vertexID] {
			findStrongComponentsCore(&componentCount, componendIDs, marked, digraph, vertexID)
			componentCount++
		}
	}
	return graph.NewConnectedComponents(componentCount, componendIDs)
}
