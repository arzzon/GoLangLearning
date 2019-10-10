/*
LAMBDA/FUNCTION LITERAL:
Lambda/function literal represents a function without a name in golang.
IMPORTANT:
Function literal VS Lambda:
Lambda/function literal is a function without a name, but it doesn't enclose and keep the free variables within it like closure.
*/

package main
import (
	"fmt"
)

type Student struct{
	Name string
	Age int
}
/*
func less(s1 *Student,s2 *Student) bool{
	if s1.Age < s2.Age{
		return true
	}
	return false
}

func sort(students []*Student){
	for i:=0;i<len(students)-1;i++{
		for j:=0;j<len(students)-1-i;j++{
			if !less(student[j],students[j+1]){
				// Bubble sort
				temp := students[j]
				students[j] = students[j+1]
				students[j+1] = temp
			}
		}
	} 
}
*/
func sort(students []*Student){
        for i:=0;i<len(students)-1;i++{
                for j:=0;j<len(students)-1-i;j++{
                        if !func(s1 *Student,s2 *Student) bool{ //This is a function literal it can be replaced with a function name as done in the above commented lines.
				if s1.Age < s2.Age{
                			return true
        			}
        			return false
			   }(students[j],students[j+1]){
				
			   // Bubble sort
                           temp := students[j]
                           students[j] = students[j+1]
                           students[j+1] = temp
                        }
                }
        }
}

func main(){
	s1 := &Student{
		Name: "abc",
		Age: 25,
	}
	s2 := &Student{
                Name: "def",
                Age: 20,
        }
	s3 := &Student{
                Name: "ghi",
                Age: 30,
        }
	s4 := &Student{
                Name: "jkl",
                Age: 15,
        }

	ss := []*Student{s1,s2,s3,s4}
	sort(ss)
	fmt.Println("After sorting:")
	for i:=0;i<len(ss);i++{
		fmt.Printf("%+v ",*ss[i])
	}
}
