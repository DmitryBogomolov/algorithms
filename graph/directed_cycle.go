package graph

func hasDirectedCycleCore(marked []bool, stack []bool, digraph Graph, vertexID int) bool {
	marked[vertexID] = true
	stack[vertexID] = true
	for _, adjacentVertexID := range digraph.AdjacentVertices(vertexID) {
		if !marked[adjacentVertexID] {
			if hasDirectedCycleCore(marked, stack, digraph, adjacentVertexID) {
				return true
			}
		} else if stack[adjacentVertexID] {
			return true
		}
	}
	stack[vertexID] = false
	return false
}

// HasDirectedCycle shows if there is a cycle in a digraph.
// https://algs4.cs.princeton.edu/44sp/DirectedCycle.java.html
func HasDirectedCycle(digraph Graph) bool {
	numVertices := digraph.NumVertices()
	marked := make([]bool, numVertices)
	stack := make([]bool, numVertices)
	for vertexID := 0; vertexID < numVertices; vertexID++ {
		if !marked[vertexID] {
			if hasDirectedCycleCore(marked, stack, digraph, vertexID) {
				return true
			}
		}
	}
	return false
}
