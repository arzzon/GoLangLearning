package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}
	fmt.Println("Agr1=>", os.Args[1]) //This is the first argument and the 0th index has the path of the currently running executable 
}
