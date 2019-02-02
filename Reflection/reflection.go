package main

import (
	"fmt"
	"reflect"
)

type vertex struct {
	X int
	Y int
}

func main() {
	i := 9.0
	X1 := vertex{X: 0, Y: 1}
	fmt.Println("Type of i:", reflect.TypeOf(i))
	fmt.Println("Type of X1:", reflect.TypeOf(X1))
	fmt.Println("Type of X1:", reflect.ValueOf(X1))
	//fmt.Println("Type of i:", reflect.TypeOf(i).Name)
}
