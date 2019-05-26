package sort

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHeapSorts(t *testing.T) {
	target := getData()

	Heap(target)

	assert.True(t, sort.IsSorted(target))
}

func TestHeapSortsAlreadySorted(t *testing.T) {
	target := getData()
	sort.Sort(target)

	Heap(target)

	assert.True(t, sort.IsSorted(target))
}
