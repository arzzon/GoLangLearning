package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func routine(site string) {
	// Decrement the counter when the goroutine completes.
	defer wg.Done()

	// do something
	fmt.Println(site)
}

func main() {
	sites := []string{
		"https://www.google.com/",
		"https://duckduckgo.com/",
		"https://www.bing.com/"}

	for _, site := range sites {
		// Increment the WaitGroup counter.
		wg.Add(1)

		go routine(site)
	}

	// wait all goroutines to finish
	wg.Wait()
}
