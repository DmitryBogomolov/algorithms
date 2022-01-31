package strings

import "github.com/DmitryBogomolov/algorithms/sorting"

func RadixLSD(items []string, radix int, alph TrieAlphabet) {
	count := len(items)
	keys := make([]int, count)
	positions := make([]int, count)
	aux1 := make([]string, count)
	copy(aux1, items)
	aux2 := make([]string, count)
	for k := 0; k < radix; k++ {
		for i, item := range aux1 {
			keys[i] = getLSDIndex(item, k, alph)
		}
		sorting.KeyIndexedCounting(keys, alph.Size(), positions)
		for i, item := range aux1 {
			aux2[positions[i]] = item
		}
		copy(aux1, aux2)
	}
	copy(items, aux1)
}

func getLSDIndex(str string, idx int, alph TrieAlphabet) int {
	runes := []rune(str)
	i := len(runes) - 1 - idx
	if i >= 0 {
		return alph.ToIndex(runes[i])
	}
	return 0
}
