package ewgraph

import (
	"math"
)

func scanMinimumSpanningTreeVertexPrim(
	verticesIndexPriorityQueue *_VerticesIndexPriorityQueue, marked []bool, edgeTo []int, distTo []float64,
	ewgraph EdgeWeightedGraph, vertexID int,
) {
	marked[vertexID] = true
	weights := ewgraph.AdjacentWeights(vertexID)
	for i, adjacentVertexID := range ewgraph.AdjacentVertices(vertexID) {
		weight := weights[i]
		if !marked[adjacentVertexID] && weight < distTo[adjacentVertexID] {
			edgeTo[adjacentVertexID] = vertexID
			distTo[adjacentVertexID] = weight
			verticesIndexPriorityQueue.updateVertex(adjacentVertexID, weight)
		}
	}
}

func processMinimumSpanningTreePrim(
	verticesIndexPriorityQueue *_VerticesIndexPriorityQueue, marked []bool, edgeTo []int, distTo []float64,
	ewgraph EdgeWeightedGraph, startVertexID int,
) {
	distTo[startVertexID] = 0
	verticesIndexPriorityQueue.updateVertex(startVertexID, 0)
	for verticesIndexPriorityQueue.Len() > 0 {
		vertexID := verticesIndexPriorityQueue.popVertex()
		scanMinimumSpanningTreeVertexPrim(verticesIndexPriorityQueue, marked, edgeTo, distTo, ewgraph, vertexID)
	}
}

// BuildMinimumSpanningTreePrim computes minimum spanning tree using Prim's algorithm.
// https://algs4.cs.princeton.edu/43mst/PrimMST.java.html
func BuildMinimumSpanningTreePrim(ewgraph EdgeWeightedGraph) EdgeWeightedGraph {
	numVertices := ewgraph.NumVertices()
	marked := make([]bool, numVertices)
	edgeTo := make([]int, numVertices)
	distTo := make([]float64, numVertices)
	verticesIndexPriorityQueue := newVerticesIndexPriorityQueue(numVertices)
	for vertexID := 0; vertexID < numVertices; vertexID++ {
		edgeTo[vertexID] = -1
		distTo[vertexID] = math.MaxFloat64
	}
	for vertexID := 0; vertexID < numVertices; vertexID++ {
		if !marked[vertexID] {
			processMinimumSpanningTreePrim(verticesIndexPriorityQueue, marked, edgeTo, distTo, ewgraph, vertexID)
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
	return _MinimumSpanningTree{
		numVertices: numVertices,
		numEdges:    numEdges,
		adjacency:   adjacency,
		weights:     weights,
	}
}
