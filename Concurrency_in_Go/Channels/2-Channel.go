// The syntax to send and receive data from a channel are:
// data := <- ch // Read from channel ch
// ch <- data  // Write to channel ch

package main

import (
    "fmt"
    //"time"
)

func hello(done chan bool) {
    fmt.Println("Hello world goroutine")
    done <- true
}

func main() {
    //Short hand declaration of channel
    done := make(chan bool)

    go hello(done)

    //time.Sleep(1 * time.Second)  <-- Use channel instead
    <-done
    fmt.Println("main function")
}
