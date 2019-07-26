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
	if x , ok := m["age"] ; ok{
		fmt.Printf("Key present: %v\n",x)
	} else{
		fmt.Println("Key absent")
	}
	// Key should be present
	if x , ok := m["name"] ; ok{
		fmt.Printf("Key present: %v\n",x)
        } else{
                fmt.Println("Key absent")
        }

}
