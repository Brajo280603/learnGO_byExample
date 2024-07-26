package main

import "fmt"

func fact(n int) int { //This calles itself until the base condition of fact(0)
	if n == 0 {
		return 1
	}

	return n * fact(n-1)
}

func main(){
	fmt.Println(fact(7))

	var fib func(n int) int // Closures can be recursive , but this requires the closure to be declared with a typed var explicitly before it's defined

	fib = func(n int) int {
		if n < 2 {
			return n
		}

		return fib(n-1) + fib(n-2)// as fib was previously declared in main , GO knows which function to call with fib here.
	}

	fmt.Println(fib(7)) 
}