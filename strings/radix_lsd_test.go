package strings_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/strings"
	"github.com/stretchr/testify/assert"
)

func TestRadixLSD(t *testing.T) {
	alph := NewRangeAlphabet('a', 'z')

	{
		RadixLSD(nil, 0, alph)
	}
	{
		items := []string{"test"}
		RadixLSD(items, 10, alph)
		assert.Equal(t, []string{"test"}, items)
	}
	{
		items := []string{
			"dab", "add", "cab", "fad", "fee", "bad", "dad", "bee", "fed", "bed", "ebb", "ace",
		}
		RadixLSD(items, 3, alph)
		assert.Equal(
			t,
			[]string{
				"ace", "add", "bad", "bed", "bee", "cab", "dab", "dad", "ebb", "fad", "fed", "fee",
			},
			items,
		)
	}
	{
		items := []string{
			"dab", "add", "cab", "fad", "fee", "bad", "dad", "bee", "fed", "bed", "ebb", "ace",
		}
		RadixLSD(items, 2, alph)
		assert.Equal(
			t, []string{
				"dab", "cab", "fad", "bad", "dad", "ebb", "ace", "add", "fed", "bed", "fee", "bee",
			},
			items,
		)
	}
	{
		items := []string{
			"dab", "add", "cab", "fad", "fee", "bad", "dad", "bee", "fed", "bed", "ebb", "ace",
		}
		RadixLSD(items, 1, alph)
		assert.Equal(
			t,
			[]string{
				"dab", "cab", "ebb", "add", "fad", "bad", "dad", "fed", "bed", "fee", "bee", "ace",
			},
			items,
		)
	}
	{
		items := []string{
			"dab", "add", "cab", "fad", "fee", "bad", "dad", "bee", "fed", "bed", "ebb", "ace",
		}
		RadixLSD(items, 0, alph)
		assert.Equal(
			t,
			[]string{
				"dab", "add", "cab", "fad", "fee", "bad", "dad", "bee", "fed", "bed", "ebb", "ace",
			},
			items,
		)
	}
}
