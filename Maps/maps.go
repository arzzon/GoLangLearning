/* *****************************************************************************
 *  Name:      Arbaaz Khan
 *  Language:  go
 *
 *  Description:  In this program we have explored maps in go.
 *
 *  % go run maps.go
 *
***************************************************************************** */
package main

import "fmt"


func main() {
	//SYNTAX:
	// map[index type] value type
	var m map[string]string
	m = make(map[string]string)
	m["name"] = "arbaaz"
	fmt.Println(m["name"])

}
