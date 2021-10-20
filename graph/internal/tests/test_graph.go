package tests

// TestGraph TEST
type TestGraph struct {
	numVertices int
	numEdges    int
	adjacency   [][]int
}

// NewTestGraph TEST
func NewTestGraph(numVertices int, connections ...int) *TestGraph {
	gr := TestGraph{
		numVertices: numVertices,
		numEdges:    len(connections) / 2,
		adjacency:   make([][]int, numVertices),
	}
	for i := 0; i < len(connections)/2; i++ {
		gr.AddEdge(connections[2*i], connections[2*i+1])
	}
	return &gr
}

// NewTestDigraph TEST
func NewTestDigraph(numVertices int, connections ...int) *TestGraph {
	dgr := TestGraph{
		numVertices: numVertices,
		numEdges:    len(connections) / 2,
		adjacency:   make([][]int, numVertices),
	}
	for i := 0; i < len(connections)/2; i++ {
		dgr.AddDirectedEdge(connections[2*i], connections[2*i+1])
	}
	return &dgr
}

// AddEdge TEST
func (gr *TestGraph) AddEdge(vertexID1, vertexID2 int) {
	gr.adjacency[vertexID1] = append(gr.adjacency[vertexID1], vertexID2)
	gr.adjacency[vertexID2] = append(gr.adjacency[vertexID2], vertexID1)
}

// AddDirectedEdge TEST
func (gr *TestGraph) AddDirectedEdge(vertexID1, vertexID2 int) {
	gr.adjacency[vertexID1] = append(gr.adjacency[vertexID1], vertexID2)
}

// NumVertices TEST
func (gr *TestGraph) NumVertices() int {
	return gr.numVertices
}

// NumEdges TEST
func (gr *TestGraph) NumEdges() int {
	return gr.numEdges
}

// AdjacentVertices TEST
func (gr *TestGraph) AdjacentVertices(vertexID int) []int {
	return gr.adjacency[vertexID]
}
