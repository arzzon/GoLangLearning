/*
   Hashset is a data structure which stores elements in a hash table and each element
   can be accessed in BigO(1) time using hash function. Unlike hash table it does not have
   key value pair rather it has only keys those are hashed.
*/

package main

import (
	"fmt"
	"reflect"
)

type HashSet struct {
	hs map[interface{}]struct{}
}

func NewHashSet() *HashSet {
	return &HashSet{
		hs: map[interface{}]struct{}{},
	}
}

func (hs *HashSet) Add(val interface{}) {
	hs.hs[val] = struct{}{}
}

func (hs *HashSet) Contains(ele interface{}) bool {
	found := false
	switch ele.(type) {
	case byte:
		_, found = hs.hs[ele.(byte)]
	case int:
		_, found = hs.hs[ele.(int)]
	case float32:
		_, found = hs.hs[ele.(float32)]
	case float64:
		_, found = hs.hs[ele.(float64)]
	default: // Default case is string
		_, found = hs.hs[ele.(string)]
	}
	return found
}

func (hs *HashSet) Remove(ele interface{}) {
	delete(hs.hs, ele)
}

func (hs *HashSet) String() string {
	res := "{ "
	for k, _ := range hs.hs {
		switch k.(type) {
		case byte:
			res += fmt.Sprintf("%s ", k.(byte))
		case int:
			res += fmt.Sprintf("%d ", k.(int))
		case float32:
			res += fmt.Sprintf("%f ", k.(float32))
		case float64:
			res += fmt.Sprintf("%f ", k.(float64))
		default: // Default case is string
			res += fmt.Sprintf("%s ", k.(string))
		}
	}
	res += "}"
	return res
}

func (hs *HashSet) Check() {
	for k, v := range hs.hs {
		if 2 == k {
			print("true")
		}
		fmt.Print(reflect.TypeOf(k), v)
	}
}

func main() {
	//hashset = map[string]struct{}{} // value occupies zero bytes of space
	hs := NewHashSet()
	fmt.Print("Added arbaaz, ak, khan.\n")
	hs.Add("arbaaz")
	hs.Add("ak")
	hs.Add("khan")
	fmt.Printf("Does it contain ak: %t\n", hs.Contains("ak"))
	fmt.Print("Removed ak.\n")
	hs.Remove("ak")
	fmt.Printf("Does it contain ak: %t\n", hs.Contains("ak"))
	fmt.Print("Contents of the set.\n")
	println(hs.String())

	// Int hash set
	hsInt := NewHashSet()
	fmt.Print("Added 1,2,3,4\n")
	hsInt.Add(1)
	hsInt.Add(2)
	hsInt.Add(3)
	hsInt.Add(4)
	fmt.Printf("Does it contain 3: %t\n", hs.Contains(3))
	fmt.Print("Removed 3.\n")
	hsInt.Remove(3)
	fmt.Printf("Does it contain 3: %t\n", hs.Contains(3))
	fmt.Print("Contents of the set.\n")
	print(hsInt.String())
	//hsInt.Check()
}

/*
The size of a struct is the sum of the size of the types of its fields, since there are no fields: no size!
*/
