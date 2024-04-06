/*
Write a program which has a function/method that takes a list of elements as argument. The list can have elements which
are integers or another list or nested lists of integer values. The function/method should return a flattened list of
elements which are in order. For example the argument can be [1,2,3,4] or [1,2,[3,4]] or [4,2,3,[6,[1,[6,9,0],3]]] and
the return value will be 1,2,3,4 or 1,2,3,4 or 4,2,36,1,6,9,0 respectively.
*/

// USING STACK
package main

import (
	"fmt"
)

func flattenListUsingStack(elements []interface{}) []int {
	var flattened []int
	stack := make([]interface{}, len(elements))
	copy(stack, elements)

	for len(stack) > 0 {
		// Pop the top element from the stack
		elem := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		switch v := elem.(type) {
		case int:
			flattened = append(flattened, v)
		case []interface{}:
			// Push elements from nested list onto the stack
			for i := len(v) - 1; i >= 0; i-- {
				stack = append(stack, v[i])
			}
		default:
			fmt.Printf("Unsupported type: %T\n", v)
		}
	}

	return flattened
}

func main() {
	list1 := []interface{}{1, 2, 3, 4}
	list2 := []interface{}{1, 2, []interface{}{3, 4}}
	list3 := []interface{}{4, 2, 3, []interface{}{6, []interface{}{1, []interface{}{6, 9, 0}, 3}}}

	fmt.Println("Flattened list 1:", flattenListUsingStack(list1))
	fmt.Println("Flattened list 2:", flattenListUsingStack(list2))
	fmt.Println("Flattened list 3:", flattenListUsingStack(list3))
}
