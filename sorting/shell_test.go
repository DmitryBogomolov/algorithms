package sort

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShellSorts(t *testing.T) {
	target := getData()

	Shell(target)

	assert.True(t, sort.IsSorted(target))
}

func TestShellSortsAlreadySorted(t *testing.T) {
	target := getData()
	sort.Sort(target)

	Shell(target)

	assert.True(t, sort.IsSorted(target))
}
