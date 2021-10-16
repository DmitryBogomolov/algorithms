package ewdigraph

import (
	"algorithms/graph/ewgraph"
	"algorithms/graph/internals/utils"
	"algorithms/indexpriorityqueue"
	"fmt"
	"math"
)

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
