/*
	Select doesn't get blocked while receiving from channel.
	So, we can perform any action even if we didn't receive
	any data in the channel.
*/
package main
import (
	"fmt"
	"time"
	"sync"
)
var wg sync.WaitGroup
type done struct{}
func main(){
	ch1 := make(chan int)
	ch2 := make(chan done)
	wg.Add(1)
	go func(ch1 chan int, ch2 chan done){
		defer wg.Done()
		for{
			time.Sleep(1*time.Second)
			select{
			case r := <- ch1:
				fmt.Println("Received: ",r)
			case <- ch2:
				fmt.Println("Termination signal received.")
				return
			default:
				fmt.Println("Working...")
			}
		}
	}(ch1,ch2)
	time.Sleep(5*time.Second)
	ch1 <- 10
	time.Sleep(3*time.Second)
	ch2 <- done{}
	wg.Wait()
}
