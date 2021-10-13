package digraph

import (
	"algorithms/graph/internals"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransitiveClosure(t *testing.T) {
	target := internals.NewTestDigraph(6,
		0, 1,
		1, 2,
		0, 3,
		3, 2,
		4, 3,
		4, 5,
		5, 4,
	)

	ret := BuildTransitiveClosure(target)

	assert.Equal(t, true, ret.Reachable(0, 2), "0 - 2")
	assert.Equal(t, false, ret.Reachable(1, 0), "1 - 0")
	assert.Equal(t, true, ret.Reachable(3, 2), "3 - 2")
	assert.Equal(t, true, ret.Reachable(4, 3), "4 - 3")
	assert.Equal(t, true, ret.Reachable(4, 2), "4 - 2")
	assert.Equal(t, false, ret.Reachable(5, 0), "5 - 0")
	assert.Equal(t, true, ret.Reachable(5, 2), "5 - 2")
}
