/*
	Singleton class is way to ensure that only one object of a class can be created.
	In go neither we have class nor do we have inbuilt singleton class.
	But we can build one using structs and the properties of a singleton class.
*/
package main

import "fmt"

var instance *singleton

type singleton struct {
	name string
}

func NewSingleton() *singleton {
	if instance == nil {
		instance = &singleton{}
	}
	return instance
}

func (s *singleton) getName() string {
	return s.name
}

func main() {
	fmt.Println("Creating object for the first time...")
	sglton1 := NewSingleton()
	fmt.Println("Setting name to arbaaz1 for the first object...")
	sglton1.name = "arbaaz1"
	fmt.Println("Get name from the first object.")
	fmt.Println(sglton1.getName())

	fmt.Println()

	fmt.Println("Trying to create new object...")
	sglton2 := NewSingleton()
	fmt.Println("Checking the value of the member name.")
	fmt.Println(sglton2.getName())
	fmt.Println("Setting name to arbaaz2 to the current object.")
	sglton2.name = "arbaaz2"
	fmt.Println("Get name from the current object.")
	fmt.Println(sglton2.getName())
	fmt.Println("Get name from the first object.")
	fmt.Println(sglton1.getName())

	fmt.Println("\nWe can clearly see that both the objects have the same value, as both are same instances\n To put it in another way only one object has been created.")
}
