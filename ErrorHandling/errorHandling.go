/*
	The following is how error is defined in go:
	#############################################################################
	type error interface {
    	Error() string
	}
	#############################################################################
	So anything that implements the Error() string method can be used as an error.
*/

package main

import (
	"fmt"
)

type ErrorDividionByZero struct {
	ErrorCode string
}

func (e *ErrorDividionByZero) Error() string {
	return e.ErrorCode
}

func NewErrorDividionByZero() *ErrorDividionByZero {
	return &ErrorDividionByZero{
		ErrorCode: "Divide by zero error!",
	}
}

func main() {
	x := 1
	y := 0
	res, err := divide(x, y)
	if err != nil {
		fmt.Printf("Encountered an error: %v", err)
	} else {
		fmt.Printf("Division result: %v", res)
	}
}

func divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, NewErrorDividionByZero()
	}
	return x / y, nil

}
