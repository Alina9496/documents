package dto

import "errors"

var (
	ErrInvalidKey   = errors.New("invalid key")
	ErrInvalidLimit = errors.New("the limit must be greater than 0")
	ErrEmptyValue   = errors.New("empty value")
)
