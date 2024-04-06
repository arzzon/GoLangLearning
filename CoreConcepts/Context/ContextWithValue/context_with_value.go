/*
	context withValue
	It returns a copy of the parent context with new done channel and a key value
	pair as provided during it's creation. This key value is passed on to wherever
	the context is passed.
*/
package main

import (
	"context"
	"fmt"
)

type details string

func main() {
	fmt.Println("Context")
	ctx := context.Background()
	//ctx = context.WithValue(ctx, "name", "ak") // use of basic string type as key is discouraged

	k := details("name")
	ctx = context.WithValue(ctx, k, "ak")
	namePrint(ctx, k)
}

func namePrint(ctx context.Context, k details) {
	if v := ctx.Value(k); v != nil {
		fmt.Println("Received value from context key name = ", v)
	} else {
		fmt.Println("Value not found!")
	}
}
