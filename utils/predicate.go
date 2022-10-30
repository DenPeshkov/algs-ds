package utils

// Predicate represents a boolean-valued function of one argument.
type Predicate[T any] func(v T) bool
