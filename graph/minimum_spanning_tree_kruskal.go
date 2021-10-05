package graph

import "container/heap"

type unionFind struct {
	parent []int
	rank   []byte
}

func (uf *unionFind) find(p int) int {
	i := p
	parent := uf.parent
	for i != parent[i] {
		i = parent[i]
	}
	return i
}

func (uf *unionFind) union(p, q int) {
	pRoot, qRoot := uf.find(p), uf.find(q)
	if pRoot == qRoot {
		return
	}
	pRank, qRank := uf.rank[pRoot], uf.rank[qRoot]
	if pRank < qRank {
		uf.parent[pRoot] = qRoot
	} else {
		uf.parent[qRoot] = pRoot
	}
	if pRank == qRank {
		uf.rank[pRoot]++
	}
}

func (uf *unionFind) connected(p, q int) bool {
	return uf.find(p) == uf.find(q)
}

func newUnionFind(size int) *unionFind {
	parent := make([]int, size)
	rank := make([]byte, size)
	for i := 0; i < size; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return &unionFind{
		parent: parent,
		rank:   rank,
	}
}

type edgesPQ struct {
	edges   []Edge
	weights []float64
}

func (pq *edgesPQ) Len() int {
	return len(pq.edges)
}
func (pq *edgesPQ) Less(i, j int) bool {
	return pq.weights[i] < pq.weights[j]
}
func (pq *edgesPQ) Swap(i, j int) {
	pq.edges[i], pq.edges[j] = pq.edges[j], pq.edges[i]
	pq.weights[i], pq.weights[j] = pq.weights[j], pq.weights[i]
}

func (pq *edgesPQ) Push(val interface{}) {
}
func (pq *edgesPQ) Pop() interface{} {
	return nil
}

func (pq *edgesPQ) push(edge Edge, weight float64) {
	pq.edges = append(pq.edges, edge)
	pq.weights = append(pq.weights, weight)
	heap.Push(pq, nil)
}
func (pq *edgesPQ) pop() (Edge, float64) {
	edge, weight := pq.edges[0], pq.weights[0]
	heap.Pop(pq)
	n := pq.Len() - 1
	pq.edges = pq.edges[:n]
	pq.weights = pq.weights[:n]
	return edge, weight
}

func newEdgesPQ() *edgesPQ {
	return &edgesPQ{}
}

// MinimumSpanningTreeKruskal computes minimum spanning tree using Kruskal's algorithm.
// https://algs4.cs.princeton.edu/43mst/KruskalMST.java.html
func MinimumSpanningTreeKruskal(graph EdgeWeightedGraph) EdgeWeightedGraph {
	pq := newEdgesPQ()
	allWeights := AllGraphWeights(graph)
	for i, edge := range AllGraphEdges(graph) {
		pq.push(edge, allWeights[i])
	}
	numVertices := graph.NumVertices()
	uf := newUnionFind(numVertices)
	adjacency := make([][]int, numVertices)
	weights := make([][]float64, numVertices)
	numEdges := 0
	for pq.Len() > 0 {
		edge, weight := pq.pop()
		v, w := edge.Vertex1(), edge.Vertex2()
		if !uf.connected(v, w) {
			uf.union(v, w)
			addWeightedEdge(adjacency, weights, v, w, weight)
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
