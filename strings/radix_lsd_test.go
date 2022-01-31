package strings_test

import (
	"testing"

	. "github.com/DmitryBogomolov/algorithms/strings"
	"github.com/stretchr/testify/assert"
)

func TestRadixLSD(t *testing.T) {
	items := []string{
		"dab",
		"add",
		"cab",
		"fad",
		"fee",
		"bad",
		"dad",
		"bee",
		"fed",
		"bed",
		"ebb",
		"ace",
	}
	RadixLSD(items, 3, NewRangeAlphabet('a', 'z'))
	assert.Equal(t, []string{
		"ace",
		"add",
		"bad",
		"bed",
		"bee",
		"cab",
		"dab",
		"dad",
		"ebb",
		"fad",
		"fed",
		"fee",
	}, items)
}
