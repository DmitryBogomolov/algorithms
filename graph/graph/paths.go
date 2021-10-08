package graph

import (
	"container/list"
	"fmt"
)

// Paths is a collection of paths from the source vertex to other vertices.
// A path is a sequence of vertices connected by edges.
type Paths struct {
	sourceVertex int
	vertexCount  int
	edgeTo       []int
}

func initPaths(graph Graph, vertexID int) Paths {
	count := graph.NumVertices()
	edgeTo := make([]int, count)
	for i := range edgeTo {
		edgeTo[i] = -1
	}
	return Paths{sourceVertex: vertexID, edgeTo: edgeTo, vertexCount: 0}
}

// SourceVertex gets source vertex.
func (paths Paths) SourceVertex() int {
	return paths.sourceVertex
}

// VertexCount gets number of vertices connected with source vertex.
func (paths Paths) VertexCount() int {
	return paths.vertexCount
}

// HasPathTo tells if a vertex is connected with source vertex.
func (paths Paths) HasPathTo(vertexID int) bool {
	if vertexID < 0 || vertexID > len(paths.edgeTo)-1 {
		panic(fmt.Sprintf("vertex '%d' is out of range", vertexID))
	}
	return paths.edgeTo[vertexID] >= 0 || vertexID == paths.sourceVertex
}

// PathTo returns path from source vertex to a vertex.
func (paths Paths) PathTo(vertexID int) []int {
	if !paths.HasPathTo(vertexID) {
		return nil
	}
	var stack []int
	for currentVertexID := vertexID; currentVertexID >= 0; currentVertexID = paths.edgeTo[currentVertexID] {
		stack = append(stack, currentVertexID)
	}
	reverseList(stack)
	return stack
}

func reverseList(list []int) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}

type _SearchVisitor interface {
	searchVisit(vertexID int, parentVertexID int)
}

func searchDepthFirstCore(graph Graph, vertexID int, marked []bool, visitor _SearchVisitor) {
	marked[vertexID] = true
	for _, adjacentVertexID := range graph.AdjacentVertices(vertexID) {
		if !marked[adjacentVertexID] {
			visitor.searchVisit(adjacentVertexID, vertexID)
			searchDepthFirstCore(graph, adjacentVertexID, marked, visitor)
		}
	}
}

func searchDepthFirst(graph Graph, vertexID int, visitor _SearchVisitor) {
	marked := make([]bool, graph.NumVertices())
	searchDepthFirstCore(graph, vertexID, marked, visitor)
}

func (paths *Paths) searchVisit(vertexID int, parentVertexID int) {
	paths.vertexCount++
	paths.edgeTo[vertexID] = parentVertexID
}

// FindPathsDepthFirst returns paths from a vertex using depth-first search.
// https://algs4.cs.princeton.edu/41graph/DepthFirstSearch.java.html
func FindPathsDepthFirst(graph Graph, vertexID int) Paths {
	paths := initPaths(graph, vertexID)
	searchDepthFirst(graph, vertexID, &paths)
	return paths
}

func searchBreadthFirstCore(graph Graph, queue *list.List, marked []bool, visitor _SearchVisitor) {
	for queue.Len() > 0 {
		vertexID := queue.Front().Value.(int)
		queue.Remove(queue.Front())
		for _, adjacentVertexID := range graph.AdjacentVertices(vertexID) {
			if !marked[adjacentVertexID] {
				marked[adjacentVertexID] = true
				queue.PushBack(adjacentVertexID)
				visitor.searchVisit(adjacentVertexID, vertexID)
			}
		}
	}
}

func searchBreadthFirst(graph Graph, vertexID int, visitor _SearchVisitor) {
	marked := make([]bool, graph.NumVertices())
	queue := list.New()
	queue.PushBack(vertexID)
	marked[vertexID] = true
	searchBreadthFirstCore(graph, queue, marked, visitor)
}

// FindPathsBreadthFirst returns paths from a vertex using breadth-first search.
// https://algs4.cs.princeton.edu/41graph/BreadthFirstPaths.java.html
func FindPathsBreadthFirst(graph Graph, vertexID int) Paths {
	paths := initPaths(graph, vertexID)
	searchBreadthFirst(graph, vertexID, &paths)
	return paths
}
