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
	for i, j := 0, len(stack)-1; i < j; i, j = i+1, j-1 {
		stack[i], stack[j] = stack[j], stack[i]
	}
	return stack
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
