package graph

func hasDirectedCycleCore(marked []bool, stack []bool, digraph Graph, current int) bool {
	marked[current] = true
	stack[current] = true
	for _, child := range digraph.AdjacentVertices(current) {
		if !marked[child] {
			if hasDirectedCycleCore(marked, stack, digraph, child) {
				return true
			}
		} else if stack[child] {
			return true
		}
	}
	stack[current] = false
	return false
}

// HasDirectedCycle shows if there is a cycle in a directed graph.
func HasDirectedCycle(digraph Graph) bool {
	numVertices := digraph.NumVertices()
	marked := make([]bool, numVertices)
	stack := make([]bool, numVertices)
	for v := 0; v < numVertices; v++ {
		if !marked[v] {
			if hasDirectedCycleCore(marked, stack, digraph, v) {
				return true
			}
		}
	}
	return false
}
