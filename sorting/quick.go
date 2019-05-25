package sort

import (
	"math/rand"
	builtinSort "sort"
)

func partition(target builtinSort.Interface, lo int, hi int) int {
	i := lo
	j := hi + 1
	pivot := lo
	for {
		for {
			i++
			if i == hi || !target.Less(i, pivot) {
				break
			}
		}
		for {
			j--
			if j == lo || !target.Less(pivot, j) {
				break
			}
		}
		if i >= j {
			break
		}
		target.Swap(i, j)
	}
	target.Swap(pivot, j)
	return j
}

func quickCore(target builtinSort.Interface, lo int, hi int) {
	if hi <= lo {
		return
	}
	pos := partition(target, lo, hi)
	quickCore(target, lo, pos-1)
	quickCore(target, pos+1, hi)
}

// Quick sorts target with "Quick sort" algorithm.
func Quick(target builtinSort.Interface) {
	rand.Shuffle(target.Len(), target.Swap)
	quickCore(target, 0, target.Len()-1)
}
