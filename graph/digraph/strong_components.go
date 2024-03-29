package digraph

import (
	"github.com/DmitryBogomolov/algorithms/graph/graph"
	"github.com/DmitryBogomolov/algorithms/graph/internal/utils"
)

func findStrongComponentsCore(
	dgr graph.Graph, marked []bool, componentCount *int, componendIDs []int,
	vertexID int,
) {
	marked[vertexID] = true
	componendIDs[vertexID] = *componentCount
	for _, adjacentVertexID := range dgr.AdjacentVertices(vertexID) {
		if !marked[adjacentVertexID] {
			findStrongComponentsCore(dgr, marked, componentCount, componendIDs, adjacentVertexID)
		}
	}
}

// StrongComponents is a collection of strongly connected components in a digraph.
type StrongComponents = graph.ConnectedComponents

// FindStrongComponents returns strongly connected components in a digraph.
// This implementation uses the Kosaraju-Sharir algorithm.
// In a digraph vertices are strongly connected if they are mutually reachable.
// https://algs4.cs.princeton.edu/42digraph/KosarajuSharirSCC.java.html
func FindStrongComponents(dgr graph.Graph) StrongComponents {
	marked := make([]bool, dgr.NumVertices())
	componentCount := 0
	componendIDs := make([]int, dgr.NumVertices())
	utils.ResetList(componendIDs)
	reversedPostorder := getReversedPostorder(ReverseDigraph(dgr))
	for _, vertexID := range reversedPostorder {
		if !marked[vertexID] {
			findStrongComponentsCore(dgr, marked, &componentCount, componendIDs, vertexID)
			componentCount++
		}
	}
	return graph.NewConnectedComponents(componentCount, componendIDs)
}
