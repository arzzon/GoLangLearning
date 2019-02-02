package main

import (
	"fmt"
	"time"
)

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
		time.Sleep(100 * time.Millisecond)
		fmt.Println(v)
	}
	c <- sum
}
func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	res1, res2 := <-c, <-c
	fmt.Println(res1, res2)
}
