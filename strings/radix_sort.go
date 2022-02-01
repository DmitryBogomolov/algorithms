package strings

import "github.com/DmitryBogomolov/algorithms/sorting"

func radixSort(items []string, radixCount int, alph Alphabet, getRune func(str string, idx int) rune) {
	count := len(items)
	keys := make([]int, count)
	positions := make([]int, count)
	src := items
	dst := make([]string, count)
	for k := 0; k < radixCount; k++ {
		for i, item := range src {
			keys[i] = getRadix(getRune(item, k), alph)
		}
		sorting.KeyIndexedCounting(keys, alph.Size(), positions)
		for i, item := range src {
			dst[positions[i]] = item
		}
		src, dst = dst, src
	}
	// After odd passes "src" points to internal array. So "items" must be explicitly updated.
	if radixCount%2 != 0 {
		copy(items, src)
	}
}

func getRadix(r rune, alph Alphabet) int {
	if r == -1 {
		return 0
	}
	return alph.ToIndex(r)
}

// RadixLSDSort sorts array of strings by character-by-character starting from the last one
// (least-significant-digit-first).
func RadixLSDSort(items []string, radixCount int, alph Alphabet) {
	radixSort(items, radixCount, alph, getLSDRadix)
}

func getLSDRadix(str string, idx int) rune {
	runes := []rune(str)
	i := len(runes) - 1 - idx
	if i >= 0 {
		return runes[i]
	}
	return -1
}

// RadixMSDSort sorts array of strings by character-by-character starting from the first one
// (most-significant-digit-first).
func RadixMSDSort(items []string, radixCount int, alph Alphabet) {
	radixSort(items, radixCount, alph, getMSDRadix)
}

func getMSDRadix(str string, idx int) rune {
	runes := []rune(str)
	if idx < len(runes) {
		return runes[idx]
	}
	return -1
}
