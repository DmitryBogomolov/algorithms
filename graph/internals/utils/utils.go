package utils

// ResetList sets all list elements to -1.
func ResetList(list []int) {
	for i := range list {
		list[i] = -1
	}
}

// ReverseList reverses list.
func ReverseList(list []int) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}

// SumList returns sum of list items.
func SumList(list []float64) float64 {
	sum := 0.0
	for _, item := range list {
		sum += item
	}
	return sum
}

// Min returns minimal element.
func Min(lhs int, rhs int) int {
	var ret int
	if lhs < rhs {
		ret = lhs
	} else {
		ret = rhs
	}
	return ret
}
