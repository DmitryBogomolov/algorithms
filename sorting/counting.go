package sorting

// KeyIndexedCounting sorts by positive integer keys in ascending order.
// Accepts array of keys and maximum key value. Returns array of positions.
func KeyIndexedCounting(keys []int, keyBound int) []int {
	count := make([]int, keyBound+1)
	for _, key := range keys {
		count[(key%keyBound)+1]++
	}
	for i := 0; i < keyBound; i++ {
		count[i+1] += count[i]
	}
	positions := make([]int, len(keys))
	for i, key := range keys {
		positions[i] = count[key%keyBound]
		count[key%keyBound]++
	}
	return positions
}
