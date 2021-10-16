package ewdigraph

import (
	"algorithms/graph/digraph"
	"algorithms/graph/ewgraph"
	"algorithms/graph/internals/utils"
	"container/list"
	"fmt"
	"math"
)

// FindShortedPathsBellmanFord returns shortest paths from a vertex.
// https://algs4.cs.princeton.edu/44sp/BellmanFordSP.java.html
func FindShortedPathsBellmanFord(ewdigraph ewgraph.EdgeWeightedGraph, vertexID int) (ShortestPaths, []int) {
	if vertexID < 0 || vertexID > ewdigraph.NumVertices()-1 {
		panic(fmt.Sprintf("vertex '%d' is out of range", vertexID))
	}
	numVertices := ewdigraph.NumVertices()
	edgeTo := make([]int, numVertices)
	distTo := make([]float64, numVertices)
	onQueue := make([]bool, numVertices)
	cost := 0
	var negativeCycle []int
	utils.ResetList(edgeTo)
	for i := range distTo {
		distTo[i] = math.MaxFloat64
	}
	distTo[vertexID] = 0.0

	verticesQueue := list.New()
	verticesQueue.PushBack(vertexID)
	onQueue[vertexID] = true
	for verticesQueue.Len() > 0 && negativeCycle == nil {
		currentVertexID := verticesQueue.Front().Value.(int)
		verticesQueue.Remove(verticesQueue.Front())
		onQueue[currentVertexID] = false
		relaxVertexBellmanFord(ewdigraph, currentVertexID, edgeTo, distTo, onQueue, verticesQueue, &cost, &negativeCycle)
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
	ewdigraph ewgraph.EdgeWeightedGraph, vertexID int,
	edgeTo []int, distTo []float64, onQueue []bool, verticesQueue *list.List,
	cost *int, negativeCycle *[]int,
) {
	weights := ewdigraph.AdjacentWeights(vertexID)
	for i, adjacentVertexID := range ewdigraph.AdjacentVertices(vertexID) {
		if distTo[adjacentVertexID] > distTo[vertexID]+weights[i] {
			distTo[adjacentVertexID] = distTo[vertexID] + weights[i]
			edgeTo[adjacentVertexID] = vertexID
			if !onQueue[adjacentVertexID] {
				verticesQueue.PushBack(adjacentVertexID)
				onQueue[adjacentVertexID] = true
			}
			*cost++
			if *cost%ewdigraph.NumVertices() == 0 {
				*negativeCycle = findNegativeCycle(edgeTo)
				if *negativeCycle != nil {
					return
				}
			}
		}
	}
}

type _TmpDigraph struct {
	numVertices int
	numEdges    int
	adjacency   [][]int
}

func (obj _TmpDigraph) NumVertices() int {
	return obj.numVertices
}

func (obj _TmpDigraph) NumEdges() int {
	return obj.numEdges
}

func (obj _TmpDigraph) AdjacentVertices(vertexID int) []int {
	return obj.adjacency[vertexID]
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
	tmpDigraph := _TmpDigraph{
		numVertices: numVertices,
		numEdges:    numEdges,
		adjacency:   adjacency,
	}
	return digraph.FindDirectedCycle(tmpDigraph)
}
