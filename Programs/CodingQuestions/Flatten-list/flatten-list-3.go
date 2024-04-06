/*
Write a program which has a function/method that takes a list of elements as argument. The list can have elements which
are integers or another list or nested lists of integer values. The function/method should return a flattened list of
elements which are in order. For example the argument can be [1,2,3,4] or [1,2,[3,4]] or [4,2,3,[6,[1,[6,9,0],3]]] and
the return value will be 1,2,3,4 or 1,2,3,4 or 4,2,36,1,6,9,0 respectively.
*/

// USING RECURSION and CUSTOM TYPE
package main

import "fmt"

type NestedList interface{}

type FlatList []int

func (fl *FlatList) flattenHelper(nl NestedList) {
	switch v := nl.(type) {
	case int:
		*fl = append(*fl, v)
	case []NestedList:
		for _, elem := range v {
			fl.flattenHelper(elem)
		}
	}
}

func Flatten(nl NestedList) FlatList {
	fl := FlatList{}
	fl.flattenHelper(nl)
	return fl
}

func main() {
	list1 := []NestedList{1, 2, 3, 4}
	list2 := []NestedList{1, 2, []NestedList{3, 4}}
	list3 := []NestedList{4, 2, 3, []NestedList{6, []NestedList{1, []NestedList{6, 9, 0}, 3}}}

	fmt.Println("Flattened list 1:", Flatten(list1))
	fmt.Println("Flattened list 2:", Flatten(list2))
	fmt.Println("Flattened list 3:", Flatten(list3))
}
