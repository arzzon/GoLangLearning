/* *****************************************************************************
 *  Name:      Arbaaz Khan
 *  Language:  go
 *
 *  Description:  This program is an example of how to work with constants.
 *
 *  % go run constants.go
 *
***************************************************************************** */
package main

import "fmt"

func main() {
	const a = 6 //int constant declaration, value cann't be changed again
	fmt.Println("a=", a)
	const b = 45.6 //float constant declaration, value cann't be changed again
	fmt.Println("b=", b)
	const c = "hello" //string constant declaration, value cann't be changed again
	fmt.Println("c=", c)
}
