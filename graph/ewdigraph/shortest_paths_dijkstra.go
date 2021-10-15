package ewdigraph

import (
	"algorithms/graph/ewgraph"
	"algorithms/graph/internals/utils"
	"algorithms/indexpriorityqueue"
	"fmt"
	"math"
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

// FindShortedPathsDijkstra returns shortest paths from a vertex.
// https://algs4.cs.princeton.edu/44sp/DijkstraSP.java.html
func FindShortedPathsDijkstra(ewdigraph ewgraph.EdgeWeightedGraph, vertexID int) ShortestPaths {
	if vertexID < 0 || vertexID > ewdigraph.NumVertices()-1 {
		panic(fmt.Sprintf("vertex '%d' is out of range", vertexID))
	}
	numVertices := ewdigraph.NumVertices()
	edgeTo := make([]int, numVertices)
	distTo := make([]float64, numVertices)
	utils.ResetList(edgeTo)
	for i := range distTo {
		distTo[i] = math.MaxFloat64
	}
	distTo[vertexID] = 0.0
	verticesPriorityQueue := indexpriorityqueue.New(func(lhs, rhs interface{}) bool {
		return lhs.(float64) < rhs.(float64)
	})
	verticesPriorityQueue.Insert(vertexID, distTo[vertexID])
	relaxVerticesDijkstra(ewdigraph, verticesPriorityQueue, distTo, edgeTo)
	vertexCount := 0
	for _, edgeID := range edgeTo {
		if edgeID >= 0 {
			vertexCount++
		}
	}
	return ShortestPaths{
		sourceVertex: vertexID,
		vertexCount:  vertexCount,
		edgeTo:       edgeTo,
		distTo:       distTo,
	}
}

func relaxVerticesDijkstra(
	ewdigraph ewgraph.EdgeWeightedGraph,
	verticesPriorityQueue indexpriorityqueue.IndexPriorityQueue,
	distTo []float64, edgeTo []int,
) {
	for verticesPriorityQueue.Size() > 0 {
		_, currentVertexID := verticesPriorityQueue.Remove()
		weights := ewdigraph.AdjacentWeights(currentVertexID)
		for i, adjacentVertexID := range ewdigraph.AdjacentVertices(currentVertexID) {
			relaxEdgeDijkstra(ewdigraph, verticesPriorityQueue, distTo, edgeTo, currentVertexID, adjacentVertexID, weights[i])
		}
	}
}

func relaxEdgeDijkstra(
	ewdigraph ewgraph.EdgeWeightedGraph,
	verticesPriorityQueue indexpriorityqueue.IndexPriorityQueue,
	distTo []float64, edgeTo []int,
	fromVertexID int, toVertexID int, weight float64,
) {
	if distTo[toVertexID] > distTo[fromVertexID]+weight {
		distTo[toVertexID] = distTo[fromVertexID] + weight
		edgeTo[toVertexID] = fromVertexID
		verticesPriorityQueue.Insert(toVertexID, distTo[toVertexID])
	}
}
