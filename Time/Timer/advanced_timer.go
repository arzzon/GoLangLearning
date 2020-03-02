/*
	Timer is used when we want a task to be performed after a certain
	amount of time.
	The amount of time that we want to wait for the task to be performed
	is passed on to the timer while creating it.
	Timer has a channel named "C" which returns after the specified amount of time.
	Timer{
		C chan time.Time
		Stop()
		Reset()
	}

	NOTE:
	One may think that timer does the same job as that of time.Sleep(),
	Using sleep() also we can wait for a specified amount of time and then execute the task,
	as we do using timer.
	But the key difference between timer and sleep is that we can cancel the timer anytime, but
	we cann't do that in case of sleep.
	For cancelling a timer we use: Stop() function.
	Functions available:


*/

package main

import (
	"fmt"
	"time"
)

// task prints a messafe of completion after the timer.
func task(t *time.Timer) {
	<-t.C
	fmt.Println("Task completed.")
}

func main() {
	fmt.Println("Timer Started...")
	timer := time.NewTimer(5 * time.Second)
	go task(timer)
	var input string
	fmt.Println("Enter 1 to stop the timer!")
	fmt.Scanln(&input)
	if input == "1" {
		fmt.Println("Timer stopped.")
		timer.Stop()
	}
	for {
	}
}
