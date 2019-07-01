package main

import (  
    "fmt"
)

func digits(number int, digitChannel chan int) {  
    for number != 0 {
        digit := number % 10
        digitChannel <- digit
        number /= 10
    }

    close(digitChannel)
}
func calcSquares(number int, squareOperation chan int) {  
    sum := 0
    digitChannel := make(chan int)

    go digits(number, digitChannel)

    for digit := range digitChannel {
        sum += digit * digit
    }

    squareOperation <- sum
}

func calcCubes(number int, cubeOperation chan int) {  
    sum := 0
    digitChannel := make(chan int)

    go digits(number, digitChannel)

    for digit := range digitChannel {
        sum += digit * digit * digit
    }

    cubeOperation <- sum
}

func main() {  
    number := 589
    squareChannel := make(chan int)
    cubeChannel := make(chan int)

    go calcSquares(number, squareChannel)
    go calcCubes(number, cubeChannel)
    
    squares, cubes := <-squareChannel, <-cubeChannel
    fmt.Println("Final output", squares+cubes)
}
