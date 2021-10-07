package internals

// SumList returns sum of list items.
func SumList(list []float64) float64 {
	sum := 0.0
	for _, item := range list {
		sum += item
	}
	return sum
}
