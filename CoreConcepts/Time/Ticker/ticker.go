/*
	Ticker is used in scenarios where we have to perform a task repeatedly
	in certain interval of time.
	Once ticker is stopped it doesn't receive any more values.
*/
package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	fmt.Println("TICKER")
	tkr := time.NewTicker(2 * time.Second)
	done := make(chan struct{})
	go func() {
		for {
			select {
			case t := <-tkr.C:
				fmt.Printf("Periodic work completed at %v.\n", t)
			case <-done:
				fmt.Println("Received done signal.")
				return
				// default:
				// 	fmt.Println("t")
			}
		}
	}()
	time.Sleep(5 * time.Second) // wait for 5 seconds then stop the ticker
	tkr.Stop()                  // Stopping ticker using Stop(), now tkr.C channel will not receive any more values
	fmt.Println(strings.Repeat("*", 80))
	fmt.Println("Even if ticker is stopped still the select will be running, thats why we have used done to return from that function.")
	done <- struct{}{}          // stop signal is sent to the done channel
	time.Sleep(5 * time.Second) // This is just to allow the above go routine to display the message.
}
