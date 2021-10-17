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

// NewEdge creates Edge instance.
func NewEdge(vertexID1 int, vertexID2 int) Edge {
	return Edge{vertex1: vertexID1, vertex2: vertexID2}
}

// AllGraphEdges returns all edges of a graph.
func AllGraphEdges(gr Graph) []Edge {
	var edges []Edge
	for vertexID := 0; vertexID < gr.NumVertices(); vertexID++ {
		for _, adjacentVertexID := range gr.AdjacentVertices(vertexID) {
			if adjacentVertexID > vertexID {
				edges = append(edges, NewEdge(vertexID, adjacentVertexID))
			}
		}
	}
	return edges
}
