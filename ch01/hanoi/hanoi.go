// Package hanoi is a Towers of Hanoi solver
package hanoi

import (
	"fmt"
	"os"
	"path"
)

// A Tower represents one of the Towers of Hanoi.
type Tower []int

// push pushes a disk onto a tower.
func (t *Tower) push(n int) {
	if len(*t) > 0 && (*t)[len(*t)-1] >= n {
		fmt.Fprintf(os.Stderr, "%s: invalid state: pushing %d onto %v",
			path.Base(os.Args[0]), n, *t)
		os.Exit(1)
	}
	*t = append(*t, n)
}

// pop pops a disk off a tower.
func (t *Tower) pop() (result int) {
	result = (*t)[len(*t)-1]
	*t = (*t)[:len(*t)-1]
	return result
}

// String returns a tower's state in native int slice format.
func (t Tower) String() string {
	return fmt.Sprint([]int(t))
}

// Solve moves disks from one tower to another, using a third temporary
// tower, maintaining the rules of the Towers of Hanoi.
func Solve(from, to, temp *Tower, n int) {
	if n < 1 {
		return
	}
	if n == 1 {
		to.push(from.pop())
		return
	}
	Solve(from, temp, to, n-1)
	Solve(from, to, temp, 1)
	Solve(temp, to, from, n-1)
}
