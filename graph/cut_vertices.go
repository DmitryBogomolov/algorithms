package graph

// In a DFS tree a vertex *v* is articulation point if:
// - *v* is root and has at least two children
// - *v* is not root and has subtree with no back edges to ancestors of *v*
//
// *dist* - distance from DFS root of *current* vertex
// *pre* - original distances
// *low* - updated distances
func findCutVerticesCore(
	articulation []bool,
	pre []int, low []int, dist int,
	graph Graph, parent int, current int,
) {
	children := 0
	pre[current] = dist
	low[current] = pre[current]
	for _, child := range graph.AdjacentVertices(current) {
		if pre[child] == -1 {
			children++
			findCutVerticesCore(articulation, pre, low, dist+1, graph, current, child)
			// If *child* distance is less than *current* distance
			// then there is back edge from *child* to ancestors of *current*.
			if low[current] > low[child] {
				low[current] = low[child]
			}
			// *current* is not root and *child* has no back edges to ancestors of *current*.
			// If *child* had back edge then its updated distance would be less
			// than *current* original distance.
			if low[child] >= pre[current] && parent != current {
				articulation[current] = true
			}
		} else if child != parent && low[current] > pre[child] {
			// Update *current* distance - it can be reached faster going through *child*.
			low[current] = pre[child]
		}
	}
	// *current* is root and has at least two children.
	if parent == current && children > 1 {
		articulation[current] = true
	}
}

// FindCutVertices finds cut-vertices in a graph.
func FindCutVertices(graph Graph) []int {
	numVertices := graph.NumVertices()
	pre := make([]int, numVertices)
	low := make([]int, numVertices)
	articulation := make([]bool, numVertices)
	for v := 0; v < numVertices; v++ {
		pre[v] = -1
		low[v] = -1
	}
	for v := 0; v < numVertices; v++ {
		if pre[v] == -1 {
			findCutVerticesCore(articulation, pre, low, 0, graph, v, v)
		}
	}
	var result []int
	for v, flag := range articulation {
		if flag {
			result = append(result, v)
		}
	}
	return result
}
