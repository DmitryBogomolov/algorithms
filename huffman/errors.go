package huffman

import (
	"errors"
	"fmt"
)

// ErrEmptyData tells that data is nil or empty.
var ErrEmptyData = errors.New("data is nil or empty")

// CompressError is returned when compress failed.
type CompressError struct {
	err error
}

func (e CompressError) Error() string {
	return fmt.Sprintf("failed to compress: %v", e.err)
}

func (e CompressError) Unwrap() error {
	return e.err
}

// ExpandError is returned when expand failed.
type ExpandError struct {
	err error
}

func (e ExpandError) Error() string {
	return fmt.Sprintf("failed to expand: %v", e.err)
}

func (e ExpandError) Unwrap() error {
	return e.err
}
