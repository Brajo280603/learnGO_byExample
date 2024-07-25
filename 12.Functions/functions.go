package main

import "fmt"

func plus(a int, b int) int {	//this function takes two ints and returns their sum as an int
	return a+b
}

func plusPlus(a,b,c int) int {	//when having multiple consecutive parameters of the same type , the type name can be omitted for the like-typed parameters up to the final parameter that declares the type
	//GO requires explicit returns it won't automatically return the value of the last experssion
	return a+b+c
}

func main() {
	res := plus(1,2)	//function call , syntax ; name(args)
	fmt.Println("1+2 = ", res)

	res = plusPlus(1,2,3)
	fmt.Println("1+2+3 = ", res)
}