package ewgraph

import (
	ipq "algorithms/indexpriorityqueue"
	"math"
)

func scanMinimumSpanningTreeVertexPrim(
	wgr EdgeWeightedGraph, marked []bool, edgeTo []int, distTo []float64, verticesQueue ipq.IndexPriorityQueue,
	vertexID int,
) {
	marked[vertexID] = true
	weights := wgr.AdjacentWeights(vertexID)
	for i, adjacentVertexID := range wgr.AdjacentVertices(vertexID) {
		weight := weights[i]
		if !marked[adjacentVertexID] && weight < distTo[adjacentVertexID] {
			edgeTo[adjacentVertexID] = vertexID
			distTo[adjacentVertexID] = weight
			verticesQueue.Insert(adjacentVertexID, weight)
		}
	}
}

func processMinimumSpanningTreePrim(
	wgr EdgeWeightedGraph, marked []bool, edgeTo []int, distTo []float64, verticesQueue ipq.IndexPriorityQueue,
	startVertexID int,
) {
	distTo[startVertexID] = 0
	verticesQueue.Insert(startVertexID, 0)
	for verticesQueue.Size() > 0 {
		_, vertexID := verticesQueue.Remove()
		scanMinimumSpanningTreeVertexPrim(wgr, marked, edgeTo, distTo, verticesQueue, vertexID)
	}
}

// BuildMinimumSpanningTreePrim computes minimum spanning tree using Prim's algorithm.
// https://algs4.cs.princeton.edu/43mst/PrimMST.java.html
func BuildMinimumSpanningTreePrim(wgr EdgeWeightedGraph) EdgeWeightedGraph {
	numVertices := wgr.NumVertices()
	marked := make([]bool, numVertices)
	edgeTo := make([]int, numVertices)
	distTo := make([]float64, numVertices)
	verticesQueue := ipq.New(func(lhs, rhs interface{}) bool {
		return lhs.(float64) < rhs.(float64)
	})
	for vertexID := 0; vertexID < numVertices; vertexID++ {
		edgeTo[vertexID] = -1
		distTo[vertexID] = math.MaxFloat64
	}
	for vertexID := 0; vertexID < numVertices; vertexID++ {
		if !marked[vertexID] {
			processMinimumSpanningTreePrim(wgr, marked, edgeTo, distTo, verticesQueue, vertexID)
		}
	}
	adjacency := make([][]int, numVertices)
	weights := make([][]float64, numVertices)
	numEdges := 0
	for vertexID := 0; vertexID < numVertices; vertexID++ {
		otherVertexID := edgeTo[vertexID]
		if otherVertexID != -1 {
			addWeightedEdge(adjacency, weights, vertexID, otherVertexID, distTo[vertexID])
			numEdges++
		}
	}
	return NewImplEdgeWeightedGraph(numVertices, numEdges, adjacency, weights)
}
