package graph

import (
	"algorithms/graph/internals/utils"
)

func findCycleCore(
	gr Graph, marked []bool, edgeTo []int,
	vertexID int, parentVertexID int,
) []int {
	marked[vertexID] = true
	edgeTo[vertexID] = parentVertexID
	for _, adjacentVertexID := range gr.AdjacentVertices(vertexID) {
		if !marked[adjacentVertexID] {
			if cycle := findCycleCore(gr, marked, edgeTo, adjacentVertexID, vertexID); cycle != nil {
				return cycle
			}
		} else if adjacentVertexID != parentVertexID {
			// If next (adjacent to current) vertex is already visited and it is not previous
			// (parent to current) vertex then it makes a cycle.
			var cycle []int
			for id := vertexID; id != adjacentVertexID; id = edgeTo[id] {
				cycle = append(cycle, id)
			}
			cycle = append(cycle, adjacentVertexID, vertexID)
			utils.ReverseList(cycle)
			return cycle
		}
	}
	return nil
}

// FindCycle returns cycle in a graph (if such cycle exists).
// Cycle is a path whose first and last vertices are the same.
// https://algs4.cs.princeton.edu/41graph/Cycle.java.html
func FindCycle(gr Graph) []int {
	marked := make([]bool, gr.NumVertices())
	edgeTo := make([]int, gr.NumVertices())
	utils.ResetList(edgeTo)
	for vertexID := 0; vertexID < gr.NumVertices(); vertexID++ {
		if !marked[vertexID] {
			if cycle := findCycleCore(gr, marked, edgeTo, vertexID, -1); cycle != nil {
				return cycle
			}
		}
	}
	return nil
}
