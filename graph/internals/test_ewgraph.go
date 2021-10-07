package internals

type TestEdgeWeightedGraph struct {
	TestGraph
	Weights [][]float64
}

type TestWeightedEdge struct {
	V1, V2 int
	Weight float64
}

func (g *TestEdgeWeightedGraph) AddEdge(v1, v2 int, w float64) {
	g.TestGraph.AddEdge(v1, v2)
	g.Weights[v1] = append(g.Weights[v1], w)
	g.Weights[v2] = append(g.Weights[v2], w)
}

func (g *TestEdgeWeightedGraph) AddDirectedEdge(v1, v2 int, w float64) {
	g.TestGraph.AddDirectedEdge(v1, v2)
	g.Weights[v1] = append(g.Weights[v1], w)
}

func (g *TestEdgeWeightedGraph) AdjacentWeights(vertex int) []float64 {
	return g.Weights[vertex]
}

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
