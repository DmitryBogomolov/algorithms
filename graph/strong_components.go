package graph

func getReversedPostOrderCore(list *[]int, marked []bool, digraph Graph, current int) {
	marked[current] = true
	for _, child := range digraph.AdjacentVertices(current) {
		if !marked[child] {
			getReversedPostOrderCore(list, marked, digraph, child)
		}
	}
	*list = append(*list, current)
}

func getReversedPostOrder(digraph Graph) []int {
	numVertices := digraph.NumVertices()
	marked := make([]bool, numVertices)
	var list []int
	for v := 0; v < numVertices; v++ {
		if !marked[v] {
			getReversedPostOrderCore(&list, marked, digraph, v)
		}
	}
	count := len(list)
	order := make([]int, count)
	for i := 0; i < count; i++ {
		order[i] = list[count-i-1]
	}
	return order
}

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
