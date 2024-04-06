package main

import "fmt"

func main() {
	//Difference between array and slice is that in array we fix the length but in slice the length can vary

	// Ways to declare an array
	var a [3]int   //int array with length 3
	fmt.Println(a) //default value is 0
	// Assigning values to the array
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a)

	// Simultaneous eclaration and initialisation
	b := [3]int{12, 78, 50} // short hand declaration to create array
	fmt.Println(b)
	c := [...]int{12, 78, 50} // ... makes the compiler determine the length
	fmt.Println(c)

	// 2D arrays
	fmt.Println("2D ARRAY:")
	twoDArray := [2][2]int{
		{1, 2},
		{3, 4},
	}
	for i := 0; i < len(twoDArray); i++ {
		for j := 0; j < len(twoDArray[i]); j++ {
			fmt.Print("%d ", twoDArray[i][j])
		}
		println()
	}
}
