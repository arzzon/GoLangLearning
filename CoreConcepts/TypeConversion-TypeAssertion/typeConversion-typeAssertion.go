/*
TypeConversion:
Type conversion in Go is used to convert a value from one type to another, provided that the conversion
is allowed by the language rules.

TypeAssertion:
Type assertion in Go is used to extract the concrete value from an interface value and test whether the
underlying concrete type implements a certain interface.

TypeConversion vs TypeAssertion:
1) Type conversion is used to convert values between compatible types, while type assertion is used
to extract the underlying concrete value from an interface value and test its type.
2) Type conversion is a compile-time operation, while type assertion is a runtime operation.
3) Type conversion is explicit and deterministic, while type assertion can fail at runtime if the
underlying concrete type does not match the asserted type.
4) Type conversion is typically used for simple conversions between types, while type assertion is
used when working with interfaces and dynamic types.
*/
package main

import (
	"fmt"
)

func main() {
	typeConversionTest()
	typeAssertionTest()
}

func typeConversionTest() {
	// Type conversion example: converting int to float64
	var x int = 10
	var y float64 = float64(x) // Convert int to float64
	fmt.Println(y)             // Outputs: 10
}

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func typeAssertionTest() {
	// Type assertion example: asserting interface value to concrete type
	var s Shape
	s = Circle{Radius: 5}

	if circle, ok := s.(Circle); ok {
		fmt.Println("Circle radius:", circle.Radius)
	} else {
		fmt.Println("Not a circle")
	}
}
