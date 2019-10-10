/*
Closure:
A closure is an implementation of a function, plus a table that binds values to the free variables that appear in the body of the function. A variable v is free in a function body f if v is used inside f, but it is not declared in f. Closures give the developer a way to pass functions around, together with some information about the context where these functions were created. Pragmatically speaking, closures allow the developer to write factories of functions.
Source: https://en.wikibooks.org/wiki/Introduction_to_Programming_Languages/Closures
*/

package main

import (
	"fmt"
)

// Declare a func type
type a func() int

// Function containing a closure, this function returns another function which is a closure.
func A() a{
	n := 0
	// Closure function
	return func() int{
		n++
		return n
	}
}

func main(){
	x := A()
	// Value of n will be stored within the closure and updated each time function is called.
	fmt.Println("Value of n =",x()) // n = 1
	fmt.Println("Value of n =",x()) // n = 2
	fmt.Println("Value of n =",x()) // n = 3
}
