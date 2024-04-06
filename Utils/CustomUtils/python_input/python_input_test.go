package main

import (
	"fmt"
	"testing"
)

func TestInput(t *testing.T) {
	x := Input()
	if x == "hi" {
		fmt.Println("PASS")
	} else {
		fmt.Println("FAIL")
	}
}
