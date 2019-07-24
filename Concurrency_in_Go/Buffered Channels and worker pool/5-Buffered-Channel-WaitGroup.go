// A WaitGroup is used to wait for a collection of Goroutines to finish executing. The control is blocked until all Goroutines finish executing.
// Lets say we have 3 concurrently executing Goroutines spawned from the main Goroutine. 
// The main Goroutines needs to wait for the 3 other Goroutines to finish before terminating. This can be accomplished using WaitGroup.

package main

import (
	"fmt"
	"sync"
	"time"
)

func process(i int, wg *sync.WaitGroup) {
	fmt.Println("Started goroutine: ", i)
	time.Sleep(2 *  time.Second)
	fmt.Println("Ended goroutines: ", i)

	wg.Done()
}

func main() {
	no := 3
	var wg sync.WaitGroup

	for i := 0; i < no; i++ {
		wg.Add(1)
		go process(i, &wg)
	}
	wg.Wait()
	fmt.Println("All goroutines finished executing")

}


// It is important to pass the address of "wg".
// If the address is not passed, then each Goroutine will have its own copy of the WaitGroup and main goroutine will not be notified when they finish executing.
