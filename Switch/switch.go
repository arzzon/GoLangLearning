package main

import (
	"fmt"
)

func main() {
	// No need of Break statements in switch
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
}
