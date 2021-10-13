package digraph

import (
	"algorithms/graph/graph"
	"algorithms/graph/internals"
)

func findDirectedCycleCore(marked []bool, stack []bool, edgeTo []int, digraph graph.Graph, vertexID int, parentVertexID int) []int {
	marked[vertexID] = true
	stack[vertexID] = true
	edgeTo[vertexID] = parentVertexID
	for _, adjacentVertexID := range digraph.AdjacentVertices(vertexID) {
		if !marked[adjacentVertexID] {
			if cycle := findDirectedCycleCore(marked, stack, edgeTo, digraph, adjacentVertexID, vertexID); cycle != nil {
				return cycle
			}
		} else if stack[adjacentVertexID] {
			var cycle []int
			for id := vertexID; id != adjacentVertexID; id = edgeTo[id] {
				cycle = append(cycle, id)
			}
			cycle = append(cycle, adjacentVertexID, vertexID)
			internals.ReverseList(cycle)
			return cycle
		}
	}
	stack[vertexID] = false
	return nil
}

// FindDirectedCycle returns directed cycle in a digraph (if such cycle exists).
// Directed cycle is a path whose first and last vertices are the same.
// https://algs4.cs.princeton.edu/42digraph/DirectedCycle.java.html
func FindDirectedCycle(digraph graph.Graph) []int {
	marked := make([]bool, digraph.NumVertices())
	stack := make([]bool, digraph.NumVertices())
	edgeTo := make([]int, digraph.NumVertices())
	internals.ResetList(edgeTo)
	for vertexID := 0; vertexID < digraph.NumVertices(); vertexID++ {
		if !marked[vertexID] {
			if cycle := findDirectedCycleCore(marked, stack, edgeTo, digraph, vertexID, -1); cycle != nil {
				return cycle
			}
		}
	}
	return nil
}
