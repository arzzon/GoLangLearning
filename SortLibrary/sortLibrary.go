package main

import (
	"fmt"
	"sort"
)

type A struct {
	x []int
}

func (a *A) Len() int {
	return len(a.x)
}

func (a *A) Less(i int, j int) bool {
	if a.x[i] < a.x[j] {
		return true
	}
	return false
}

func (a *A) Swap(i int, j int) {
	a.x[j] = a.x[i] + a.x[j]
	a.x[i] = a.x[j] + a.x[i]
	a.x[j] = a.x[j] + a.x[i]
}

func main() {
	// Integer slice sorting
	x := []int{3, 2, 1}
	sort.Ints(x)
	fmt.Printf("%v", x)
	// Checks if sorted
	x = []int{3, 2, 1}
	r := sort.IntsAreSorted(x)
	fmt.Printf("%v", r)

	a := &A{
		x: []int{1, 2, 3},
	}
	r = sort.IsSorted(a)
	fmt.Printf("%v", r)
}
