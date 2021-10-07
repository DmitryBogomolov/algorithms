package graph

func isBipartiteCore(marked []bool, colors []bool, graph Graph, vertexID int) bool {
	marked[vertexID] = true
	for _, adjacentVertexID := range graph.AdjacentVertices(vertexID) {
		if !marked[adjacentVertexID] {
			colors[adjacentVertexID] = !colors[vertexID]
			if !isBipartiteCore(marked, colors, graph, adjacentVertexID) {
				return false
			}
		} else if colors[adjacentVertexID] == colors[vertexID] {
			return false
		}
	}
	return true
}

// IsBipartite tells if graph is two-colorable. Uses depth-first search.
// https://algs4.cs.princeton.edu/41graph/Bipartite.java.html
func IsBipartite(graph Graph) bool {
	numVertices := graph.NumVertices()
	marked := make([]bool, numVertices)
	colors := make([]bool, numVertices)
	for vertexID := 0; vertexID < numVertices; vertexID++ {
		if !marked[vertexID] {
			if !isBipartiteCore(marked, colors, graph, vertexID) {
				return false
			}
		}
	}
	return true
}
