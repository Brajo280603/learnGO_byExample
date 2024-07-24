package main

import "fmt"

func main(){

	if 7%2 == 0 { // just an if else condition
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 8%4 == 0 { // if condition without else
		fmt.Println("8 is divisible by 4")
	}

	if 8%2 == 0 || 7%2 == 0 { // using logical operators (logical OR in this case)
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 { //You can write a statement before the conditions , any variable declared in this statement is available in the current and all subsequent branches
		fmt.Println(num, "is negetive")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// NOTE : parenthesis () around conditions are not required , however braces {} are required
}