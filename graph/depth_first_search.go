package graph

// DepthFirstSearchResult represents result of depth-first search.
type DepthFirstSearchResult struct {
	source int
	count  int
	marked []bool
	edgeTo []int
}

// Source show source vertex.
func (r DepthFirstSearchResult) Source() int {
	return r.source
}

// Marked shows if *vertex* connected with source vertex.
func (r DepthFirstSearchResult) Marked(vertex int) bool {
	return r.marked[vertex]
}

// Count shows number of vertices connected with source vertex.
func (r DepthFirstSearchResult) Count() int {
	return r.count
}

// PathTo shows a path from source vertex to *vertex*.
func (r DepthFirstSearchResult) PathTo(vertex int) []int {
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

func depthFirstSearch(r *DepthFirstSearchResult, graph Graph, vertex int) {
	r.marked[vertex] = true
	r.count++
	for _, v := range graph.AdjacentVertices(vertex) {
		if !r.marked[v] {
			r.edgeTo[v] = vertex
			depthFirstSearch(r, graph, v)
		}
	}
}

// DepthFirstSearch performs depth-first search for a specified vertex.
func DepthFirstSearch(graph Graph, vertex int) DepthFirstSearchResult {
	result := DepthFirstSearchResult{
		source: vertex,
		marked: make([]bool, graph.NumVertices()),
		edgeTo: make([]int, graph.NumVertices()),
	}
	depthFirstSearch(&result, graph, vertex)
	return result
}
