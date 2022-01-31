package sorting_test

import (
	"math/rand"
	"sort"
	"testing"

	. "github.com/DmitryBogomolov/algorithms/sorting"
	"github.com/stretchr/testify/assert"
)

func TestKeyIndexedCounting(t *testing.T) {
	{
		KeyIndexedCounting(nil, 0, nil)
	}
	{
		pos := make([]int, 1)
		KeyIndexedCounting([]int{2}, 10, pos)
		assert.Equal(t, []int{0}, pos)
	}
	{
		pos := make([]int, 9)
		KeyIndexedCounting(
			[]int{3, 2, 1, 2, 4, 7, 2, 4, 8},
			10,
			pos,
		)
		assert.Equal(t, []int{4, 1, 0, 2, 5, 7, 3, 6, 8}, pos)
	}
	{
		pos := make([]int, 4)
		KeyIndexedCounting(
			[]int{1, 2, 3, 4},
			100,
			pos,
		)
		assert.Equal(t, []int{0, 1, 2, 3}, pos)
	}
	{
		pos := make([]int, 4)
		KeyIndexedCounting(
			[]int{4, 3, 2, 1},
			100,
			pos,
		)
		assert.Equal(t, []int{3, 2, 1, 0}, pos)
	}
}

func TestKeyIndexedCounting_RandData(t *testing.T) {
	keys := make([]int, 10000)
	for i := range keys {
		keys[i] = rand.Intn(200)
	}
	pos := make([]int, len(keys))
	KeyIndexedCounting(keys, 200, pos)

	test := make([]int, len(keys))
	for i := range keys {
		test[pos[i]] = keys[i]
	}
	sort.IsSorted(sort.IntSlice(test))
}
