package graph

func findStrongComponentsCore(cc *ConnectedComponents, marked []bool, digraph Graph, vertexID int) {
	marked[vertexID] = true
	cc.components[vertexID] = cc.count
	for _, adjacentVertexID := range digraph.AdjacentVertices(vertexID) {
		if !marked[adjacentVertexID] {
			findStrongComponentsCore(cc, marked, digraph, adjacentVertexID)
		}
	}
}

// FindStrongComponents finds strongly connected components in a digraph.
// https://algs4.cs.princeton.edu/42digraph/KosarajuSharirSCC.java.html
func FindStrongComponents(digraph Graph) ConnectedComponents {
	numVertices := digraph.NumVertices()
	result := newConnectedComponents(numVertices)
	marked := make([]bool, numVertices)
	reversedPostOrder := getReversedPostOrder(ReverseDigraph(digraph))
	for _, vertexID := range reversedPostOrder {
		if !marked[vertexID] {
			findStrongComponentsCore(&result, marked, digraph, vertexID)
			result.count++
		}
	}
	return result
}
