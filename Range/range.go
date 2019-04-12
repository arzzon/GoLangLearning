package main

import (
	"fmt"
)

func main() {
	//range returns index and the values.
	slc := []int{1, 2, 3, 4, 5}
	for i, v := range slc { //Range does not always iterate over the slice/array in order
		fmt.Println("index: ", i, " value: ", v)
	}
}
