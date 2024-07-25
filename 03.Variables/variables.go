package main

import "fmt"

func main(){
	var a = "initial" // var declares a with String type inferred by go as it is also initialized with a string in declaration
	fmt.Println(a)

	var b,c = 1,2 // var can declare multiple variables
	fmt.Println(b,c)

	var d = true // declaration and initialization of boolean variable
	fmt.Println(d)

	var e int // declaration of integer 
	fmt.Println(e)

	f:= "apple" //this is the shorthand syntax for declaring and initializing a variable, NOTE: this syntax is only available inside functions.
	fmt.Println(f)
}