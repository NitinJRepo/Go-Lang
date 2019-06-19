//It is also possible to create unidirectional channels, that is channels that only send or receive data.
package main

import "fmt"

func sendData(sendch chan<- int) {
    sendch <- 10
}

func main() {
    sendch := make(chan<- int) //unidirectional channel
    go sendData(sendch)
    fmt.Println(<-sendch)
}

