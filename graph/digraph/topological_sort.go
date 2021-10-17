package digraph

import (
	"algorithms/graph/graph"
	"algorithms/graph/internals/utils"
)

func getReversedPostorderCore(
	dgr graph.Graph, marked []bool, vertices *[]int,
	vertexID int,
) {
	marked[vertexID] = true
	for _, adjacentVertexID := range dgr.AdjacentVertices(vertexID) {
		if !marked[adjacentVertexID] {
			getReversedPostorderCore(dgr, marked, vertices, adjacentVertexID)
		}
	}
	*vertices = append(*vertices, vertexID)
}

func getReversedPostorder(dgr graph.Graph) []int {
	marked := make([]bool, dgr.NumVertices())
	var vertices []int
	for vertexID := 0; vertexID < dgr.NumVertices(); vertexID++ {
		if !marked[vertexID] {
			getReversedPostorderCore(dgr, marked, &vertices, vertexID)
		}
	}
	utils.ReverseList(vertices)
	return vertices
}

// TopologicalSort puts the vertices in order such that all directed edges
// point from a vertex earlier in the order to a vertex later in the order.
// https://algs4.cs.princeton.edu/42digraph/Topological.java.html
func TopologicalSort(dgr graph.Graph) []int {
	if FindDirectedCycle(dgr) != nil {
		return nil
	}
	return getReversedPostorder(dgr)
}
