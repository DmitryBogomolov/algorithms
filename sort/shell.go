package sort

// Shell sorts target with "Shell sort" algorithm.
func Shell(target Interface) {
	len := target.Len()
	h := 1
	for h < len/3 {
		h = 3*h + 1
	}
	for h >= 1 {
		for i := h; i < len; i++ {
			for j := i; j >= h && target.Less(j, j-h); j -= h {
				target.Swap(j, j-h)
			}
		}
		h /= 3
	}
}
