/*
	Stringer is a interface that contains String() function.
    type Stringer interface{
		String() string
	}
    Packages like fmt look for this interface in print functions.
 */
package main

import "fmt"

type Person struct{
 	Name string
 	Age int
}

func (p *Person) String() string{
	return fmt.Sprintf("%s has age %d.", p.Name, p.Age)
}

func main(){
	p := &Person{
		Name: "AK",
		Age: 23,
	}
	fmt.Println(p)
}
