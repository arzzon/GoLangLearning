/*
	ldflags is also known as linker flag and it's used to provide such information/data/value which are dynamic in nature
	to the go program at build time. So the binary will have the value provided during build time.
	With -X flag we can write information to the variables at link time.
	Syntax:
	go build -ldflags="-X 'package_path.variable_name=new_value'"
*/
package main

import (
	"fmt"
)

var version string

func main() {
	fmt.Println("Version :", version)
}
