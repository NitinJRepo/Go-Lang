//The for range form of the for loop can be used to receive values from a channel until it is closed.

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

    for v := range channel {
	fmt.Println("Received ", v )
    }
}
