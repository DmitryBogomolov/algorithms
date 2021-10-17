package graph

func isBipartiteCore(
	gr Graph, marked []bool, colors []bool,
	vertexID int, parentColor bool,
) bool {
	marked[vertexID] = true
	color := !parentColor
	colors[vertexID] = color
	for _, adjacentVertexID := range gr.AdjacentVertices(vertexID) {
		if !marked[adjacentVertexID] {
			colors[adjacentVertexID] = !colors[vertexID]
			if !isBipartiteCore(gr, marked, colors, adjacentVertexID, color) {
				return false
			}
		} else if colors[adjacentVertexID] == colors[vertexID] {
			return false
		}
	}
	return true
}

// IsBipartite tells if graph is two-colorable - such that each no edge connects vertices of the same color.
// In a bipartite graph vertices can be divided into two sets such that all edges connect a vertex
// in one set with a vertex in other set.
// https://algs4.cs.princeton.edu/41graph/Bipartite.java.html
func IsBipartite(gr Graph) bool {
	marked := make([]bool, gr.NumVertices())
	colors := make([]bool, gr.NumVertices())
	for vertexID := 0; vertexID < gr.NumVertices(); vertexID++ {
		if !marked[vertexID] {
			if !isBipartiteCore(gr, marked, colors, vertexID, true) {
				return false
			}
		}
	}
	return true
}
