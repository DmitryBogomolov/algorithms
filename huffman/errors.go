package huffman

import (
	"errors"
	"fmt"
)

// ErrEmptyData tells that data is nil or empty.
var ErrEmptyData = errors.New("data is nil or empty")

// DataCorruptedError tells that compressed data is corrupted.
type DataCorruptedError struct {
	err error
}

func (e DataCorruptedError) Error() string {
	return fmt.Sprintf("compressed data is corrupted: %v", e.err)
}

func (e DataCorruptedError) Unwrap() error {
	return e.err
}
