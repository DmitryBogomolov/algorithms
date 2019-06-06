package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestVerticesPQ(t *testing.T) {
	pq := newVerticesPQ(10)

	pq.update(2, 1.2)
	pq.update(3, 0.3)
	pq.update(5, 3.1)
	pq.update(8, 2.2)
	pq.update(0, 6.1)
	pq.update(3, 5.4)
	pq.update(8, 0.9)
	pq.update(1, 2.5)

	assert.Equal(t, 6, pq.Len())
	var data []int
	for pq.Len() > 0 {
		data = append(data, pq.pop())
	}
	assert.Equal(t, []int{8, 2, 1, 5, 3, 0}, data)
}
