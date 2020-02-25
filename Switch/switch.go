/*
	There are basically 2 types of switch in go
	1) Switch with expression
	2) Switch with type
	In go switch cases are not fallthrough by default, which means that if a matching condition is found the control gets out of the switch,
	So, we don't need break statements in go switch. But if we want the fallthrough to happen then we can use go keyword "fallthrough".
*/
package main

import (
	"fmt"
	//"math/rand"
)

func main() {
	// No need of Break statements in switch
	fmt.Println("Switch case no break statement needed:")
	x := "a"
	switch x {
	case "b":
		fmt.Println("b")
	case "a":
		fmt.Println("a")
	case "c":
		fmt.Println("c")
	default:
		fmt.Println("Aphabet")
	}

	// No need of Break statements in switch
	//switch y:= rand.Int31n(50) {
	fmt.Println("Switch case with break which is not required:")
	y:= 20 
	switch y {
        case 10:
                fmt.Println("b")
		break
        case 20:
                fmt.Println("a")
		break
        case 40:
                fmt.Println("c")
		break
        default:
                fmt.Println("Aphabet")
        }

	// Fallthrough
	fmt.Println("Switch case with fallthrough:")
        y = 5
        switch { // if no expression is mentioned then the case with true condition is considered.
        case y < 2:
                fmt.Println("b")
		fallthrough
        case y < 10:
                fmt.Println("a")
		fallthrough
        case y < 20:
                fmt.Println("c")
		fallthrough
        default:
                fmt.Println("Aphabet")
        }

	// Switch case with multiple conditions in cases
	fmt.Println("Switch case with multiple conditions:")
	y = 5
        switch { // if no expression is mentioned then the case with true condition is considered.
        case y < 2, y < 10: // two conditions seperated by comma
                fmt.Println("condition is true as y < 10 although y not < 2")
        case y < 10:
                fmt.Println("a")
        case y < 20:
                fmt.Println("c")
        default:
                fmt.Println("Aphabet")
        }


	// Type switch
	fmt.Println("Type Switch case:")
        var i interface{}
	i = 8
        switch i.(type){ // i's type is considered (important non interface cann't be used here)
        case string:
                fmt.Println("string")
        case float64:
                fmt.Println("float")
        case int:
                fmt.Println("int")
        default:
                fmt.Println("Unknown")
        }

}
