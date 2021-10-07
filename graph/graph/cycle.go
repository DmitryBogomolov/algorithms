package graph

func hasCycleCore(marked []bool, graph Graph, parentVertexID int, currentVertexID int) bool {
	marked[currentVertexID] = true
	for _, childVertexID := range graph.AdjacentVertices(currentVertexID) {
		if !marked[childVertexID] {
			if hasCycleCore(marked, graph, currentVertexID, childVertexID) {
				return true
			}
		} else if childVertexID != parentVertexID {
			return true
		}
	}
	return false
}

// HasCycle tells if there is a cycle in a graph.
// Cycle is a path whose first and last vertices are the same.
// https://algs4.cs.princeton.edu/41graph/Cycle.java.html
func HasCycle(graph Graph) bool {
	numVertices := graph.NumVertices()
	marked := make([]bool, numVertices)
	for vertexID := 0; vertexID < numVertices; vertexID++ {
		if !marked[vertexID] {
			if hasCycleCore(marked, graph, -1, vertexID) {
				return true
			}
		}
	}
	return false
}
