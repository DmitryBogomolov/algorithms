package graph

func findStrongComponentsCore(cc *ConnectedComponents, marked []bool, digraph Graph, current int) {
	marked[current] = true
	cc.components[current] = cc.count
	for _, child := range digraph.AdjacentVertices(current) {
		if !marked[child] {
			findStrongComponentsCore(cc, marked, digraph, child)
		}
	}
}

// FindStrongComponents finds strongly connected components in a digraph.
func FindStrongComponents(digraph Graph) ConnectedComponents {
	numVertices := digraph.NumVertices()
	result := newConnectedComponents(numVertices)
	marked := make([]bool, numVertices)
	reversedPostorder := getReversedPostOrder(Reverse(digraph))
	for _, v := range reversedPostorder {
		if !marked[v] {
			findStrongComponentsCore(&result, marked, digraph, v)
			result.count++
		}
	}
	return result
}
