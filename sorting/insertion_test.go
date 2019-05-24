package sort

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var sampleData []int

func init() {
	sampleData = make([]int, 200)
	for i := 0; i < len(sampleData); i++ {
		sampleData[i] = rand.Intn(100)
	}
}

func getData() sort.IntSlice {
	data := make([]int, len(sampleData))
	copy(data, sampleData)
	return sort.IntSlice(data)
}

func TestInsertionSorts(t *testing.T) {
	target := getData()

	Insertion(target)

	assert.True(t, sort.IsSorted(target))
}

func TestInsertionSortsAlreadySorted(t *testing.T) {
	target := getData()
	sort.Sort(target)

	Insertion(target)

	assert.True(t, sort.IsSorted(target))
}
