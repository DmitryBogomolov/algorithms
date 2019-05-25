package sort

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQuickSorts(t *testing.T) {
	target := getData()

	Quick(target)

	assert.True(t, sort.IsSorted(target))
}

func TestQuickSortsAlreadySorted(t *testing.T) {
	target := getData()
	sort.Sort(target)

	Quick(target)

	assert.True(t, sort.IsSorted(target))
}
