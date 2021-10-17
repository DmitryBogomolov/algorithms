package graph

// ImplGraph is an implementation of Graph.
type ImplGraph struct {
	numVertices int
	numEdges    int
	adjacency   [][]int
}

// NumVertices gets number of graph vertices.
func (gr *ImplGraph) NumVertices() int {
	return gr.numVertices
}

// NumEdges gets number of graph edges.
func (gr *ImplGraph) NumEdges() int {
	return gr.numEdges
}

// AdjacentVertices returns vertices adjacent to the vertex.
func (gr *ImplGraph) AdjacentVertices(vertexID int) []int {
	return gr.adjacency[vertexID]
}

// NewImplGraph creates instance of ImplGraph.
func NewImplGraph(numVertices int, numEdges int, adjacency [][]int) *ImplGraph {
	return &ImplGraph{
		numVertices: numVertices,
		numEdges:    numEdges,
		adjacency:   adjacency,
	}
}
