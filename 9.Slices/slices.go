package main

import (
	"fmt"
	"slices"
)

func main() {

	var s []string //slices can be typed ,by the elements they contain(not the number of elements)

	fmt.Println("uninit:", s, s == nil, len(s) == 0) //uninitialized slices = nil and has length = 0

	s = make([]string, 3)	// this is how you make an empty slice with non-zero length
	fmt.Println("emp:", s, "len:", len(s), "cap:", cap(s))


	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	fmt.Println("len:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f") //in addition to basic operations just like array , slices support more , like append() which returns a slice containing one or more new values , NOTE: we need to accept a return value from append as we get a new slice value
	fmt.Println("apd:", s)

	c := make([]string, len(s))
	copy(c,s)	//to make a copy use the copy() , here we are copying s into c
	fmt.Println("cpy:", c)

	l := s[2:5]	//slices support a "slice" operator the syntax is like [low:high], with this you can get the slice of the elements s[2],s[3] and s[4]
	fmt.Println("sl1:", l)

	l = s[:5]	//here we are slicing upto (not including) s[5]
	fmt.Println("sl2:", l)

	l = s[2:]	//here we are slicing from (and including) s[2]
	fmt.Println("sl3:", l)

	t := []string{"g", "h", "i"} //how to declare and initialize a variable for slices
	fmt.Println("dcl:",t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {	//and example of a utility function contained in the "slices" package , here it compares them to see if both are equal
		fmt.Println("t == t2")
	}

	twoD := make([][]int, 3) //making multi-dimensional data structures with slices, the length of inner slices can vary , unlinke multi-dimensional arrays
	for i := 0; i<3;i++{
		innerLen := i+1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++{
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d:", twoD)

	//however slices are different types than arrays , both are rendered similarly by fmt.Println
}