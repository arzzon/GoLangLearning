/*
	Channels in go are used for communication between go routines.
	It is a type defined by "chan T", where T is the type of chan.
	There are 3 operations that can be performed on the channels:
	1) Send:      ch <- <some value>
	2) Recieve:   <- ch
	3) Close:     close(ch)
	There are broadly two types of channels:
	1) Based on the size of channel:
		|----------->  Buffered: Can hold more values in a queue and is non blocking unless it's full.
		|                        While sending if it's full then it's blocking, otherwise it's non blocking.
		|                        While receiving if there is no element in the channel then it's blocking, otherwise it's non blocking.
		|----------->  UnBuffered: Can hold only one value and is blocking.
	2) Based on the direction of data flow to/from the channel.
		|-----------> Bidirectional channel: Both send and receive can be done.
		|-----------> Unidirectional channel: It's capable of either sending or receiving.
							|
							|-----------> Send only channel: data can only be sent to it, but cann't be received from it.
							|             Declaration:     var ch chan<- int
							|-----------> Receive only channel: data can only be received from the channel, but cann't be sent to it.
							|             Declaration:     var ch <-chan int
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	// Normal bidirectional channel
	var chB chan int
	chB = make(chan int) // It's compulsary to make channels, declaring by var keyword and sending channels to goroutines won't work, it will fail with channel nil error.
	go forBidirectionalCh(chB)
	// Send to channel
	fmt.Println("Sent 1 to channel in main:")
	chB <- 1
	// Receive
	fmt.Println("Recieved in main:", <-chB)

	fmt.Println()

	fmt.Println("For ReceiveOnlyChannel")
	go forUnidirectionalReceiveOnlyCh(chB)
	// Send to channel
	fmt.Println("Sent 1 to channel in main:")
	chB <- 4

	fmt.Println()

	fmt.Println("For SendOnlyChannel")
	go forUnidirectionalSendOnlyCh(chB)
	// Receive
	fmt.Println("Recieved in main:", <-chB)

	fmt.Println()
	// Closing a channel
	fmt.Println("Closing the channel:")
	close(chB)

	// Check if channel is closed or not
	_, ok := <-chB
	if !ok {
		fmt.Println("Channel closed successfully.")
	} else {
		fmt.Println("Channel has not been closed, some value might have been received but we have muted it since we know that we have already closed it.")
	}

	fmt.Println()

	// Buffered Channels
	fmt.Println("Buffered channels")
	chBuff := make(chan int, 5) // It can hold 2 values without blocking
	// Sending values to buffered chan channel
	for i := 1; i <= 5; i++ {
		chBuff <- i // Can hold a max of 5 values due to its size
	}
	go forBufferedChannel(chBuff)
	time.Sleep(1 * time.Second) // Just for the shake of main not terminating earlier than goroutine
	fmt.Println("\nExit")
}

func forBidirectionalCh(ch chan int) {
	// Receive
	fmt.Println("Received in forBidirectionalCh:", <-ch)
	// Send
	fmt.Println("Sent 0 to channel in forBidirectionalCh:")
	ch <- 2
}

// Here we have used a sendonly channel
func forUnidirectionalSendOnlyCh(ch chan<- int) {
	// Can't recReceive
	fmt.Println("Can't receive in forUnidirectionalSendOnlyCh.") //,<-ch)
	// Send
	fmt.Println("Sent 0 to channel in forUnidirectionalSendOnlyCh:")
	ch <- 3
}

// Here we have used a receive only channel
func forUnidirectionalReceiveOnlyCh(ch <-chan int) {
	// Recieve
	fmt.Println("Received in forUnidirectionalReceiveOnlyCh:", <-ch)
	// Can't Send
	fmt.Println("Can't send anything to channel in forUnidirectionalReceiveOnlyCh.")
	// ch <- 0
}

func forBufferedChannel(ch chan int) {
	fmt.Println("Receiving from buffered channel")
	for i := 0; i <= 5; i++ {
		fmt.Print(" ", <-ch) // Reading those 5 values
	}
}
