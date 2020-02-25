/*
	goto allows to transfer the control to an arbitary location.
	if goto is used then we write target labels to which control is transferred.
	These target labels can be defined anywhere in the code.

*/
package main

import "fmt"

func main() {
	// In this program the value of x will be reduced to 0 by decreamenting or increamenting by 1.
	var x int
	x = -3
Start: // This is a target label
	for {
		switch {
		case x < 0:
			goto Negative
		case x > 0:
			goto Positive
		default:
			break Start
		}
	Positive: // This is another target label
		fmt.Printf("%d ", x)
		x--
		continue Start
	Negative: // This is another target label
		fmt.Printf("%d ", x)
		x++
		continue Start
	}
	fmt.Println("\n Final value", x)
}
