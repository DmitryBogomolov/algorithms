package digraph

import (
	"algorithms/graph/graph"
	"algorithms/graph/internals/utils"
)

func getReversedPostorderCore(list *[]int, marked []bool, digraph graph.Graph, vertexID int) {
	marked[vertexID] = true
	for _, adjacentVertexID := range digraph.AdjacentVertices(vertexID) {
		if !marked[adjacentVertexID] {
			getReversedPostorderCore(list, marked, digraph, adjacentVertexID)
		}
	}
	*list = append(*list, vertexID)
}

func getReversedPostorder(digraph graph.Graph) []int {
	marked := make([]bool, digraph.NumVertices())
	var list []int
	for vertexID := 0; vertexID < digraph.NumVertices(); vertexID++ {
		if !marked[vertexID] {
			getReversedPostorderCore(&list, marked, digraph, vertexID)
		}
	}
	utils.ReverseList(list)
	return list
}

// TopologicalSort puts the vertices in order such that all directed edges
// point from a vertex earlier in the order to a vertex later in the order.
// https://algs4.cs.princeton.edu/42digraph/Topological.java.html
func TopologicalSort(digraph graph.Graph) []int {
	if FindDirectedCycle(digraph) != nil {
		return nil
	}
	return getReversedPostorder(digraph)
}
