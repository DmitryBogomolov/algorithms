package huffman

import "errors"

// ErrEmptyData tells that data is nil or empty.
var ErrEmptyData = errors.New("data is nil or empty")

// ErrDataCorrupted tells that compressed data is corrupted.
var ErrDataCorrupted = errors.New("compressed data is corrupted")
