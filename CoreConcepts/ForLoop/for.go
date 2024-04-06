/*
	There is no while loop in go, the only loop construct that is available is for loop.
	3 variations of for loop
	1) general for loop
	2) for loop with range
	3) infinite loop
*/
package main

import (
	"fmt"
)

func main() {
	arr := [5]string{"a", "b", "c", "d", "e"}
	// General for loop
	fmt.Println("General for loop")
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// For loop with range
	fmt.Println("\nFor loop with range")
	for index, value := range arr {
		fmt.Println(index, value)
	}

	// Another variant where only index is returned
	fmt.Println("\nAnother variant where only index is returned")
	for index := range arr {
		fmt.Println(index)
	}

	// Infinite loop
	fmt.Println("\nInfinite loop")
	i := 0
	for {
		if i == 10 {
			fmt.Println("Forcefully stopping infinite loop")
			break
		}
		fmt.Println(i)
		i++
	}
}
