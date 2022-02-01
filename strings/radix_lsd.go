package strings

import "github.com/DmitryBogomolov/algorithms/sorting"

// RadixLSD sorts array of strings by character-by-character starting from the last one
// (least-significant-digit-first).
func RadixLSD(items []string, radixCount int, alph TrieAlphabet) {
	count := len(items)
	keys := make([]int, count)
	positions := make([]int, count)
	src := items
	dst := make([]string, count)
	for k := 0; k < radixCount; k++ {
		for i, item := range src {
			keys[i] = getLSDIndex(item, k, alph)
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

func getLSDIndex(str string, idx int, alph TrieAlphabet) int {
	runes := []rune(str)
	i := len(runes) - 1 - idx
	if i >= 0 {
		return alph.ToIndex(runes[i])
	}
	return 0
}
