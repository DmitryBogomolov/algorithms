package graph

// Graph is a set of vertices and edges.
type Graph interface {
	// NumVertices gets number of graph vertices.
	NumVertices() int
	// NumEdges gets number of graph edges.
	NumEdges() int
	// AdjacentVertices returns vertices adjacent to the vertex.
	AdjacentVertices(vertex int) []int
}

// EdgeWeightedGraph is a graph where each edge has an associated weight.
type EdgeWeightedGraph interface {
	Graph
	// AdjacentWeights returns weights of edges adjacent to the vertex.
	AdjacentWeights(vertex int) []float64
}

// Edge is a pair of connected vertices in a graph.
type Edge struct {
	// Vertex1 is one of edge vertices.
	Vertex1 int
	// Vertex2 is one of edge vertices.
	Vertex2 int
}

// AllGraphEdges returns all edges of a graph.
func AllGraphEdges(graph Graph) []Edge {
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

// AllDigraphEdges returns all edges of a digraph.
func AllDigraphEdges(digraph Graph) []Edge {
	var ret []Edge
	for v := 0; v < digraph.NumVertices(); v++ {
		for _, w := range digraph.AdjacentVertices(v) {
			ret = append(ret, Edge{v, w})
		}
	}
	return ret
}

// AllGraphWeights returns all edge weights of an edge-weighted graph.
func AllGraphWeights(graph EdgeWeightedGraph) []float64 {
	var ret []float64
	for v := 0; v < graph.NumVertices(); v++ {
		weights := graph.AdjacentWeights(v)
		for i, w := range graph.AdjacentVertices(v) {
			if w > v {
				ret = append(ret, weights[i])
			}
		}
	}
	return ret
}

// AllDigraphWeights returns all edges of an edge-weighted digraph.
func AllDigraphWeights(digraph EdgeWeightedGraph) []float64 {
	var ret []float64
	for v := 0; v < digraph.NumVertices(); v++ {
		weights := digraph.AdjacentWeights(v)
		for i := range digraph.AdjacentVertices(v) {
			ret = append(ret, weights[i])
		}
	}
	return ret
}

func sumList(list []float64) float64 {
	sum := 0.0
	for _, item := range list {
		sum += item
	}
	return sum
}

// TotalGraphWeight returns total weight of a graph.
func TotalGraphWeight(graph EdgeWeightedGraph) float64 {
	return sumList(AllGraphWeights(graph))
}

// TotalDigraphWeight returns total weight of a digraph.
func TotalDigraphWeight(digraph EdgeWeightedGraph) float64 {
	return sumList(AllDigraphWeights(digraph))
}
