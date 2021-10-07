package digraph

import "algorithms/graph/graph"

// AllDigraphEdges returns all edges of a digraph.
func AllDigraphEdges(digraph graph.Graph) []graph.Edge {
	var edges []graph.Edge
	for vertexID := 0; vertexID < digraph.NumVertices(); vertexID++ {
		for _, otherVertexID := range digraph.AdjacentVertices(vertexID) {
			edges = append(edges, graph.NewEdge(vertexID, otherVertexID))
		}
	}
	return edges
}
