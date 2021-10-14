package ewgraph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerticesIndexPQ(t *testing.T) {
	pq := newVerticesIndexPriorityQueue(10)

	pq.updateVertex(2, 1.2)
	pq.updateVertex(3, 0.3)
	pq.updateVertex(5, 3.1)
	pq.updateVertex(8, 2.2)
	pq.updateVertex(0, 6.1)
	pq.updateVertex(3, 5.4)
	pq.updateVertex(8, 0.9)
	pq.updateVertex(1, 2.5)

	assert.Equal(t, 6, pq.Len())
	var data []int
	for pq.Len() > 0 {
		data = append(data, pq.popVertex())
	}
	assert.Equal(t, []int{8, 2, 1, 5, 3, 0}, data)
}
