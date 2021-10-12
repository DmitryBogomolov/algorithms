package graph

// In a DFS tree edge "u-v" is bridge if "v" subtree has no back edges to ancestors of "u".
func findCutEdgesCore(
	cutEdges *[]Edge,
	// original vertex distances
	distances []int,
	// updated vertex distances
	updatedDistances []int,
	// distance from DFS root to current vertex
	distance int,
	graph Graph, parentVertexID int, vertexID int,
) {
	distances[vertexID] = distance
	updatedDistances[vertexID] = distances[vertexID]
	for _, adjacentVertexID := range graph.AdjacentVertices(vertexID) {
		if distances[adjacentVertexID] == -1 {
			findCutEdgesCore(cutEdges, distances, updatedDistances, distance+1, graph, vertexID, adjacentVertexID)
			// If child vertex distance is less than current vertex distance
			// then there is back edge from child vertex to ancestors of current vertex.
			updatedDistances[vertexID] = min(updatedDistances[vertexID], updatedDistances[adjacentVertexID])
			// If child vertex had back edge then its updated distance would be less then its original distance.
			if updatedDistances[adjacentVertexID] == distances[adjacentVertexID] {
				*cutEdges = append(*cutEdges, NewEdge(vertexID, adjacentVertexID))
			}
		} else if adjacentVertexID != parentVertexID {
			// Update current vertex distance - it can be reached faster going through child vertex.
			updatedDistances[vertexID] = min(updatedDistances[vertexID], distances[adjacentVertexID])
		}
	}
}

// FindCutEdges finds cut-edges in a graph.
// Cut-edge is an edge whose deletion increases number of connected components.
// An edge is a bridge iif it is not contained in any cycle.
// https://algs4.cs.princeton.edu/41graph/Bridge.java.html
func FindCutEdges(graph Graph) []Edge {
	distances := make([]int, graph.NumVertices())
	updatedDistances := make([]int, graph.NumVertices())
	resetList(distances)
	resetList(updatedDistances)
	var cutEdges []Edge
	for vertexID := 0; vertexID < graph.NumVertices(); vertexID++ {
		if distances[vertexID] == -1 {
			findCutEdgesCore(&cutEdges, distances, updatedDistances, 0, graph, vertexID, vertexID)
		}
	}
	return cutEdges
}
