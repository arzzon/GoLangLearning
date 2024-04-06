/*
	python_input package contains methods similar to python that we use to read inputs in competitive programming.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// StringType represents a string
type StringType string

var reader = bufio.NewReader(os.Stdin)

// Input reads the entire line from the std input.
func Input() StringType {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Enter a string>> ")
	inputStr, err := reader.ReadString('\n') // Reads a string till enter is clicked
	if err != nil {
		log.Print(err)
	}
	// fmt.Printf("String read = %s\n", inputStr)
	return StringType(inputStr)
}

// Strip method removes the leading and trailing chars from the string
func (s StringType) Strip(cutset ...string) StringType {
	if len(cutset) > 0 && cutset[0] != " " {
		fmt.Println(cutset[0])
		return StringType(strings.Trim(string(s), cutset[0]))
	}
	return StringType(strings.TrimSpace(string(s)))
}

// Split returns the slice of strings after splitting the string using the separator
func (s StringType) Split(sep string) []string {
	return strings.Split(string(s), sep)
}

func main() {
	fmt.Println(Input().Strip().Split(" "))
	// fmt.Printf("%c", Input()[2])
}
