package ewdigraph

import (
	"container/list"
	"fmt"
	"math"

	"github.com/DmitryBogomolov/algorithms/graph/digraph"
	"github.com/DmitryBogomolov/algorithms/graph/ewgraph"
	"github.com/DmitryBogomolov/algorithms/graph/graph"
	"github.com/DmitryBogomolov/algorithms/graph/internal/utils"
)

// FindShortedPathsBellmanFord returns shortest paths from a vertex.
// https://algs4.cs.princeton.edu/44sp/BellmanFordSP.java.html
func FindShortedPathsBellmanFord(wdgr ewgraph.EdgeWeightedGraph, vertexID int) (ShortestPaths, []int) {
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
	onQueue := make([]bool, numVertices)
	cost := 0
	var negativeCycle []int
	verticesQueue := list.New()
	verticesQueue.PushBack(vertexID)
	onQueue[vertexID] = true
	for verticesQueue.Len() > 0 && negativeCycle == nil {
		currentVertexID := verticesQueue.Front().Value.(int)
		verticesQueue.Remove(verticesQueue.Front())
		onQueue[currentVertexID] = false
		relaxVertexBellmanFord(wdgr, edgeTo, distTo, onQueue, verticesQueue, &cost, &negativeCycle, currentVertexID)
	}

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
	}, negativeCycle
}

func relaxVertexBellmanFord(
	wdgr ewgraph.EdgeWeightedGraph, edgeTo []int, distTo []float64, onQueue []bool, verticesQueue *list.List,
	cost *int, negativeCycle *[]int,
	vertexID int,
) {
	weights := wdgr.AdjacentWeights(vertexID)
	for i, adjacentVertexID := range wdgr.AdjacentVertices(vertexID) {
		if distTo[adjacentVertexID] > distTo[vertexID]+weights[i] {
			distTo[adjacentVertexID] = distTo[vertexID] + weights[i]
			edgeTo[adjacentVertexID] = vertexID
			if !onQueue[adjacentVertexID] {
				verticesQueue.PushBack(adjacentVertexID)
				onQueue[adjacentVertexID] = true
			}
			*cost++
			if *cost%wdgr.NumVertices() == 0 {
				*negativeCycle = findNegativeCycle(edgeTo)
				if *negativeCycle != nil {
					return
				}
			}
		}
	}
}

func findNegativeCycle(edgeTo []int) []int {
	numVertices := len(edgeTo)
	numEdges := 0
	adjacency := make([][]int, numVertices)
	for vertexID := 0; vertexID < numVertices; vertexID++ {
		if fromVertexID := edgeTo[vertexID]; fromVertexID != -1 {
			adjacency[fromVertexID] = append(adjacency[fromVertexID], vertexID)
			numEdges++
		}
	}
	dgr := graph.NewImplGraph(numVertices, numEdges, adjacency)
	return digraph.FindDirectedCycle(dgr)
}
