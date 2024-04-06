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
 1. A slice literal is declared just like an array literal, except you leave out the element count:
    letters := []string{"a", "b", "c", "d"}
 2. A slice can be created with the built-in function called make, which has the signature:
    func make([]T, len, cap) []	T

Link: https://blog.golang.org/slices-intro
*/
package main

import (
	"fmt"
	"math/rand"
)

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

	// 2D slice
	fmt.Println("\n2D slice")
	/*
		[	|	|	|	|	|] Treat as ROWS
		  |	  |   |   |   |
		  v   v   v   v   v
		 s1   s2  s3  s4  s5
		s1 = [1,2,3] // slice with len 3
		s2 = [] // slice with len 0
	*/
	var slice2D [][]int
	initialRows := 3
	slice2D = make([][]int, initialRows) // [ | | |] Each cell pointing to an empty slice
	// Let's generate some random column lengths
	fmt.Println("2D Slice with number of rows: ", initialRows)
	var columns [3]int // array of length initialRows [IMP cann't mention initialRows directly as it's only evaluated by compiler at run time and for arrays we have to fix the size from the very beginning]
	for i := range columns {
		columns[i] = rand.Intn(10)
	}
	// create the generated columns
	for i, v := range columns {
		slice2D[i] = make([]int, v)
	}
	//Print new 2D slice now
	for i, _ := range slice2D {
		fmt.Println(slice2D[i])
	}
	// Increase the size of the 2D slice to support 2 more rows
	fmt.Println("2D Slice increase the number of rows by 2 ")
	// new columns for the new 2 rows
	newRows := make([][]int, 2) // new 2D slice that will be appended to the existing
	var newColumnsLens [2]int
	for i := range newColumnsLens {
		newColumnsLens[i] = rand.Intn(10)
	}
	// add the generated columns to the newRows
	for i, v := range newColumnsLens {
		newRows[i] = make([]int, v)
	}
	// Append the new 2D slice to the old
	slice2D = append(slice2D, newRows...)
	//Print new 2D slice now
	for i, _ := range slice2D {
		fmt.Println(slice2D[i])
	}

}
