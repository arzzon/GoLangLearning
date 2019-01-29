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

}
