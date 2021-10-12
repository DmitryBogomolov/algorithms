package graph

func findCycleCore(edgeTo []int, graph Graph, marked []bool, vertexID int, parentVertexID int) []int {
	marked[vertexID] = true
	edgeTo[vertexID] = parentVertexID
	for _, adjacentVertexID := range graph.AdjacentVertices(vertexID) {
		if !marked[adjacentVertexID] {
			cycle := findCycleCore(edgeTo, graph, marked, adjacentVertexID, vertexID)
			if cycle != nil {
				return cycle
			}
		} else if adjacentVertexID != parentVertexID {
			// If next (adjacent to current) vertex is already visited and it is not previous
			// (parent to current) vertex then it makes a cycle.
			var cycle []int
			for id := vertexID; id != adjacentVertexID; id = edgeTo[id] {
				cycle = append(cycle, id)
			}
			cycle = append(cycle, adjacentVertexID, vertexID)
			reverseList(cycle)
			return cycle
		}
	}
	return nil
}

// FindCycle returns cycle in a graph (if such cycle exists).
// Cycle is a path whose first and last vertices are the same.
// https://algs4.cs.princeton.edu/41graph/Cycle.java.html
func FindCycle(graph Graph) []int {
	marked := make([]bool, graph.NumVertices())
	edgeTo := make([]int, graph.NumVertices())
	resetList(edgeTo)
	for vertexID := 0; vertexID < graph.NumVertices(); vertexID++ {
		if !marked[vertexID] {
			cycle := findCycleCore(edgeTo, graph, marked, vertexID, -1)
			if cycle != nil {
				return cycle
			}
		}
	}
	return nil
}
