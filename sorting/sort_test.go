package sorting_test

import (
	"math/rand"
	"sort"
	"testing"

	. "github.com/DmitryBogomolov/algorithms/sorting"
	"github.com/stretchr/testify/assert"
)

type sortFunc = func(sort.Interface)
type sortDesc struct {
	name string
	f    sortFunc
}

var sampleData []int

var items = []sortDesc{
	{"Insertion", Insertion},
	{"Shell", Shell},
	{"Quick", Quick},
	{"Heap", Heap},
}

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

func TestMethods(t *testing.T) {
	for _, item := range items {
		t.Run(item.name, func(t *testing.T) {
			target := getData()
			item.f(target)
			assert.True(t, sort.IsSorted(target))
		})
	}
}

func BenchmarkMethods(b *testing.B) {
	for _, item := range items {
		b.Run(item.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				target := getData()
				b.StartTimer()
				item.f(target)
			}
		})
	}
}
