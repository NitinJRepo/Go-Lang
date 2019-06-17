package main

import (
    "fmt"
)

func calculateSquares(number int, squareOperation chan int) {
    sum := 0
    for number != 0 {
        digit := number % 10
        sum += digit * digit
        number /= 10
    }
    squareOperation <- sum
}

func calculateCubes(number int, cubeOperation chan int) {
    sum := 0
    for number != 0 {
        digit := number % 10
        sum += digit * digit * digit
        number /= 10
    }
    cubeOperation <- sum
}

func main() {
    number := 589

    squarechannel := make(chan int)
    cubeChannel := make(chan int)

    go calculateSquares(number, squarechannel)
    go calculateCubes(number, cubeChannel)

    squares, cubes := <-squarechannel, <-cubeChannel
    fmt.Println("Final output", squares + cubes)
}
