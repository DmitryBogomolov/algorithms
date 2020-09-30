package graph

func isBipartiteCore(marked []bool, colors []bool, graph Graph, current int) bool {
	marked[current] = true
	for _, child := range graph.AdjacentVertices(current) {
		if !marked[child] {
			colors[child] = !colors[current]
			if !isBipartiteCore(marked, colors, graph, child) {
				return false
			}
		} else if colors[child] == colors[current] {
			return false
		}
	}
	return true
}

// IsBipartite shows if graph is two-colorable. Uses depth-first search.
// https://algs4.cs.princeton.edu/41graph/Bipartite.java.html
func IsBipartite(graph Graph) bool {
	numVertices := graph.NumVertices()
	marked := make([]bool, numVertices)
	colors := make([]bool, numVertices)
	for v := 0; v < numVertices; v++ {
		if !marked[v] {
			if !isBipartiteCore(marked, colors, graph, v) {
				return false
			}
		}
	}
	return true
}
