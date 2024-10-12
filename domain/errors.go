package main

import "errors"

// Custom error types
var (
	ErrBookNotFound  = errors.New("book not found")
	ErrInvalidAuthor = errors.New("invalid author")
)
