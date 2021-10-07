package internals

type TestGraph struct {
	numVertices int
	numEdges    int
	adjacency   [][]int
}

func NewTestGraph(numVertices int, connections ...int) *TestGraph {
	graph := TestGraph{
		numVertices: numVertices,
		numEdges:    len(connections) / 2,
		adjacency:   make([][]int, numVertices),
	}
	for i := 0; i < len(connections)/2; i++ {
		graph.AddEdge(connections[2*i], connections[2*i+1])
	}
	return &graph
}

func NewTestDigraph(numVertices int, connections ...int) *TestGraph {
	digraph := TestGraph{
		numVertices: numVertices,
		numEdges:    len(connections) / 2,
		adjacency:   make([][]int, numVertices),
	}
	for i := 0; i < len(connections)/2; i++ {
		digraph.AddDirectedEdge(connections[2*i], connections[2*i+1])
	}
	return &digraph
}

func (g *TestGraph) AddEdge(v1, v2 int) {
	g.adjacency[v1] = append(g.adjacency[v1], v2)
	g.adjacency[v2] = append(g.adjacency[v2], v1)
}

func (g *TestGraph) AddDirectedEdge(v1, v2 int) {
	g.adjacency[v1] = append(g.adjacency[v1], v2)
}

func (g *TestGraph) NumVertices() int {
	return g.numVertices
}

func (g *TestGraph) NumEdges() int {
	return g.numEdges
}

func (g *TestGraph) AdjacentVertices(vertex int) []int {
	return g.adjacency[vertex]
}
