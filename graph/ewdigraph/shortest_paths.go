package ewdigraph

import (
	"algorithms/graph/internals/utils"
	"fmt"
)

// ShortestPaths is a collection of paths from the source vertex to other vertices.
type ShortestPaths struct {
	sourceVertex int
	vertexCount  int
	edgeTo       []int
	distTo       []float64
}

// SourceVertex gets source vertex.
func (paths ShortestPaths) SourceVertex() int {
	return paths.sourceVertex
}

// VertexCount gets number of vertices connected with source vertex.
func (paths ShortestPaths) VertexCount() int {
	return paths.vertexCount
}

// HasPathTo tells if a vertex is connected with source vertex.
func (paths ShortestPaths) HasPathTo(vertexID int) bool {
	if vertexID < 0 || vertexID > len(paths.edgeTo)-1 {
		panic(fmt.Sprintf("vertex '%d' is out of range", vertexID))
	}
	return paths.edgeTo[vertexID] >= 0 || vertexID == paths.sourceVertex
}

// PathTo returns path from source vertex to a vertex.
func (paths ShortestPaths) PathTo(vertexID int) ([]int, float64) {
	if !paths.HasPathTo(vertexID) {
		return nil, 0
	}
	var stack []int
	for currentVertexID := vertexID; currentVertexID >= 0; currentVertexID = paths.edgeTo[currentVertexID] {
		stack = append(stack, currentVertexID)
	}
	utils.ReverseList(stack)
	return stack, paths.distTo[vertexID]
}
