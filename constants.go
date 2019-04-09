package main

import "fmt"
import "math"

const s string = "constant"

func main(){
	fmt.Println(s)

	//Uncomment and try this. 
	//s = "another value"

	const n = 500000000

	//Constant expressions perform arithmetic with arbitrary precision
	const d = 3e20 / n
	fmt.Println(d)

	//A numeric constant has no type until it’s given one, such as by an explicit cast
	fmt.Println(int64(d))

	//A number can be given a type by using it in a context that requires one, such as a variable assignment or function call. For example, here math.Sin expects a float64
	fmt.Println(math.Sin(n))
}

