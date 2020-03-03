package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	fmt.Println("Context")
	ctx := context.Background()
	sleepThenWrite(ctx, 3*time.Second, "This is an example of background context.")

}

func sleepThenWrite(ctx context.Context, t time.Duration, msg string) {
	time.Sleep(t)
	fmt.Println(msg)
}
