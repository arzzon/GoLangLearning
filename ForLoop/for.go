package main

import (
	"fmt"
)

func main() {
	arr := [5]string{"a", "b", "c", "d", "e"}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}
