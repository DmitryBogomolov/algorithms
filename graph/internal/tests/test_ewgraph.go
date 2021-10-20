package tests

// TestEdgeWeightedGraph TEST
type TestEdgeWeightedGraph struct {
	TestGraph
	Weights [][]float64
}

// TestWeightedEdge TEST
type TestWeightedEdge struct {
	V1, V2 int
	Weight float64
}

// AddEdge TEST
func (gr *TestEdgeWeightedGraph) AddEdge(v1, v2 int, w float64) {
	gr.TestGraph.AddEdge(v1, v2)
	gr.Weights[v1] = append(gr.Weights[v1], w)
	gr.Weights[v2] = append(gr.Weights[v2], w)
}

// AddDirectedEdge TEST
func (gr *TestEdgeWeightedGraph) AddDirectedEdge(v1, v2 int, w float64) {
	gr.TestGraph.AddDirectedEdge(v1, v2)
	gr.Weights[v1] = append(gr.Weights[v1], w)
}

// AdjacentWeights TEST
func (gr *TestEdgeWeightedGraph) AdjacentWeights(vertex int) []float64 {
	return gr.Weights[vertex]
}

// NewTestEdgeWeightedGraph TEST
func NewTestEdgeWeightedGraph(numVertices int, edges []TestWeightedEdge) *TestEdgeWeightedGraph {
	wgr := TestEdgeWeightedGraph{
		TestGraph: TestGraph{
			numVertices: numVertices,
			numEdges:    len(edges),
			adjacency:   make([][]int, numVertices),
		},
		Weights: make([][]float64, numVertices),
	}
	for _, edge := range edges {
		wgr.AddEdge(edge.V1, edge.V2, edge.Weight)
	}
	return &wgr
}

// NewTestEdgeWeightedDigraph TEST
func NewTestEdgeWeightedDigraph(numVertices int, edges []TestWeightedEdge) *TestEdgeWeightedGraph {
	wdgr := TestEdgeWeightedGraph{
		TestGraph: TestGraph{
			numVertices: numVertices,
			numEdges:    len(edges),
			adjacency:   make([][]int, numVertices),
		},
		Weights: make([][]float64, numVertices),
	}
	for _, edge := range edges {
		wdgr.AddDirectedEdge(edge.V1, edge.V2, edge.Weight)
	}
	return &wdgr
}
