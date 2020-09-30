package sorting

import "sort"

func insertionCore(target sort.Interface, lo int, hi int) {
	for i := lo; i <= hi; i++ {
		for j := i; j > 0 && target.Less(j, j-1); j-- {
			target.Swap(j, j-1)
		}
	}
}

// Insertion sorts using *Insertion sort* algorithm.
// https://algs4.cs.princeton.edu/21elementary/
func Insertion(target sort.Interface) {
	insertionCore(target, 0, target.Len()-1)
}
