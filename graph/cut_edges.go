package graph

// CutEdge represents cut-edge in a graph.
type CutEdge = [2]int

// In a DFS tree edge *u-v* is bridge if *v* subtree has no back edges to ancestors of *u*.
//
// *dist* - distance from DFS root of *current* vertex
// *pre* - original distances
// *low* - updated distances
func findCutEdgesCore(
	result *[]CutEdge,
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
				*result = append(*result, CutEdge{current, child})
			}
		} else if child != parent && low[current] > pre[child] {
			// Update *current* distance - it can be reached faster going through *child*.
			low[current] = pre[child]
		}
	}
}

// FindCutEdges finds cut-edges in a graph.
func FindCutEdges(graph Graph) []CutEdge {
	numVertices := graph.NumVertices()
	pre := make([]int, numVertices)
	low := make([]int, numVertices)
	for v := 0; v < numVertices; v++ {
		pre[v] = -1
		low[v] = -1
	}
	var result []CutEdge
	for v := 0; v < numVertices; v++ {
		if pre[v] == -1 {
			findCutEdgesCore(&result, pre, low, 0, graph, v, v)
		}
	}
	return result
}
