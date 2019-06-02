package graph

func hasCycleCore(marked []bool, graph Graph, parent int, current int) bool {
	marked[current] = true
	for _, child := range graph.AdjacentVertices(current) {
		if !marked[child] {
			if hasCycleCore(marked, graph, current, child) {
				return true
			}
		} else if child != parent {
			return true
		}
	}
	return false
}

// HasCycle shows if there is a cycle in a graph.
func HasCycle(graph Graph) bool {
	numVertices := graph.NumVertices()
	marked := make([]bool, numVertices)
	for v := 0; v < numVertices; v++ {
		if !marked[v] {
			if hasCycleCore(marked, graph, -1, v) {
				return true
			}
		}
	}
	return false
}
