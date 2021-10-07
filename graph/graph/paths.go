package graph

import (
	"container/list"
)

// Paths is a collection of paths from the selected vertex to other vertices.
// A path is a sequence of vertices connected by edges.
type Paths struct {
	origin int
	count  int
	edgeTo []int
}

func initGraphPaths(graph Graph, vertexID int) (Paths, []bool) {
	count := graph.NumVertices()
	edgeTo := make([]int, count)
	marked := make([]bool, count)
	for i := range edgeTo {
		edgeTo[i] = -1
	}
	return Paths{origin: vertexID, edgeTo: edgeTo}, marked
}

// Origin gets initial vertex.
func (paths Paths) Origin() int {
	return paths.origin
}

// Count gets number of vertices connected with initial vertex.
func (paths Paths) Count() int {
	return paths.count
}

// HasPathTo tells if a vertex is connected with initial vertex.
func (paths Paths) HasPathTo(targetVertexID int) bool {
	return paths.edgeTo[targetVertexID] >= 0 || targetVertexID == paths.origin
}

// PathTo returns a path from initial vertex to a vertex.
func (paths Paths) PathTo(targetVertexID int) []int {
	if !paths.HasPathTo(targetVertexID) {
		return nil
	}
	var stack []int
	for vertexID := targetVertexID; vertexID >= 0; vertexID = paths.edgeTo[vertexID] {
		stack = append(stack, vertexID)
	}
	reverseList(stack)
	return stack
}

func reverseList(list []int) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}

func findPathsDepthFirstCore(paths *Paths, marked []bool, graph Graph, vertexID int) {
	marked[vertexID] = true
	paths.count++
	for _, adjacentVertexID := range graph.AdjacentVertices(vertexID) {
		if !marked[adjacentVertexID] {
			paths.edgeTo[adjacentVertexID] = vertexID
			findPathsDepthFirstCore(paths, marked, graph, adjacentVertexID)
		}
	}
}

// FindPathsDepthFirst finds paths from a vertex using depth-first search.
// https://algs4.cs.princeton.edu/41graph/DepthFirstSearch.java.html
func FindPathsDepthFirst(graph Graph, vertexID int) Paths {
	paths, marked := initGraphPaths(graph, vertexID)
	findPathsDepthFirstCore(&paths, marked, graph, vertexID)
	return paths
}

func findPathsBreadthFirstCore(paths *Paths, marked []bool, graph Graph, vertexID int) {
	queue := list.New()
	queue.PushBack(vertexID)
	marked[vertexID] = true
	paths.count++
	for queue.Len() > 0 {
		currentVertexID := queue.Front().Value.(int)
		queue.Remove(queue.Front())
		for _, adjacentVertexID := range graph.AdjacentVertices(currentVertexID) {
			if !marked[adjacentVertexID] {
				marked[adjacentVertexID] = true
				paths.count++
				paths.edgeTo[adjacentVertexID] = currentVertexID
				queue.PushBack(adjacentVertexID)
			}
		}
	}
}

// FindPathsBreadthFirst finds paths from a vertex using breadth-first search.
// https://algs4.cs.princeton.edu/41graph/BreadthFirstPaths.java.html
func FindPathsBreadthFirst(graph Graph, vertexID int) Paths {
	result, marked := initGraphPaths(graph, vertexID)
	findPathsBreadthFirstCore(&result, marked, graph, vertexID)
	return result
}
