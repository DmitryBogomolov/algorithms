package unionfind_test

import (
	"fmt"
	"testing"

	. "github.com/DmitryBogomolov/algorithms/unionfind"
	"github.com/stretchr/testify/assert"
)

func TestUnionFind(t *testing.T) {
	uf := New(10)

	check := func(expected []int) {
		for i := 0; i < 10; i++ {
			assert.Equal(t, uf.Find(i), expected[i], fmt.Sprintf("%d - %d", i, expected[i]))
		}
	}

	assert.Equal(t, uf.Count(), 10)
	check([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	uf.Union(0, 1)
	uf.Union(4, 5)
	uf.Union(5, 7)
	uf.Union(5, 8)
	uf.Union(4, 9)

	assert.Equal(t, uf.Count(), 5)
	check([]int{0, 0, 2, 3, 4, 4, 6, 4, 4, 4})

	assert.True(t, uf.Connected(1, 0), "1 and 0")
	assert.False(t, uf.Connected(0, 2), "0 and 2")
	assert.True(t, uf.Connected(5, 9), "5 and 9")
	assert.True(t, uf.Connected(4, 8), "4 and 8")
	assert.False(t, uf.Connected(2, 7), "2 and 7")
}
