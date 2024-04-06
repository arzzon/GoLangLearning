/*
	GoLang does not support constructor by default.
	We can use Init() functions for the same.
*/
package main

import (
	"fmt"
)

type Car struct {
	Name  string
	Model string
}

func (c *Car) Init() {
	c.setName("DefaultName")
	c.setModel("DefaultModel")
}

// Go does not support function overloading(so we can use either the above method or the below one but not both) and does not support user defined operators.
/*func (c *Car) Init(n string, m string) {
	c.setName("DefaultName")
	c.setModel("DefaultModel")
}*/

func (c *Car) setName(s string) {
	c.Name = s
}

func (c *Car) getName() {
	fmt.Println("Car Name: %s", c.Name)
}

func (c *Car) setModel(m string) {
	c.Model = m
}

func (c *Car) getModel() {
	fmt.Println("Car Name: %s", c.Model)
}

func main() {

	var c1 = &Car{
		Name:  "Tata Indigo",
		Model: "TA30",
	}
	var c2 = new(Car)
	c2.Init()

	fmt.Printf("C1: %v\nC2: %v", *c1, *c2)

}
