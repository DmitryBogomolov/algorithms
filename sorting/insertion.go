package sort

import builtinSort "sort"

// Insertion sorts target with "Insertion sort" algorithm.
func Insertion(target builtinSort.Interface) {
	for i := 0; i < target.Len(); i++ {
		for j := i; j > 0 && target.Less(j, j-1); j-- {
			target.Swap(j, j-1)
		}
	}
}
