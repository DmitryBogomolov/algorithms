package digraph

import "algorithms/graph/graph"

// AllDigraphEdges returns all edges of a digraph.
func AllDigraphEdges(dgr graph.Graph) []graph.Edge {
	var edges []graph.Edge
	for vertexID := 0; vertexID < dgr.NumVertices(); vertexID++ {
		for _, otherVertexID := range dgr.AdjacentVertices(vertexID) {
			edges = append(edges, graph.NewEdge(vertexID, otherVertexID))
		}
	}
	return edges
}
