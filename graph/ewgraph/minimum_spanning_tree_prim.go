package ewgraph

import (
	"math"
)

type minimumSpanningTree struct {
	origin    EdgeWeightedGraph
	numEdges  int
	adjacency [][]int
	weights   [][]float64
}

func (t minimumSpanningTree) NumVertices() int {
	return t.origin.NumVertices()
}
func (t minimumSpanningTree) NumEdges() int {
	return t.numEdges
}
func (t minimumSpanningTree) AdjacentVertices(vertex int) []int {
	return t.adjacency[vertex]
}
func (t minimumSpanningTree) AdjacentWeights(vertex int) []float64 {
	return t.weights[vertex]
}

func addWeightedEdge(adjacency [][]int, weights [][]float64, v1, v2 int, weight float64) {
	adjacency[v1] = append(adjacency[v1], v2)
	adjacency[v2] = append(adjacency[v2], v1)
	weights[v1] = append(weights[v1], weight)
	weights[v2] = append(weights[v2], weight)
}

func scanMinimumSpanningTreeVertex(
	pq *_VerticesIndexPriorityQueue, marked []bool, edgeTo []int, distTo []float64,
	graph EdgeWeightedGraph, current int,
) {
	marked[current] = true
	weights := graph.AdjacentWeights(current)
	for i, v := range graph.AdjacentVertices(current) {
		weight := weights[i]
		if !marked[v] && weight < distTo[v] {
			edgeTo[v] = current
			distTo[v] = weight
			pq.updateVertex(v, weight)
		}
	}
}

func processMinimumSpanningTree(
	pq *_VerticesIndexPriorityQueue, marked []bool, edgeTo []int, distTo []float64,
	graph EdgeWeightedGraph, start int,
) {
	distTo[start] = 0
	pq.updateVertex(start, 0)
	for pq.Len() > 0 {
		v := pq.popVertex()
		scanMinimumSpanningTreeVertex(pq, marked, edgeTo, distTo, graph, v)
	}
}

// MinimumSpanningTreePrim computes minimum spanning tree using Prim's algorithm.
// https://algs4.cs.princeton.edu/43mst/PrimMST.java.html
func MinimumSpanningTreePrim(graph EdgeWeightedGraph) EdgeWeightedGraph {
	numVertices := graph.NumVertices()
	marked := make([]bool, numVertices)
	edgeTo := make([]int, numVertices)
	distTo := make([]float64, numVertices)
	pq := newVerticesIndexPriorityQueue(numVertices)
	for vertexID := 0; vertexID < numVertices; vertexID++ {
		edgeTo[vertexID] = -1
		distTo[vertexID] = math.MaxFloat64
	}
	for vertexID := 0; vertexID < numVertices; vertexID++ {
		if !marked[vertexID] {
			processMinimumSpanningTree(pq, marked, edgeTo, distTo, graph, vertexID)
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
	return minimumSpanningTree{
		origin:    graph,
		numEdges:  numEdges,
		adjacency: adjacency,
		weights:   weights,
	}
}
