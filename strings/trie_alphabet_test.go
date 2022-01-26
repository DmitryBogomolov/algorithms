package strings_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/strings"
	"github.com/stretchr/testify/assert"
)

func TestTrieRangeAlphabet(t *testing.T) {
	alph := NewRangeAlphabet('a', 'z')

	assert.Equal(t, 26, alph.Size())
	assert.Equal(t, 0, alph.ToIndex('a'))
	assert.Equal(t, 1, alph.ToIndex('b'))
	assert.Equal(t, 2, alph.ToIndex('c'))
	assert.Equal(t, 'z', alph.ToSymbol(25))
	assert.Equal(t, 'y', alph.ToSymbol(24))
	assert.Equal(t, 'x', alph.ToSymbol(23))
}

func TestTrieASCIIAlphabet(t *testing.T) {
	alph := ASCIIAlphabet

	assert.Equal(t, 128, alph.Size())
	assert.Equal(t, 48, alph.ToIndex('0'))
	assert.Equal(t, '9', alph.ToSymbol(57))
}
