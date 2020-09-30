package graph

// In a DFS tree edge *u-v* is bridge if *v* subtree has no back edges to ancestors of *u*.
//
// *dist* - distance from DFS root of *current* vertex
// *pre* - original distances
// *low* - updated distances
func findCutEdgesCore(
	result *[]Edge,
	pre []int, low []int, dist int,
	graph Graph, parent int, current int,
) {
	pre[current] = dist
	low[current] = pre[current]
	for _, child := range graph.AdjacentVertices(current) {
		if pre[child] == -1 {
			findCutEdgesCore(result, pre, low, dist+1, graph, current, child)
			// If *child* distance is less than *current* distance
			// then there is back edge from *child* to ancestors of *current*.
			if low[current] > low[child] {
				low[current] = low[child]
			}
			// *child* has no back edges to ancestors of *current*.
			// If *child* had back edge then its updated distance would be less
			// then its original distance.
			if low[child] == pre[child] {
				*result = append(*result, Edge{current, child})
			}
		} else if child != parent && low[current] > pre[child] {
			// Update *current* distance - it can be reached faster going through *child*.
			low[current] = pre[child]
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
	for v := 0; v < numVertices; v++ {
		pre[v] = -1
		low[v] = -1
	}
	var result []Edge
	for v := 0; v < numVertices; v++ {
		if pre[v] == -1 {
			findCutEdgesCore(&result, pre, low, 0, graph, v, v)
		}
	}
	return result
}
