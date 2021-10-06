package graph

func getReversedPostOrderCore(list *[]int, marked []bool, digraph Graph, vertexID int) {
	marked[vertexID] = true
	for _, adjacentVertexID := range digraph.AdjacentVertices(vertexID) {
		if !marked[adjacentVertexID] {
			getReversedPostOrderCore(list, marked, digraph, adjacentVertexID)
		}
	}
	*list = append(*list, vertexID)
}

func getReversedPostOrder(digraph Graph) []int {
	numVertices := digraph.NumVertices()
	marked := make([]bool, numVertices)
	var list []int
	for vertexID := 0; vertexID < numVertices; vertexID++ {
		if !marked[vertexID] {
			getReversedPostOrderCore(&list, marked, digraph, vertexID)
		}
	}
	reverseList(list)
	return list
}

// TopologicalSort puts the vertices in order such that all directed edges
// point from a vertex earlier in the order to a vertex later in the order.
// https://algs4.cs.princeton.edu/42digraph/Topological.java.html
func TopologicalSort(digraph Graph) []int {
	if HasDirectedCycle(digraph) {
		return nil
	}
	return getReversedPostOrder(digraph)
}
