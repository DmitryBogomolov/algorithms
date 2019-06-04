package graph

// Graph represents an undirected graph.
type Graph interface {
	NumVertices() int
	NumEdges() int
	AdjacentVertices(vertex int) []int
}
