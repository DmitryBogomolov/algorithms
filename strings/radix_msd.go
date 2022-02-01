package strings

import "github.com/DmitryBogomolov/algorithms/sorting"

// RadixMSD sorts array of strings by character-by-character starting from the first one
// (most-significant-digit-first).
func RadixMSD(items []string, radixCount int, alph TrieAlphabet) {
	count := len(items)
	keys := make([]int, count)
	positions := make([]int, count)
	src := items
	dst := make([]string, count)
	for k := 0; k < radixCount; k++ {
		for i, item := range src {
			keys[i] = getMSDIndex(item, k, alph)
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

func getMSDIndex(str string, idx int, alph TrieAlphabet) int {
	runes := []rune(str)
	if idx < len(runes) {
		return alph.ToIndex(runes[idx])
	}
	return 0
}
