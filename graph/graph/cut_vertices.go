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
	graph Graph, parentVertexID int, currentVertexID int,
) {
	children := 0
	pre[currentVertexID] = dist
	low[currentVertexID] = pre[currentVertexID]
	for _, child := range graph.AdjacentVertices(currentVertexID) {
		if pre[child] == -1 {
			children++
			findCutVerticesCore(articulation, pre, low, dist+1, graph, currentVertexID, child)
			// If *child* distance is less than *current* distance
			// then there is back edge from *child* to ancestors of *current*.
			if low[currentVertexID] > low[child] {
				low[currentVertexID] = low[child]
			}
			// *current* is not root and *child* has no back edges to ancestors of *current*.
			// If *child* had back edge then its updated distance would be less
			// than *current* original distance.
			if low[child] >= pre[currentVertexID] && parentVertexID != currentVertexID {
				articulation[currentVertexID] = true
			}
		} else if child != parentVertexID && low[currentVertexID] > pre[child] {
			// Update *current* distance - it can be reached faster going through *child*.
			low[currentVertexID] = pre[child]
		}
	}
	// *current* is root and has at least two children.
	if parentVertexID == currentVertexID && children > 1 {
		articulation[currentVertexID] = true
	}
}

// FindCutVertices finds cut-vertices in a graph.
// Cut-vertex is a vertex whose removal increases number of connected components.
// https://algs4.cs.princeton.edu/41graph/Biconnected.java.html
func FindCutVertices(graph Graph) []int {
	numVertices := graph.NumVertices()
	pre := make([]int, numVertices)
	low := make([]int, numVertices)
	articulation := make([]bool, numVertices)
	for vertexID := 0; vertexID < numVertices; vertexID++ {
		pre[vertexID] = -1
		low[vertexID] = -1
	}
	for vertexID := 0; vertexID < numVertices; vertexID++ {
		if pre[vertexID] == -1 {
			findCutVerticesCore(articulation, pre, low, 0, graph, vertexID, vertexID)
		}
	}
	var result []int
	for vertexID, flag := range articulation {
		if flag {
			result = append(result, vertexID)
		}
	}
	return result
}
