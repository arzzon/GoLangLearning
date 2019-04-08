package main
import "fmt"
func main(){
	// Ways to declare slices
	s1 := make([]int, 3, 10)
	s2 := make([]int, 3)
	s3 := []int{}
	s4 := []int{1,2,3,4}
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)

	// Methods available for slices

	s4 = append(s4,5)
        fmt.Println(s4)
        fmt.Println(cap(s1))
        fmt.Println(len(s2))

	ret := copy(s4, s2) //The builtin copy(dst, src) copies min(len(dst), len(src)) elements.
	fmt.Println("Number of elements copied: ",ret)
	fmt.Println("Copied slice: ", s4)

}
