package graph

import (
	"algorithms/graph/internals/utils"
)

// In a DFS tree a vertex is articulation point if:
// - it is root and has at least two children
// - it is not root and has subtree with no back edges to ancestors
func findCutVerticesCore(
	gr Graph,
	// original vertex distances
	distances []int,
	// updated vertex distances
	updatedDistances []int,
	articulation []bool,
	// distance from DFS root to current vertex
	distance int,
	parentVertexID int, vertexID int,
) {
	children := 0
	distances[vertexID] = distance
	updatedDistances[vertexID] = distances[vertexID]
	for _, adjacentVertexID := range gr.AdjacentVertices(vertexID) {
		if distances[adjacentVertexID] == -1 {
			children++
			findCutVerticesCore(gr, distances, updatedDistances, articulation, distance+1, vertexID, adjacentVertexID)
			// If child vertex distance is less than current vertex distance
			// then there is back edge from child vertex to ancestors of current vertex.
			updatedDistances[vertexID] = utils.Min(updatedDistances[vertexID], updatedDistances[adjacentVertexID])
			// If child vertex had back edge then its updated distance would be less
			// than current vertex original distance.
			if updatedDistances[adjacentVertexID] >= distances[vertexID] && parentVertexID != vertexID {
				articulation[vertexID] = true
			}
		} else if adjacentVertexID != parentVertexID {
			// Update current vertex distance - it can be reached faster going through child vertex.
			updatedDistances[vertexID] = utils.Min(updatedDistances[vertexID], distances[adjacentVertexID])
		}
	}
	// Current vertex is root and has at least two children.
	if parentVertexID == vertexID && children > 1 {
		articulation[vertexID] = true
	}
}

// FindCutVertices finds cut-vertices in a graph.
// Cut-vertex (articulation vertex) is a vertex whose removal increases number of connected components.
// A graph is biconnected if it has no articulation vertices.
// https://algs4.cs.princeton.edu/41graph/Biconnected.java.html
func FindCutVertices(gr Graph) []int {
	distances := make([]int, gr.NumVertices())
	updatedDistances := make([]int, gr.NumVertices())
	articulation := make([]bool, gr.NumVertices())
	utils.ResetList(distances)
	utils.ResetList(updatedDistances)
	for vertexID := 0; vertexID < gr.NumVertices(); vertexID++ {
		if distances[vertexID] == -1 {
			findCutVerticesCore(gr, distances, updatedDistances, articulation, 0, vertexID, vertexID)
		}
	}
	var result []int
	for vertexID, flag := range articulation {
		if flag {
			result = append(result, vertexID)
		}
	}
	return result
}
