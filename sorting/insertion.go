package sort

import builtinSort "sort"

func insertionCore(target builtinSort.Interface, lo int, hi int) {
	for i := lo; i <= hi; i++ {
		for j := i; j > 0 && target.Less(j, j-1); j-- {
			target.Swap(j, j-1)
		}
	}
}

// Insertion sorts target with "Insertion sort" algorithm.
func Insertion(target builtinSort.Interface) {
	insertionCore(target, 0, target.Len()-1)
}
