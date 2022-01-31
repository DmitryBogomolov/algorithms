package sorting_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/sorting"
	"github.com/stretchr/testify/assert"
)

func TestKeyIndexedCounting(t *testing.T) {
	arr := []int{3, 2, 1, 2, 4, 7, 2, 4, 8}
	pos := KeyIndexedCounting(arr, 10)
	assert.Equal(t, []int{4, 1, 0, 2, 5, 7, 3, 6, 8}, pos)
}
