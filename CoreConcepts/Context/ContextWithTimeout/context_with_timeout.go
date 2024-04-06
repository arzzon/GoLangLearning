/*
	context withTimeout
	It returns a copy of the parent context with new done channel and a cancel function.
	In context.WithTimeout we provide a timeout value and the context will be automatically
	cancelled after that time. Once the context is cancelled the Done channel is closed.
	We know that a channel can return two values x, ok := <-ch,
	So when a channel is closed it returns false, which is stored in ok in the above example.
*/
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Context")
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, 3*time.Second)
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
