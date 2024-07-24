package main

import "fmt"

func main(){
	var a [5]int	//a one dimensional integer array , you need to specify the length of the array (how many items the array should hold)
	fmt.Println("emp:", a) 

	a[4] = 100	//setting the value at an index of the array
	fmt.Println("set:", a)
	fmt.Println("get:", a[4]) //accessing one item on the array
	fmt.Println("len:", len(a)) //accessing the length of the array using the len function

	b := [5]int{1, 2, 3, 4, 5} //this syntax to declare and initialize an array
	fmt.Println("dcl:", b) 

	b = [...]int{1, 2, 3, 4, 5} //this to declare an array with length depending on the number of elements initialized , (the compiler takes care of this)
	fmt.Println("dcl:", b)

	b = [...]int{100, 3: 400, 500} //here 400 is set to the index 3 and the elements in between will be zeroed
	fmt.Println("idx:", b)

	var twoD [2][3]int		//a two dimensional array, array types are one dimensional but types can be composed to build multi-dimensional data structures
	for i := 0; i < 2; i++{
		for j := 0; j<3;j++{
			twoD[i][j] = i + j
		}
	}

	fmt.Println("2d:", twoD)

	twoD = [2][3]int{
		{1, 2, 3},
		{1, 2, 3},
	} //create and intialize multi-dimensional arrays
	fmt.Println("2d: ", twoD)
}