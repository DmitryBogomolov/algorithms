package utils

import "golang.org/x/exp/constraints"

// ResetList sets all list elements to -1.
func ResetList[T constraints.Signed](list []T) {
	for i := range list {
		list[i] = -1
	}
}

// ReverseList reverses list.
func ReverseList[T any](list []T) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}

// SumList returns sum of list items.
func SumList[T constraints.Float](list []T) T {
	sum := T(0.0)
	for _, item := range list {
		sum += item
	}
	return sum
}

// Min returns minimal element.
func Min[T constraints.Ordered](lhs, rhs T) T {
	var ret T
	if lhs < rhs {
		ret = lhs
	} else {
		ret = rhs
	}
	return ret
}
