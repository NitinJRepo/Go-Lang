package main

import "fmt"

func main(){
	//declare one or more variables at once
	var a = "initial"
	fmt.Println(a)

	var b, c  int = 1, 2
	fmt.Println(b, c)

	//Go will infer the type of initialized variables.
	var d = true
	fmt.Println(d)

	//Variables declared without a corresponding initialization are zero-valued.
	var e int
	fmt.Println(e)

	//The := syntax is shorthand for declaring and initializing a variable
	f := "some value"
	fmt.Println(f)

}
