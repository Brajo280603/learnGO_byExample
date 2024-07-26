package main

import "fmt"

func vals() (int, int){ // the (int, int) signature shows that function returns 2 ints
	//GO has built-in supports for __multiple__ return values. This features is used often in indiomatic GO , for example to return both result and error values from a function
	return 3,7 
}

func main() {
	a, b := vals() // here we can get 2 different return values from the call with multiple assignments // here we can get 2 different return values from the call with multiple assignments // here we can get 2 different return values from the call with multiple assignments // here we can get 2 different return values from the call with multiple assignments
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals() // here we are only getting a subset of the returned value , so we ignore the first one with blank identifier (_)
	fmt.Println(c)
}

