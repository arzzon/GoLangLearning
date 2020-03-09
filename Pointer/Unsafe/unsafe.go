/*
	Unsafe is a package in go that contains operations to step around the type safety in golang.
	One of it's uses is that we can convert two different structures with the same members between
	them.
*/
package main

import (
	"fmt"
	"unsafe"
)

type animal struct {
	name   string
	age    int
	gender string
}

type human struct {
	name   string
	age    int
	gender string
}

func main() {
	fmt.Println("Unsafe")
	var a animal
	a = animal{
		name:   "tom",
		age:    10,
		gender: "male",
	}
	var h human
	//h = (human)a    //NOT ALLOWED

	h = *(*human)(unsafe.Pointer(&a))

	fmt.Printf("After initialisation of animal we have animal = %v\n", a)
	fmt.Printf("After converting animal to human we have human = %v\n", h)
	fmt.Println("Sizeof function returns the size of the parameter provided in bytes:")
	fmt.Println("Size of structs unsafe.Sizeof(a) =", unsafe.Sizeof(a), "bytes which is {[name]string(16), [age]int(8), [gender]string(16)}")
	fmt.Println("unsafe.Sizeof(true) =", unsafe.Sizeof(true), "bytes")
}

/*
	Explanation:
	a = animal{				|  								10000     10016    10024       10040
		name:   "tom",		|								|-------------------------------|--------------------
		age:    10,			|								|  "tom"   |  10   |   "male"   |   storing some other data
		gender: "male",		|								|-------------------------------|-------------------
	}														|	 |		   |		 |				|
															|    |         |		 |				|
	________________________________________________________|	 |		   |___		 |				|
    |															 |__________   |	 |				|
	|																		|  |	 |				|
	|>>>>>> h = *(*human)(unsafe.Pointer(&a)) 			h = human{			|  |	 |				|
																	name: <<|  |	 |				|
																	age:  <<<<<| 	 |				|
																	gender:	<<<<<<<<<|				|
																	ifexists:<<<<<<<<<<<<<<<<<<<<<<<| garbage will be stored
															}

*/

/*
type animal struct {
	name   string
	age    int
	gender string
}

type human struct {
	name   string
	age    int
	gender string
	a      int       // This will be storing some garbage data whatever will be present in the location
}

func main() {
	fmt.Println("Unsafe")
	var a animal
	a = animal{
		name:   "tom",
		age:    10,
		gender: "male",
	}
	var h human
	h = *(*human)(unsafe.Pointer(&a))
	fmt.Println("%v", h)
}
*/
