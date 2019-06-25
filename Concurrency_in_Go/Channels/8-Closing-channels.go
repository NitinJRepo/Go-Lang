//Senders have the ability to close the channel to notify receivers that no more data will be sent on the channel.
//Receivers can use an additional variable while receiving data from the channel to check whether the channel has been closed.
//v, ok := <- ch
//
//In the above statement ok is true if the value was received by a successful send operation to a channel. If ok is false it means that we are reading from a closed channel. 
//The value read from a closed channel will be the zero value of the channel's type. For example if the channel is an int channel, then the value received from a closed channel will be 0.

package main

import (
    "fmt"
)

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
	    ch <- i
	}
	close(ch)
}

func main() {
	channel := make(chan int)
	
	go producer(channel)

	for{
		v, ok :=  <-channel
		if ok == false {
		    break
		}
	    fmt.Println("received", v, ok)
	}
}
