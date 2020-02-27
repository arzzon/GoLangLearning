/*
	There are different go tools that are available for go.
	They have been listed here.
*/
package main

import "fmt"

func main() {
	fmt.Println(` There are many go tools available:
	1. golint: 
	   >> Helps in detecting errors in go code.
	   >> golint example.go
	2. go doc
	   >> Shows the documentation of a particular package that we have asked it for
	   >> go doc <package-name> 
	3. gofmt
	   >> It takes care of formating, spaces, indentation etc.
	   >> gofmt expamle.go     // It only shows the go code after correcting it, but it never actually overwrites the code.
	   >> gofmt -w example.go  // It actually overwrites the go code.
	4. goimport
	   >> It searches for packages in the GOPATH, it is used in IDEs to search for packages that we have used in the code, but we have not imported them.
	   >> goimport searches for those packages and imports them.
	`)
}
