// Package fib calculates Fibonacci numbers.
package fib

// Plain calculates Fibonacci numbers using a naive algorithm.
func Plain(n uint) uint {
	if n < 2 {
		return n
	}
	return Plain(n-2) + Plain(n-1)
}

var fibs = map[uint]uint{0: 0, 1: 1}

// Memo calculates Fibonacci numbers using memoization.
func Memo(n uint) uint {
	if result, ok := fibs[n]; ok {
		return result
	}
	fibs[n] = Memo(n-2) + Memo(n-1)
	return fibs[n]
}

// Iter calculates Fibonacci numbers iteratively.
func Iter(n uint) uint {
	if n == 0 {
		return 0
	}
	var last, next uint = 0, 1
	for i := uint(1); i < n; i++ {
		last, next = next, last+next
	}
	return next
}
