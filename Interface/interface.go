/* *****************************************************************************
 *  Name:      Arbaaz Khan
 *  Language:  go
 *
 *  Description:  In this program we have explored interfaces in go.
 *
 *  % go run interface.go
 *
***************************************************************************** */
package main

import "fmt"

type Customer struct{
	Name string
	Age int
	Address string
}
type Client interface{
	PrintName()
	PrintAge()
	PrintAddress()
}
func main() {
       // Interfaces are used in two different ways.
       
       // 1. It's used as contract to create a type.
       // It has signatures of methods that a particular type should be defining
       c := Customer{
       		Name: "Arbaaz",
		Age: 100,
		Address: "India",
       }
       c.PrintName()
}

func (c Customer) PrintName(){
	fmt.Printf("Customer name: %v\n", c.Name)
}

