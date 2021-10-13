package internals

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
