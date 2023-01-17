package ewdigraph

import (
	"fmt"
	"math"

	"github.com/DmitryBogomolov/algorithms/graph/ewgraph"
	"github.com/DmitryBogomolov/algorithms/graph/internal/utils"
	ipq "github.com/DmitryBogomolov/algorithms/indexpriorityqueue"
)

// FindShortedPathsDijkstra returns shortest paths from a vertex.
// https://algs4.cs.princeton.edu/44sp/DijkstraSP.java.html
func FindShortedPathsDijkstra(wdgr ewgraph.EdgeWeightedGraph, vertexID int) ShortestPaths {
	if vertexID < 0 || vertexID > wdgr.NumVertices()-1 {
		panic(fmt.Sprintf("vertex '%d' is out of range", vertexID))
	}
	numVertices := wdgr.NumVertices()
	edgeTo := make([]int, numVertices)
	utils.ResetList(edgeTo)
	distTo := make([]float64, numVertices)
	for i := range distTo {
		distTo[i] = math.MaxFloat64
	}
	distTo[vertexID] = 0.0
	verticesQueue := ipq.New(func(lhs, rhs interface{}) bool {
		return lhs.(float64) < rhs.(float64)
	})
	verticesQueue.Insert(vertexID, distTo[vertexID])
	relaxVerticesDijkstra(wdgr, edgeTo, distTo, verticesQueue)
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
	wdgr ewgraph.EdgeWeightedGraph, edgeTo []int, distTo []float64, verticesQueue ipq.IndexPriorityQueue[any],
) {
	for verticesQueue.Size() > 0 {
		_, currentVertexID := verticesQueue.Remove()
		weights := wdgr.AdjacentWeights(currentVertexID)
		for i, adjacentVertexID := range wdgr.AdjacentVertices(currentVertexID) {
			relaxEdgeDijkstra(wdgr, edgeTo, distTo, verticesQueue, currentVertexID, adjacentVertexID, weights[i])
		}
	}
}

func relaxEdgeDijkstra(
	wdgr ewgraph.EdgeWeightedGraph, edgeTo []int, distTo []float64, verticesQueue ipq.IndexPriorityQueue[any],
	fromVertexID int, toVertexID int, weight float64,
) {
	if distTo[toVertexID] > distTo[fromVertexID]+weight {
		distTo[toVertexID] = distTo[fromVertexID] + weight
		edgeTo[toVertexID] = fromVertexID
		verticesQueue.Insert(toVertexID, distTo[toVertexID])
	}
}
