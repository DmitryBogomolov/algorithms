package graph

// Graph is a set of vertices and edges.
type Graph interface {
	// NumVertices gets number of graph vertices.
	NumVertices() int
	// NumEdges gets number of graph edges.
	NumEdges() int
	// AdjacentVertices returns vertices adjacent to the vertex.
	AdjacentVertices(vertexID int) []int
}

// EdgeWeightedGraph is a graph where each edge has an associated weight.
type EdgeWeightedGraph interface {
	Graph
	// AdjacentWeights returns weights of edges adjacent to the vertex.
	AdjacentWeights(vertexID int) []float64
}

// Edge is a pair of connected vertices in a graph.
type Edge struct {
	vertex1 int
	vertex2 int
}

// Vertex1 gets one of edge vertices.
func (edge Edge) Vertex1() int {
	return edge.vertex1
}

// Vertex2 gets one of edge vertices.
func (edge Edge) Vertex2() int {
	return edge.vertex2
}

// AllGraphEdges returns all edges of a graph.
func AllGraphEdges(graph Graph) []Edge {
	var edges []Edge
	for vertexID := 0; vertexID < graph.NumVertices(); vertexID++ {
		for _, otherVertexID := range graph.AdjacentVertices(vertexID) {
			if otherVertexID > vertexID {
				edges = append(edges, Edge{vertexID, otherVertexID})
			}
		}
	}
	return edges
}

// AllDigraphEdges returns all edges of a digraph.
func AllDigraphEdges(digraph Graph) []Edge {
	var edges []Edge
	for vertexID := 0; vertexID < digraph.NumVertices(); vertexID++ {
		for _, otherVertexID := range digraph.AdjacentVertices(vertexID) {
			edges = append(edges, Edge{vertexID, otherVertexID})
		}
	}
	return edges
}

// AllGraphWeights returns all edge weights of an edge-weighted graph.
func AllGraphWeights(graph EdgeWeightedGraph) []float64 {
	var list []float64
	for vertexID := 0; vertexID < graph.NumVertices(); vertexID++ {
		weights := graph.AdjacentWeights(vertexID)
		for i, otherVertexID := range graph.AdjacentVertices(vertexID) {
			if otherVertexID > vertexID {
				list = append(list, weights[i])
			}
		}
	}
	return list
}

// AllDigraphWeights returns all edges of an edge-weighted digraph.
func AllDigraphWeights(digraph EdgeWeightedGraph) []float64 {
	var list []float64
	for vertexID := 0; vertexID < digraph.NumVertices(); vertexID++ {
		weights := digraph.AdjacentWeights(vertexID)
		for i := range digraph.AdjacentVertices(vertexID) {
			list = append(list, weights[i])
		}
	}
	return list
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
