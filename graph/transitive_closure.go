package graph

// TransitiveClosure represents transitive closure of a digraph.
// It is another digraph with the same set of vertices but with an edge from *v* to *w*
// if and only if *w* is reachable from *v*.
type TransitiveClosure struct {
	data []VertexPaths
}

// Reachable tells if there is a directed path from *vertex1* to *vertex2*.
func (tc TransitiveClosure) Reachable(vertex1 int, vertex2 int) bool {
	return tc.data[vertex1].HasPathTo(vertex2)
}

// BuildTransitiveClosure computes transitive closure of a digraph by running depth-first search from each vertex.
// https://algs4.cs.princeton.edu/42digraph/TransitiveClosure.java.html
func BuildTransitiveClosure(graph Graph) TransitiveClosure {
	numVertices := graph.NumVertices()
	data := make([]VertexPaths, numVertices)
	for v := 0; v < numVertices; v++ {
		data[v] = FindPathsDepthFirst(graph, v)
	}
	return TransitiveClosure{data}
}
