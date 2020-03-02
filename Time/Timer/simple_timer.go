/*
	Timer is used when we want a task to be performed after a certain
	amount of time.
	The amount of time that we want to wait for the task to be performed
	is passed on to the timer while creating it.
	Timer has a channel named "C" which returns after the specified amount of time.
	Timer{
		C chan time.Time
	}

	NOTE:
	One may think that timer does the same job as that of time.Sleep(),
	Using sleep() also we can wait for a specified amount of time and then execute the task,
	as we do using timer.
	But the key difference between timer and sleep is that we can cancel the timer anytime, but
	we cann't do that in case of sleep.
	For cancelling a timer we use: Stop() function.

*/

package main

import (
	"fmt"
	"time"
)

func task() {
	fmt.Println("Task completed.")
}

func main() {
	fmt.Println("Timer Started...")
	timer := time.NewTimer(5 * time.Second)
	//var input int
	//fmt.Scanln("Enter 1 to stop the timer!", &input)
	<-timer.C
	task()
}
