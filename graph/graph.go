package graph

// Graph represents a undirected graph.
type Graph interface {
	NumVertices() int
	NumEdges() int
	AdjacentVertices(vertex int) []int
}
