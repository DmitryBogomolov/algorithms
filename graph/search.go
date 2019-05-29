package graph

import "container/list"

// VertexPaths represents paths from a vertex.
type VertexPaths struct {
	source int
	count  int
	marked []bool
	edgeTo []int
}

func newVertexPaths(graph Graph, vertex int) VertexPaths {
	count := graph.NumVertices()
	return VertexPaths{
		source: vertex,
		marked: make([]bool, count),
		edgeTo: make([]int, count),
	}
}

// Origin show initial vertex.
func (r VertexPaths) Origin() int {
	return r.source
}

// HasPathTo shows if *vertex* is connected with initial vertex.
func (r VertexPaths) HasPathTo(vertex int) bool {
	return r.marked[vertex]
}

// Count shows number of vertices connected with initial vertex.
func (r VertexPaths) Count() int {
	return r.count
}

// PathTo shows a path from source vertex to *vertex*.
func (r VertexPaths) PathTo(vertex int) []int {
	if !r.marked[vertex] {
		return nil
	}
	var stack []int
	for v := vertex; v != r.source; v = r.edgeTo[v] {
		stack = append(stack, v)
	}
	count := len(stack)
	path := make([]int, count+1)
	for i, v := range stack {
		path[count-i] = v
	}
	path[0] = r.source
	return path
}

func findPathsDepthFirstCore(r *VertexPaths, graph Graph, vertex int) {
	r.marked[vertex] = true
	r.count++
	for _, v := range graph.AdjacentVertices(vertex) {
		if !r.marked[v] {
			r.edgeTo[v] = vertex
			findPathsDepthFirstCore(r, graph, v)
		}
	}
}

// FindPathsDepthFirst finds paths from "vertex" using depth-first search.
func FindPathsDepthFirst(graph Graph, vertex int) VertexPaths {
	result := newVertexPaths(graph, vertex)
	findPathsDepthFirstCore(&result, graph, vertex)
	return result
}

func findPathsBreadthFirstCore(r *VertexPaths, graph Graph, vertex int) {
	queue := list.New()
	queue.PushBack(vertex)
	r.marked[vertex] = true
	r.count++
	for queue.Len() > 0 {
		current := queue.Front().Value.(int)
		queue.Remove(queue.Front())
		for _, v := range graph.AdjacentVertices(current) {
			if !r.marked[v] {
				r.marked[v] = true
				r.count++
				r.edgeTo[v] = current
				queue.PushBack(v)
			}
		}
	}
}

// FindPathsBreadthFirst finds paths from "vertex" using breadth-first search.
func FindPathsBreadthFirst(graph Graph, vertex int) VertexPaths {
	result := newVertexPaths(graph, vertex)
	findPathsBreadthFirstCore(&result, graph, vertex)
	return result
}

// ConnectedComponents respresents connected components of a graph.
type ConnectedComponents struct {
	marked     []bool
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

func findConnectedComponentsCore(cc *ConnectedComponents, graph Graph, vertex int) {
	cc.marked[vertex] = true
	cc.components[vertex] = cc.count
	for _, v := range graph.AdjacentVertices(vertex) {
		if !cc.marked[v] {
			findConnectedComponentsCore(cc, graph, v)
		}
	}
}

// FindConnectedComponents finds connected components in a graph.
func FindConnectedComponents(graph Graph) ConnectedComponents {
	count := graph.NumVertices()
	result := ConnectedComponents{
		marked:     make([]bool, count),
		components: make([]int, count),
	}
	for v := 0; v < count; v++ {
		if !result.marked[v] {
			findConnectedComponentsCore(&result, graph, v)
			result.count++
		}
	}
	return result
}
