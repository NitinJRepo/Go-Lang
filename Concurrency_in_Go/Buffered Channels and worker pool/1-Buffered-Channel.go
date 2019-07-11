// All the channels we discussed in the previously were basically unbuffered. There we discussed sends and receives to an unbuffered channel are blocking.
// It is possible to create a channel with a buffer. Sends to a buffered channel are blocked only when the buffer is full. 
// Similarly receives from a buffered channel are blocked only when the buffer is empty.
// Buffered channels can be created by passing an additional capacity parameter to the make function which specifies the size of the buffer.
//
// ch := make(chan type, capacity)
//
// capacity in the above syntax should be greater than 0 for a channel to have a buffer.
// The capacity for an unbuffered channel is 0 by default and hence previously we omitted the capacity parameter while creating channels. 

package main

import (
	"fmt"
)

func main() {
	ch := make(chan string, 2)
	ch <- "Hello"
	ch <- "Nitin"
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
