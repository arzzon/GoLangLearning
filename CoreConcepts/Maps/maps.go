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
	// SYNTAX:
	// map[index type] value type
	var m map[string]string
	m = make(map[string]string)
	m["name"] = "arbaaz"
	m["country"] = "Ind"
	fmt.Println(m) //Displays the key value pairs.

	// Deleting a key-value pair
	delete(m, "country")
	fmt.Println(m) //Displays the key value pairs.

	//Checking whether the map contains a key or not
	//Key should be absent
	if x, ok := m["age"]; ok {
		fmt.Printf("Key present: %v\n", x)
	} else {
		fmt.Println("Key absent")
	}
	// Key should be present
	if x, ok := m["name"]; ok {
		fmt.Printf("Key present: %v\n", x)
	} else {
		fmt.Println("Key absent")
	}
	fmt.Println("Calling testingMapPassedAsReference and passing map m")
	testingMapPassedAsReference(m)
	fmt.Println("Checking if map m is also modified in the caller side")
	for k, v := range m {
		fmt.Println(k, ":", v)
	}
	fmt.Println("Yes, map m is modified in the caller side.")
	fmt.Println("This proves map is always passed as reference to a function")
}

func testingMapPassedAsReference(m map[string]string) {
	fmt.Println("I'm testingMapPassedAsReference function and I've added a new key value to map m")
	m["new-key"] = "new-value"
}
