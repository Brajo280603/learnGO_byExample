package main

import "fmt"

// GO supports anonymous functions , which can form closures
func intSeq() func() int {  // here this function returns a function
	i := 0
	return func() int {// heres we define the function anonymously , this function closes over the variable i to form a closure
		i++
		return i
	}
}

func main() {
	nextInt := intSeq() // here we call intSeq, assigning the result to nextInt, here it captures its own value which wil be updated next time we call nextInt()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()	// also check that the state is unique to that particular function , this function doesn't return the previous value 
	fmt.Println(newInts())
}