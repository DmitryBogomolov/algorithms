package sort

func heapSink(target Interface, pos int, len int) {
	for k := pos; 2*k <= len; {
		p := 2 * k
		if p < len && target.Less(p-1, p) {
			p++
		}
		if !target.Less(k-1, p-1) {
			break
		}
		target.Swap(k-1, p-1)
		k = p
	}
}

// Heap sorts target with "Heap sort" algorithm.
func Heap(target Interface) {
	len := target.Len()
	for k := len / 2; k >= 1; k-- {
		heapSink(target, k, len)
	}
	for len > 1 {
		target.Swap(0, len-1)
		len--
		heapSink(target, 1, len)
	}
}
