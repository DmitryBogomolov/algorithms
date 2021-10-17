package graph

import (
	"algorithms/graph/internals/utils"
)

// In a DFS tree edge "u-v" is bridge if "v" subtree has no back edges to ancestors of "u".
func findCutEdgesCore(
	gr Graph,
	// original vertex distances
	distances []int,
	// updated vertex distances
	updatedDistances []int,
	cutEdges *[]Edge,
	// distance from DFS root to current vertex
	distance int,
	parentVertexID int, vertexID int,
) {
	distances[vertexID] = distance
	updatedDistances[vertexID] = distances[vertexID]
	for _, adjacentVertexID := range gr.AdjacentVertices(vertexID) {
		if distances[adjacentVertexID] == -1 {
			findCutEdgesCore(gr, distances, updatedDistances, cutEdges, distance+1, vertexID, adjacentVertexID)
			// If child vertex distance is less than current vertex distance
			// then there is back edge from child vertex to ancestors of current vertex.
			updatedDistances[vertexID] = utils.Min(updatedDistances[vertexID], updatedDistances[adjacentVertexID])
			// If child vertex had back edge then its updated distance would be less then its original distance.
			if updatedDistances[adjacentVertexID] == distances[adjacentVertexID] {
				*cutEdges = append(*cutEdges, NewEdge(vertexID, adjacentVertexID))
			}
		} else if adjacentVertexID != parentVertexID {
			// Update current vertex distance - it can be reached faster going through child vertex.
			updatedDistances[vertexID] = utils.Min(updatedDistances[vertexID], distances[adjacentVertexID])
		}
	}
}

// FindCutEdges finds cut-edges in a graph.
// Cut-edge is an edge whose deletion increases number of connected components.
// An edge is a bridge iif it is not contained in any cycle.
// https://algs4.cs.princeton.edu/41graph/Bridge.java.html
func FindCutEdges(gr Graph) []Edge {
	distances := make([]int, gr.NumVertices())
	updatedDistances := make([]int, gr.NumVertices())
	utils.ResetList(distances)
	utils.ResetList(updatedDistances)
	var cutEdges []Edge
	for vertexID := 0; vertexID < gr.NumVertices(); vertexID++ {
		if distances[vertexID] == -1 {
			findCutEdgesCore(gr, distances, updatedDistances, &cutEdges, 0, vertexID, vertexID)
		}
	}
	return cutEdges
}
