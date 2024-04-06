/*
Write a program which has a function/method that takes a list of elements as argument. The list can have elements which
are integers or another list or nested lists of integer values. The function/method should return a flattened list of
elements which are in order. For example the argument can be [1,2,3,4] or [1,2,[3,4]] or [4,2,3,[6,[1,[6,9,0],3]]] and
the return value will be 1,2,3,4 or 1,2,3,4 or 4,2,36,1,6,9,0 respectively.
*/

// USING RECURSION
package main

import (
	"fmt"
	"reflect"
)

func flattenList(elements []interface{}) []int {
	var flattened []int

	for _, elem := range elements {
		switch v := elem.(type) {
		case int:
			flattened = append(flattened, v)
		case []interface{}:
			flattened = append(flattened, flattenList(v)...)
		default:
			fmt.Printf("Unsupported type: %v\n", reflect.TypeOf(v))
		}
	}

	return flattened
}

func main() {
	list1 := []interface{}{1, 2, 3, 4}
	list2 := []interface{}{1, 2, []interface{}{3, 4}}
	list3 := []interface{}{4, 2, 3, []interface{}{6, []interface{}{1, []interface{}{6, 9, 0}, 3}}}

	fmt.Println("Flattened list 1:", flattenList(list1))
	fmt.Println("Flattened list 2:", flattenList(list2))
	fmt.Println("Flattened list 3:", flattenList(list3))
}
