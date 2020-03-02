package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Context")
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	go func() {
		time.Sleep(time.Second)
		fmt.Println("finished")
		cancel()
	}()
	sleepThenWrite(ctx, 10*time.Second, "This is an example of background context.")

}

func sleepThenWrite(ctx context.Context, t time.Duration, msg string) {
	select {
	case <-time.After(t):
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Print(ctx.Err())
	}
}
