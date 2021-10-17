package ewgraph

import (
	"algorithms/graph/graph"
	pq "algorithms/priorityqueue"
	"algorithms/unionfind"
)

type _EdgesQueueItem struct {
	edge   graph.Edge
	weight float64
}

// BuildMinimumSpanningTreeKruskal computes minimum spanning tree using Kruskal's algorithm.
// https://algs4.cs.princeton.edu/43mst/KruskalMST.java.html
func BuildMinimumSpanningTreeKruskal(wgr EdgeWeightedGraph) EdgeWeightedGraph {
	edgesQueue := pq.New(func(lhs, rhs interface{}) bool {
		return lhs.(_EdgesQueueItem).weight < rhs.(_EdgesQueueItem).weight
	})
	allWeights := AllGraphWeights(wgr)
	for i, edge := range graph.AllGraphEdges(wgr) {
		edgesQueue.Insert(_EdgesQueueItem{edge, allWeights[i]})
	}
	numVertices := wgr.NumVertices()
	uf := unionfind.New(numVertices)
	adjacency := make([][]int, numVertices)
	weights := make([][]float64, numVertices)
	numEdges := 0
	for edgesQueue.Size() > 0 {
		queueItem := edgesQueue.Remove().(_EdgesQueueItem)
		vertexID1, vertexID2 := queueItem.edge.Vertex1(), queueItem.edge.Vertex2()
		if !uf.Connected(vertexID1, vertexID2) {
			uf.Union(vertexID1, vertexID2)
			addWeightedEdge(adjacency, weights, vertexID1, vertexID2, queueItem.weight)
			numEdges++
		}
	}
	return NewImplEdgeWeightedGraph(numVertices, numEdges, adjacency, weights)
}
