/*
Package utils provides common functionality used by various algorithms and data structures.
*/
package utils

/*
Comparator returns:
  - -1 if a is less than b,
  - 0 if a is equal to b,
  - 1 if a is greater than b.
*/
type Comparator[T any] func(a, b T) int
