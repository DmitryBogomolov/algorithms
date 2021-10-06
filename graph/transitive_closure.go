package graph

// TransitiveClosure represents transitive closure of a digraph.
// It is another digraph with the same set of vertices but with an edge from *v* to *w*
// if and only if *w* is reachable from *v*.
type TransitiveClosure struct {
	data []Paths
}

// Reachable tells if there is a directed path from *vertex1* to *vertex2*.
func (tc TransitiveClosure) Reachable(vertexID1 int, vertexID2 int) bool {
	return tc.data[vertexID1].HasPathTo(vertexID2)
}

// BuildTransitiveClosure computes transitive closure of a digraph by running depth-first search from each vertex.
// https://algs4.cs.princeton.edu/42digraph/TransitiveClosure.java.html
func BuildTransitiveClosure(graph Graph) TransitiveClosure {
	numVertices := graph.NumVertices()
	data := make([]Paths, numVertices)
	for vertexID := 0; vertexID < numVertices; vertexID++ {
		data[vertexID] = FindPathsDepthFirst(graph, vertexID)
	}
	return TransitiveClosure{data}
}
