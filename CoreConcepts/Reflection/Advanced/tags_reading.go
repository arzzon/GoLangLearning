package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name string `json:"student-name"`
	Age int	`json:"student-age"`
	Addr string `json:"student-address"`
}

func main() {
	st1 := Student{Name: "Arbaaz", Age: 60, Addr: "India"}
	t := reflect.TypeOf(st1)
	fmt.Println("Type of st1:", t)
	fmt.Println("Value of st1:", reflect.ValueOf(st1))
	fmt.Println("Value of st1:", reflect.ValueOf(st1).Kind())
	
}
