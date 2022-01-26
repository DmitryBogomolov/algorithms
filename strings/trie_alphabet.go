package strings

// TrieAlphabet describes all available trie symbols.
type TrieAlphabet interface {
	// Size gets alphabet size.
	Size() int
	// ToIndex converts symbol to index.
	ToIndex(symbol rune) int
	// ToSymbol converts index to symbol.
	ToSymbol(idx int) rune
}

// RangeAlphabet contains symbols from a specified range.
type RangeAlphabet struct {
	size        int
	startSymbol rune
}

// Size gets alphabet size.
func (alph RangeAlphabet) Size() int {
	return alph.size
}

// ToIndex converts symbol to index.
func (alph RangeAlphabet) ToIndex(symbol rune) int {
	return int(symbol-alph.startSymbol) % alph.size
}

// ToSymbol converts index to symbol.
func (alph RangeAlphabet) ToSymbol(idx int) rune {
	return alph.startSymbol + rune(idx)
}

// NewRangeAlphabet creates RangeAlphabet with a specified range.
func NewRangeAlphabet(start rune, end rune) RangeAlphabet {
	return RangeAlphabet{
		size:        int(end) - int(start) + 1,
		startSymbol: start,
	}
}

// ASCIIAlphabet contains 0..127 symbols.
var ASCIIAlphabet TrieAlphabet = NewRangeAlphabet(0, 127)
