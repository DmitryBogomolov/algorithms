package ewgraph

type _MinimumSpanningTree struct {
	numVertices int
	numEdges    int
	adjacency   [][]int
	weights     [][]float64
}

func (msTree _MinimumSpanningTree) NumVertices() int {
	return msTree.numVertices
}
func (msTree _MinimumSpanningTree) NumEdges() int {
	return msTree.numEdges
}
func (msTree _MinimumSpanningTree) AdjacentVertices(vertex int) []int {
	return msTree.adjacency[vertex]
}
func (msTree _MinimumSpanningTree) AdjacentWeights(vertex int) []float64 {
	return msTree.weights[vertex]
}

func addWeightedEdge(adjacency [][]int, weights [][]float64, v1, v2 int, weight float64) {
	adjacency[v1] = append(adjacency[v1], v2)
	adjacency[v2] = append(adjacency[v2], v1)
	weights[v1] = append(weights[v1], weight)
	weights[v2] = append(weights[v2], weight)
}
