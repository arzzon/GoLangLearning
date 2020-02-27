/*
	Select allows us to select multiple channels.
	Select doesn't get blocked while receiving from channel.
	So, we can perform any action even if we didn't receive
	any data in the channel.
*/
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

type done struct{}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan done)
	wg.Add(1)
	go func(ch1 chan int, ch2 chan done) {
		defer wg.Done()
		for {
			time.Sleep(1 * time.Second)
			select {
			case r := <-ch1:
				fmt.Println("Received: ", r)
			case <-ch2:
				fmt.Println("Termination signal received.")
				return
			default:
				fmt.Println("Working...")
			}
		}
	}(ch1, ch2)
	time.Sleep(5 * time.Second)
	ch1 <- 10
	time.Sleep(3 * time.Second)
	ch2 <- done{}
	wg.Wait()

	// Select timeout
	fmt.Println("\nSELECT TIMEOUT")
	go func() {
		defer close(ch2) // returns handle to channel
		for {

		}
	}()

	select {
	case <-ch2:
		fmt.Println("Goroutine finished work!!!!")
	case <-time.After(2 * time.Second): // After 2 seconds the select terminates if nothing could be received from ch2
		// The time.After() returns a channel which closes after the mentioned duration [IMPORTANT]
		fmt.Println("Sorry, the goroutine took too long.")
	}

}
