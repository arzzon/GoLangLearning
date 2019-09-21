/*
	This program demonstrates the use of Go's inbuilt library `pprof` that can be used for profiling go programs.
*/
package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go func() {
		log.Printf("Starting Server! \t Go to http://localhost:6060/debug/pprof/\n")
		err := http.ListenAndServe("localhost:6060", nil)
		if err != nil {
			log.Printf("Failed to start the server! Error: %v", err)
			wg.Done()
		}
	}()
	wg.Add(1)
	go RandomProcessing()
	wg.Wait()
}

// Creating some random functions
func RandomProcessing() {
	for i := 0; i < 50; i++ {
		time.Sleep(5 * time.Second)
		wg.Add(1)
		go Sleepy()
	}
	wg.Done()
}

// This function runs for 50 seconds
func Sleepy() {
	time.Sleep(50 * time.Second)
	wg.Done()
}
