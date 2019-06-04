package graph

// TransitiveClosure represents transitive closure of a graph.
type TransitiveClosure struct {
	data []VertexPaths
}

// Reachable tells if there is a directed path from *vertex1* to *vertex2*.
func (tc TransitiveClosure) Reachable(vertex1 int, vertex2 int) bool {
	return tc.data[vertex1].HasPathTo(vertex2)
}

// BuildTransitiveClosure builds transitive closure for a graph.
func BuildTransitiveClosure(graph Graph) TransitiveClosure {
	numVertices := graph.NumVertices()
	data := make([]VertexPaths, numVertices)
	for v := 0; v < numVertices; v++ {
		data[v] = FindPathsDepthFirst(graph, v)
	}
	return TransitiveClosure{data}
}
