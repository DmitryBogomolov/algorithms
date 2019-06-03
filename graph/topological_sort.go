package graph

func topologicalSortCore(list *[]int, marked []bool, digraph Graph, current int) {
	marked[current] = true
	for _, child := range digraph.AdjacentVertices(current) {
		if !marked[child] {
			topologicalSortCore(list, marked, digraph, child)
		}
	}
	*list = append(*list, current)
}

// TopologicalSort performs topological sort on a digraph.
func TopologicalSort(digraph Graph) []int {
	if HasDirectedCycle(digraph) {
		return nil
	}
	numVertices := digraph.NumVertices()
	marked := make([]bool, numVertices)
	var list []int
	for v := 0; v < numVertices; v++ {
		if !marked[v] {
			topologicalSortCore(&list, marked, digraph, v)
		}
	}
	count := len(list)
	order := make([]int, count)
	for i := 0; i < count; i++ {
		order[i] = list[count-i-1]
	}
	return order
}
