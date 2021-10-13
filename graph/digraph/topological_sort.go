package digraph

import (
	"algorithms/graph/graph"
	"algorithms/graph/internals"
)

func topologicalSortCore(list *[]int, marked []bool, digraph graph.Graph, vertexID int) {
	marked[vertexID] = true
	for _, adjacentVertexID := range digraph.AdjacentVertices(vertexID) {
		if !marked[adjacentVertexID] {
			topologicalSortCore(list, marked, digraph, adjacentVertexID)
		}
	}
	*list = append(*list, vertexID)
}

// TopologicalSort puts the vertices in order such that all directed edges
// point from a vertex earlier in the order to a vertex later in the order.
// https://algs4.cs.princeton.edu/42digraph/Topological.java.html
func TopologicalSort(digraph graph.Graph) []int {
	if FindDirectedCycle(digraph) != nil {
		return nil
	}

	marked := make([]bool, digraph.NumVertices())
	var list []int
	for vertexID := 0; vertexID < digraph.NumVertices(); vertexID++ {
		if !marked[vertexID] {
			topologicalSortCore(&list, marked, digraph, vertexID)
		}
	}
	internals.ReverseList(list)
	return list
}
