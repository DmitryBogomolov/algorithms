package sort

import (
	"math/rand"
	builtinSort "sort"
)

func partition(target builtinSort.Interface, lo int, hi int) (int, int) {
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

func quickCore(target builtinSort.Interface, lo int, hi int) {
	if hi <= lo {
		return
	}
	lt, gt := partition(target, lo, hi)
	quickCore(target, lo, lt-1)
	quickCore(target, gt+1, hi)
}

// Quick sorts target with "Quick sort" algorithm.
func Quick(target builtinSort.Interface) {
	rand.Shuffle(target.Len(), target.Swap)
	quickCore(target, 0, target.Len()-1)
}
