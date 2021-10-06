package graph

// In a DFS tree edge *u-v* is bridge if *v* subtree has no back edges to ancestors of *u*.
//
// *dist* - distance from DFS root of *current* vertex
// *pre* - original distances
// *low* - updated distances
func findCutEdgesCore(
	result *[]Edge,
	pre []int, low []int, dist int,
	graph Graph, parentVertexID int, currentVertexID int,
) {
	pre[currentVertexID] = dist
	low[currentVertexID] = pre[currentVertexID]
	for _, child := range graph.AdjacentVertices(currentVertexID) {
		if pre[child] == -1 {
			findCutEdgesCore(result, pre, low, dist+1, graph, currentVertexID, child)
			// If *child* distance is less than *current* distance
			// then there is back edge from *child* to ancestors of *current*.
			if low[currentVertexID] > low[child] {
				low[currentVertexID] = low[child]
			}
			// *child* has no back edges to ancestors of *current*.
			// If *child* had back edge then its updated distance would be less
			// then its original distance.
			if low[child] == pre[child] {
				*result = append(*result, Edge{currentVertexID, child})
			}
		} else if child != parentVertexID && low[currentVertexID] > pre[child] {
			// Update *current* distance - it can be reached faster going through *child*.
			low[currentVertexID] = pre[child]
		}
	}
}

// FindCutEdges finds cut-edges in a graph.
// Cut-edge is an edge whose deletion increases number of connected components.
// https://algs4.cs.princeton.edu/41graph/Bridge.java.html
func FindCutEdges(graph Graph) []Edge {
	numVertices := graph.NumVertices()
	pre := make([]int, numVertices)
	low := make([]int, numVertices)
	for vertexID := 0; vertexID < numVertices; vertexID++ {
		pre[vertexID] = -1
		low[vertexID] = -1
	}
	var result []Edge
	for vertexID := 0; vertexID < numVertices; vertexID++ {
		if pre[vertexID] == -1 {
			findCutEdgesCore(&result, pre, low, 0, graph, vertexID, vertexID)
		}
	}
	return result
}
