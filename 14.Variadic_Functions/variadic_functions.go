package main

import "fmt" 


func sum(nums ...int){ //here this function will take an arbitrary number of ints as arguments
	fmt.Print(nums, " ")//within this function the type of nums is equivalent to []int. We can call len(nums) and iterate over it with range etc.
	total := 0

	for _, num := range nums {
		total += num
	}

	fmt.Println(total)
}

func main(){

	sum(1,2) //the call of variadic functions work the same way with individual arguments
	sum(1,2,3)

	nums := []int{1,2,3,4}
	sum(nums...)	//if we have multiple args in a slice , we can apply them to a variadic function using funcName(slice...)
}