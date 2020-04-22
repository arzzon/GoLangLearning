/*
	Slice is a composite type which uses array internally but it's size is dynamic.
	Internally it creates an array and there is a pointer to the start of the array.
	The way it is defined:
	type slice struct{
		array unsafe.Pointer
		len int
		cap int
	}
	How to declare a slice:
	1) A slice literal is declared just like an array literal, except you leave out the element count:
	   letters := []string{"a", "b", "c", "d"}
	2) A slice can be created with the built-in function called make, which has the signature:
		func make([]T, len, cap) []	T
	Link: https://blog.golang.org/slices-intro

*/
package main

import "fmt"

func main() {
	// Ways to declare slices
	s1 := make([]int, 3, 10) // make() returns a slice value not a pointer although
	s2 := make([]int, 3)
	s3 := []int{}
	s4 := []int{1, 2, 3, 4}
	var s5 []string
	s6 := s4[1:3] // Slice can be created by slicing another slice or array
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s6)
	// Methods available for slices

	s4 = append(s4, 5)
	fmt.Println(s4)
	fmt.Println(cap(s1)) // Capacity is the maximum number of elements that the slice can store
	fmt.Println(len(s2)) // Length is the current length of the slice

	// Copy slice
	/*
		X Wrong way X
		s7 = s6 // What it does is that it creates a new slice value, but it internally points to the same array
	*/
	// Clone using inbuilt copy() function
	ret := copy(s4, s2) //The builtin copy(dst, src) copies min(len(dst), len(src)) elements.
	fmt.Println("Number of elements copied: ", ret)
	fmt.Println("Copied slice: ", s4)

}
