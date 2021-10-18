package graph

import (
	"container/list"
	"fmt"

	"github.com/DmitryBogomolov/algorithms/graph/internals/utils"
)

// Paths is a collection of paths from the source vertex to other vertices.
// A path is a sequence of vertices connected by edges.
type Paths struct {
	sourceVertex int
	vertexCount  int
	edgeTo       []int
}

func initPaths(gr Graph, vertexID int) Paths {
	count := gr.NumVertices()
	edgeTo := make([]int, count)
	utils.ResetList(edgeTo)
	return Paths{
		sourceVertex: vertexID,
		vertexCount:  0,
		edgeTo:       edgeTo,
	}
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
	utils.ReverseList(stack)
	return stack
}

type _SearchPathsVisitor interface {
	searchPathsVisit(vertexID int, parentVertexID int)
}

func searchPathsDepthFirstCore(gr Graph, marked []bool, visitor _SearchPathsVisitor, vertexID int) {
	marked[vertexID] = true
	for _, adjacentVertexID := range gr.AdjacentVertices(vertexID) {
		if !marked[adjacentVertexID] {
			visitor.searchPathsVisit(adjacentVertexID, vertexID)
			searchPathsDepthFirstCore(gr, marked, visitor, adjacentVertexID)
		}
	}
}

func searchPathsDepthFirst(gr Graph, visitor _SearchPathsVisitor, vertexID int) {
	marked := make([]bool, gr.NumVertices())
	searchPathsDepthFirstCore(gr, marked, visitor, vertexID)
}

func (paths *Paths) searchPathsVisit(vertexID int, parentVertexID int) {
	paths.vertexCount++
	paths.edgeTo[vertexID] = parentVertexID
}

// FindPathsDepthFirst returns paths from a vertex using depth-first search.
// https://algs4.cs.princeton.edu/41graph/DepthFirstSearch.java.html
func FindPathsDepthFirst(gr Graph, vertexID int) Paths {
	paths := initPaths(gr, vertexID)
	searchPathsDepthFirst(gr, &paths, vertexID)
	return paths
}

func searchPathsBreadthFirstCore(gr Graph, marked []bool, queue *list.List, visitor _SearchPathsVisitor) {
	for queue.Len() > 0 {
		vertexID := queue.Front().Value.(int)
		queue.Remove(queue.Front())
		for _, adjacentVertexID := range gr.AdjacentVertices(vertexID) {
			if !marked[adjacentVertexID] {
				marked[adjacentVertexID] = true
				queue.PushBack(adjacentVertexID)
				visitor.searchPathsVisit(adjacentVertexID, vertexID)
			}
		}
	}
}

func searchPathsBreadthFirst(gr Graph, visitor _SearchPathsVisitor, vertexID int) {
	marked := make([]bool, gr.NumVertices())
	queue := list.New()
	queue.PushBack(vertexID)
	marked[vertexID] = true
	searchPathsBreadthFirstCore(gr, marked, queue, visitor)
}

// FindPathsBreadthFirst returns paths from a vertex using breadth-first search.
// https://algs4.cs.princeton.edu/41graph/BreadthFirstPaths.java.html
func FindPathsBreadthFirst(gr Graph, vertexID int) Paths {
	paths := initPaths(gr, vertexID)
	searchPathsBreadthFirst(gr, &paths, vertexID)
	return paths
}
