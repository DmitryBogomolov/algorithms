package digraph

import (
	"github.com/DmitryBogomolov/algorithms/graph/graph"
	"github.com/DmitryBogomolov/algorithms/graph/internals/utils"
)

func findDirectedCycleCore(
	dgr graph.Graph, marked []bool, stack []bool, edgeTo []int,
	vertexID int, parentVertexID int,
) []int {
	marked[vertexID] = true
	stack[vertexID] = true
	edgeTo[vertexID] = parentVertexID
	for _, adjacentVertexID := range dgr.AdjacentVertices(vertexID) {
		if !marked[adjacentVertexID] {
			if cycle := findDirectedCycleCore(dgr, marked, stack, edgeTo, adjacentVertexID, vertexID); cycle != nil {
				return cycle
			}
		} else if stack[adjacentVertexID] {
			var cycle []int
			for id := vertexID; id != adjacentVertexID; id = edgeTo[id] {
				cycle = append(cycle, id)
			}
			cycle = append(cycle, adjacentVertexID, vertexID)
			utils.ReverseList(cycle)
			return cycle
		}
	}
	stack[vertexID] = false
	return nil
}

// FindDirectedCycle returns directed cycle in a digraph (if such cycle exists).
// Directed cycle is a path whose first and last vertices are the same.
// https://algs4.cs.princeton.edu/42digraph/DirectedCycle.java.html
func FindDirectedCycle(dgr graph.Graph) []int {
	marked := make([]bool, dgr.NumVertices())
	stack := make([]bool, dgr.NumVertices())
	edgeTo := make([]int, dgr.NumVertices())
	utils.ResetList(edgeTo)
	for vertexID := 0; vertexID < dgr.NumVertices(); vertexID++ {
		if !marked[vertexID] {
			if cycle := findDirectedCycleCore(dgr, marked, stack, edgeTo, vertexID, -1); cycle != nil {
				return cycle
			}
		}
	}
	return nil
}
