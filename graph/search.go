package graph

import (
	"container/list"
)

// VertexPaths represents paths from a vertex.
type VertexPaths struct {
	origin int
	count  int
	edgeTo []int
}

func newVertexPaths(graph Graph, vertex int) (VertexPaths, []bool) {
	count := graph.NumVertices()
	edgeTo := make([]int, count)
	marked := make([]bool, count)
	for i := range edgeTo {
		edgeTo[i] = -1
	}
	return VertexPaths{
		origin: vertex,
		edgeTo: edgeTo,
	}, marked
}

// Origin show initial vertex.
func (r VertexPaths) Origin() int {
	return r.origin
}

// HasPathTo shows if *vertex* is connected with initial vertex.
func (r VertexPaths) HasPathTo(vertex int) bool {
	return r.edgeTo[vertex] >= 0 || vertex == r.origin
}

// Count shows number of vertices connected with initial vertex.
func (r VertexPaths) Count() int {
	return r.count
}

// PathTo shows a path from source vertex to *vertex*.
func (r VertexPaths) PathTo(vertex int) []int {
	if !r.HasPathTo(vertex) {
		return nil
	}
	var stack []int
	for v := vertex; v >= 0; v = r.edgeTo[v] {
		stack = append(stack, v)
	}
	count := len(stack)
	path := make([]int, count)
	for i, v := range stack {
		path[count-i-1] = v
	}
	return path
}

func findPathsDepthFirstCore(r *VertexPaths, marked []bool, graph Graph, vertex int) {
	marked[vertex] = true
	r.count++
	for _, v := range graph.AdjacentVertices(vertex) {
		if !marked[v] {
			r.edgeTo[v] = vertex
			findPathsDepthFirstCore(r, marked, graph, v)
		}
	}
}

// FindPathsDepthFirst finds paths from "vertex" using depth-first search.
func FindPathsDepthFirst(graph Graph, vertex int) VertexPaths {
	result, marked := newVertexPaths(graph, vertex)
	findPathsDepthFirstCore(&result, marked, graph, vertex)
	return result
}

func findPathsBreadthFirstCore(r *VertexPaths, marked []bool, graph Graph, vertex int) {
	queue := list.New()
	queue.PushBack(vertex)
	marked[vertex] = true
	r.count++
	for queue.Len() > 0 {
		current := queue.Front().Value.(int)
		queue.Remove(queue.Front())
		for _, v := range graph.AdjacentVertices(current) {
			if !marked[v] {
				marked[v] = true
				r.count++
				r.edgeTo[v] = current
				queue.PushBack(v)
			}
		}
	}
}

// FindPathsBreadthFirst finds paths from "vertex" using breadth-first search.
func FindPathsBreadthFirst(graph Graph, vertex int) VertexPaths {
	result, marked := newVertexPaths(graph, vertex)
	findPathsBreadthFirstCore(&result, marked, graph, vertex)
	return result
}

// ConnectedComponents represents connected components of a graph.
type ConnectedComponents struct {
	count      int
	components []int
}

// Count returns number of connected components.
func (cc ConnectedComponents) Count() int {
	return cc.count
}

// Connected tells if two vertices are connected.
func (cc ConnectedComponents) Connected(vertex1 int, vertex2 int) bool {
	return cc.components[vertex1] == cc.components[vertex2]
}

// ComponentID returns the component to which vertex belongs.
func (cc ConnectedComponents) ComponentID(vertex int) int {
	return cc.components[vertex]
}

// Component returns vertices of connected component.
func (cc ConnectedComponents) Component(component int) []int {
	var ret []int
	for i := 0; i < len(cc.components); i++ {
		if cc.components[i] == component {
			ret = append(ret, i)
		}
	}
	return ret
}

func findConnectedComponentsCore(cc *ConnectedComponents, marked []bool, graph Graph, current int) {
	marked[current] = true
	cc.components[current] = cc.count
	for _, child := range graph.AdjacentVertices(current) {
		if !marked[child] {
			findConnectedComponentsCore(cc, marked, graph, child)
		}
	}
}

// FindConnectedComponents finds connected components in a graph.
func FindConnectedComponents(graph Graph) ConnectedComponents {
	numVertices := graph.NumVertices()
	result := ConnectedComponents{
		components: make([]int, numVertices),
	}
	marked := make([]bool, numVertices)
	for v := 0; v < numVertices; v++ {
		if !marked[v] {
			findConnectedComponentsCore(&result, marked, graph, v)
			result.count++
		}
	}
	return result
}

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

func isBipartiteCore(marked []bool, colors []bool, graph Graph, current int) bool {
	marked[current] = true
	for _, child := range graph.AdjacentVertices(current) {
		if !marked[child] {
			colors[child] = !colors[current]
			if !isBipartiteCore(marked, colors, graph, child) {
				return false
			}
		} else if colors[child] == colors[current] {
			return false
		}
	}
	return true
}

// IsBipartite shows if graph is two-colorable.
func IsBipartite(graph Graph) bool {
	numVertices := graph.NumVertices()
	marked := make([]bool, numVertices)
	colors := make([]bool, numVertices)
	for v := 0; v < numVertices; v++ {
		if !marked[v] {
			if !isBipartiteCore(marked, colors, graph, v) {
				return false
			}
		}
	}
	return true
}

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
