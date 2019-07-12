// we write 3 strings to a buffered channel of capacity 2. When the control reaches the third write, the write is blocked since the channel has exceeded its capacity. 
// Now some Goroutine must read from the channel in order for the write to proceed, but in this case there is no concurrent routine reading from this channel.
// Hence there will be a deadlock and the program will panic at run time

package main

import (  
    "fmt"
)

func main() {  
    ch := make(chan string, 2)
    ch <- "Hello"
    ch <- "Nitin"
    ch <- "Sample program for Deadlock using buffered channel"
    fmt.Println(<-ch)
    fmt.Println(<-ch)
}
