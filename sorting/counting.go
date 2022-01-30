package sorting

// Counting sorts items using counting sort algorithm.
func Counting(items []interface{}, keyRange int, keyFunc func(item interface{}) int) {
	count := make([]int, len(items)+1)
	aux := make([]interface{}, len(items))
	for _, item := range items {
		key := keyFunc(item)
		count[key+1]++
	}
	for i := 1; i < len(count); i++ {
		count[i] += count[i-1]
	}
	for _, item := range items {
		key := keyFunc(item)
		aux[count[key]] = item
		count[key]++
	}
	for i := range items {
		items[i] = aux[i]
	}
}
