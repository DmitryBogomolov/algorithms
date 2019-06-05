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
	reverseList(list)
	return list
}

// TopologicalSort performs topological sort on a digraph.
func TopologicalSort(digraph Graph) []int {
	if HasDirectedCycle(digraph) {
		return nil
	}
	return getReversedPostOrder(digraph)
}
