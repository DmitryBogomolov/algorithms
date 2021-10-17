package ewgraph

func addWeightedEdge(adjacency [][]int, weights [][]float64, v1, v2 int, weight float64) {
	adjacency[v1] = append(adjacency[v1], v2)
	adjacency[v2] = append(adjacency[v2], v1)
	weights[v1] = append(weights[v1], weight)
	weights[v2] = append(weights[v2], weight)
}
