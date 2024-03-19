/* *****************************************************************************
 *  Name:      Arbaaz Khan
 *  Language:  go
 *
 *  Description:  In this program is an example of how to work with variables.
 *                Like varibale declaration, intialization and use.
 *
 *  % go run variables.go
 *
***************************************************************************** */
package main

import "fmt"

func main() {
	//Ways to declare variables
	var age int              // Only declaration
	name := "Go"             // Declaration and initialization
	var x = 11               // In such cases the data type int is inferred from the value that we assign to them
	fmt.Println("Age=", age) //Age= 0 by default value of int variables is 0
	fmt.Println("name=", name)
	fmt.Println("x=", x)
	//Assignment after declaration
	age = 10
	fmt.Println("Age=", age)

	//Multiple variable declaration
	var a, b, c = 5, 1.9, "hi" // Or a, b, c := 4, 3, 5
	fmt.Println("a =", a, "b =", b, "c =", c)

	//Shadow variables
	fmt.Println("Variable shadowing")
	fmt.Println("shadow variable refers to a variable that is declared within a nested scope" +
		"with the same name as a variable in an outer scope. When this occurs, the outer variable is" +
		"said to be 'shadowed' by the inner variable.")
	var sv int
	sv = 100
	if sv == 100 {
		sv := 1 // sv is a new variable that shadows the sv present out of the if block
		fmt.Println("Variable sv shadowed with value, sv:", sv)
	}
	fmt.Println("No shadow effect on variable sv out of the block, sv:", sv)

}
