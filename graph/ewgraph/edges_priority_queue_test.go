package ewgraph

import (
	"algorithms/graph/graph"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEdgesPQ(t *testing.T) {
	pq := newEdgesPriorityQueue()

	pq.pushEdge(graph.NewEdge(1, 2), 1.2)
	pq.pushEdge(graph.NewEdge(2, 0), 2.4)
	pq.pushEdge(graph.NewEdge(4, 1), 0.4)
	pq.pushEdge(graph.NewEdge(3, 4), 1.5)
	pq.pushEdge(graph.NewEdge(2, 4), 1.9)

	assert.Equal(t, 5, pq.Len())
	var edges []graph.Edge
	var weights []float64
	for pq.Len() > 0 {
		edge, weight := pq.popEdge()
		edges = append(edges, edge)
		weights = append(weights, weight)
	}
	assert.Equal(
		t,
		[]graph.Edge{
			graph.NewEdge(4, 1),
			graph.NewEdge(1, 2),
			graph.NewEdge(3, 4),
			graph.NewEdge(2, 4),
			graph.NewEdge(2, 0),
		},
		edges,
	)
	assert.Equal(t, []float64{0.4, 1.2, 1.5, 1.9, 2.4}, weights)
}
