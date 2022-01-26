package strings

// TrieAlphabet describes all available trie symbols.
type TrieAlphabet interface {
	Size() int
	ToIndex(symbol rune) int
	ToSymbol(idx int) rune
}
