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
func (g *TestEdgeWeightedGraph) AddEdge(v1, v2 int, w float64) {
	g.TestGraph.AddEdge(v1, v2)
	g.Weights[v1] = append(g.Weights[v1], w)
	g.Weights[v2] = append(g.Weights[v2], w)
}

// AddDirectedEdge TEST
func (g *TestEdgeWeightedGraph) AddDirectedEdge(v1, v2 int, w float64) {
	g.TestGraph.AddDirectedEdge(v1, v2)
	g.Weights[v1] = append(g.Weights[v1], w)
}

// AdjacentWeights TEST
func (g *TestEdgeWeightedGraph) AdjacentWeights(vertex int) []float64 {
	return g.Weights[vertex]
}

// NewTestEdgeWeightedGraph TEST
func NewTestEdgeWeightedGraph(numVertices int, edges []TestWeightedEdge) *TestEdgeWeightedGraph {
	graph := TestEdgeWeightedGraph{
		TestGraph: TestGraph{
			numVertices: numVertices,
			numEdges:    len(edges),
			adjacency:   make([][]int, numVertices),
		},
		Weights: make([][]float64, numVertices),
	}
	for _, edge := range edges {
		graph.AddEdge(edge.V1, edge.V2, edge.Weight)
	}
	return &graph
}

// NewTestEdgeWeightedDigraph TEST
func NewTestEdgeWeightedDigraph(numVertices int, edges []TestWeightedEdge) *TestEdgeWeightedGraph {
	digraph := TestEdgeWeightedGraph{
		TestGraph: TestGraph{
			numVertices: numVertices,
			numEdges:    len(edges),
			adjacency:   make([][]int, numVertices),
		},
		Weights: make([][]float64, numVertices),
	}
	for _, edge := range edges {
		digraph.AddDirectedEdge(edge.V1, edge.V2, edge.Weight)
	}
	return &digraph
}
