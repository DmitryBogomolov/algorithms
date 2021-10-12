package graph

func resetList(list []int) {
	for i := range list {
		list[i] = -1
	}
}

func reverseList(list []int) {
	for i, j := 0, len(list)-1; i < j; i, j = i+1, j-1 {
		list[i], list[j] = list[j], list[i]
	}
}

func min(lhs int, rhs int) int {
	var ret int
	if lhs < rhs {
		ret = lhs
	} else {
		ret = rhs
	}
	return ret
}
