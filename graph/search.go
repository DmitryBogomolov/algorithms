package graph

import "container/list"

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

// ConnectedComponents respresents connected components of a graph.
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

// ComponentID tells the component which vertex belongs to.
func (cc ConnectedComponents) ComponentID(vertex int) int {
	return cc.components[vertex]
}

func findConnectedComponentsCore(cc *ConnectedComponents, marked []bool, graph Graph, vertex int) {
	marked[vertex] = true
	cc.components[vertex] = cc.count
	for _, v := range graph.AdjacentVertices(vertex) {
		if !marked[v] {
			findConnectedComponentsCore(cc, marked, graph, v)
		}
	}
}

// FindConnectedComponents finds connected components in a graph.
func FindConnectedComponents(graph Graph) ConnectedComponents {
	count := graph.NumVertices()
	result := ConnectedComponents{
		components: make([]int, count),
	}
	marked := make([]bool, count)
	for v := 0; v < count; v++ {
		if !marked[v] {
			findConnectedComponentsCore(&result, marked, graph, v)
			result.count++
		}
	}
	return result
}

func hasCycleCore(marked []bool, graph Graph, initialVertex int, vertex int) bool {
	marked[vertex] = true
	for _, v := range graph.AdjacentVertices(vertex) {
		if !marked[v] {
			if hasCycleCore(marked, graph, vertex, v) {
				return true
			}
		} else if v != initialVertex {
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

func isBipartiteCore(marked []bool, colors []bool, graph Graph, vertex int) bool {
	marked[vertex] = true
	for _, v := range graph.AdjacentVertices(vertex) {
		if !marked[v] {
			colors[v] = !colors[vertex]
			if !isBipartiteCore(marked, colors, graph, v) {
				return false
			}
		} else if colors[v] == colors[vertex] {
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
