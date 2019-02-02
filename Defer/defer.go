package main

import (
	"fmt"
)

func main() {
	f()
	defer fmt.Println("defer1")
	defer fmt.Println("defer 2")
	fmt.Println("In main")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered")
		}
	}()
	g(0)
	fmt.Println("In f")
}
func g(i int) {
	if i > 2 {
		panic("In g") //Exits this function and goes
	}
	defer fmt.Println("defer in g")
	g(i + 1)
	fmt.Println("In End of g")
}
