/*
	Decorator is a design pattern which allows to add custom functionalities to an existing functionality
	without making any changes in the implementation of the latter.
*/
package main

import (
	"fmt"
	"time"
)

// work1 does some work
func work1() {
	time.Sleep(3 * time.Second)
}

// work2 does some other work
func work2() {
	time.Sleep(5 * time.Second)
}

// workHandler is another function which can do some specific work,
// however it is designed in such a way that along with doing it's
// own work, some other functionalities can also be added to it, like
// it can do the work of work1 and work2 as well(decoration).
func workHandler(work func()) {
	start := time.Now()
	work()
	end := time.Now()
	fmt.Printf("Work took %v seconds.\n", end.Unix()-start.Unix())
}

func main() {
	fmt.Println("DECORATOR EXAMPLE:")
	workHandler(work1)
	workHandler(work2)
}
