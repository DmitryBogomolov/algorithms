package graph

// Graph represents a graph.
type Graph interface {
	NumVertices() int
	NumEdges() int
	AdjacentVertices(vertex int) []int
}

// Edge represents cut-edge in a graph.
type Edge = [2]int

// Edges returns all edges of an undirected graph.
func Edges(graph Graph) []Edge {
	var ret []Edge
	for v := 0; v < graph.NumVertices(); v++ {
		for _, w := range graph.AdjacentVertices(v) {
			if w > v {
				ret = append(ret, Edge{v, w})
			}
		}
	}
	return ret
}

// DirectedEdges returns all edges of a directed graph.
func DirectedEdges(digraph Graph) []Edge {
	var ret []Edge
	for v := 0; v < digraph.NumVertices(); v++ {
		for _, w := range digraph.AdjacentVertices(v) {
			ret = append(ret, Edge{v, w})
		}
	}
	return ret
}
