package sorting

import (
	"math/rand"
	"sort"
)

// QuickToInsertionCutoff defines threshold when quick sort is switched to insertion sort.
const QuickToInsertionCutoff = 11

func partition(target sort.Interface, lo int, hi int) (int, int) {
	lt := lo
	gt := hi
	for i := lo + 1; i <= gt; {
		if target.Less(i, lt) {
			target.Swap(lt, i)
			lt++
			i++
		} else if target.Less(lt, i) {
			target.Swap(i, gt)
			gt--
		} else {
			i++
		}
	}
	return lt, gt
}

func quickCore(target sort.Interface, lo int, hi int) {
	if hi-lo <= QuickToInsertionCutoff {
		insertionCore(target, lo, hi)
		return
	}
	lt, gt := partition(target, lo, hi)
	quickCore(target, lo, lt-1)
	quickCore(target, gt+1, hi)
}

// Quick sorts using *Quick sort* algorithm.
// https://algs4.cs.princeton.edu/23quicksort/
func Quick(target sort.Interface) {
	rand.Shuffle(target.Len(), target.Swap)
	quickCore(target, 0, target.Len()-1)
}
