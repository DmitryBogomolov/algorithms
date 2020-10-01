package strings

// Alphabet is all available symbols.
type Alphabet interface {
	Size() int
	ToIndex(symbol rune) int
	ToSymbol(idx int) rune
}
