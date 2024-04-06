package main

import (
	"flag"
	"fmt"
	"os"
)

var flagSet = flag.NewFlagSet("Java", flag.ContinueOnError)

var version = flagSet.Bool("v", false, "Display version.")
var path = flagSet.Bool("path", false, "Display the java binary path.")
var help = flagSet.Bool("help", false, "Display the available options.")

func main() {
	err := flagSet.Parse(os.Args[1:])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if *version {
		fmt.Println("v1.12")
	} else if *path {
		fmt.Println("/usr/local/java")
	} else if *help {
		fmt.Println("-v Version\n-path Path of the binary\n-help Displays all the options available.")
	} else {
		fmt.Println("Please check all the available options with go run java.go help=true")
	}
}
