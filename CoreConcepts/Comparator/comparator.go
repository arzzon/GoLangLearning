package main

import (
	"fmt"
)

type Comparator interface {
	Less(interface{}) bool
	Swap(interface{})
	Len()
	Sort(interface{})
}

type Student struct {
	Name string
	Age  int
}

func (s1 *Student) Less(s2 Student) bool {
	if s1.Age < s2.Age {
		return true
	}
	return false
}

//func (s *Student) Swap(interface{}) {}

//func (s *Student) Len() {}

func (*Student) Sort(interface{}) {
	fmt.Println("Sort")
}

func main() {
	fmt.Println("Hi")
	s1 := &Student{
		Name: "Arbaaz",
		Age:  50,
	}
	s2 := &Student{
		Name: "Suraj",
		Age:  34,
	}

	var S []*Student
	S = append(S, s1)
	S = append(S, s2)

	(s1).Sort(S)

}
