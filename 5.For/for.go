package main

import "fmt"

func main(){
	
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	fmt.Println("end of first loop")

	for j := 0 ; j < 3 ; j++ { //the classic initial / condition / after for loop
		fmt.Println(j)
	}

	fmt.Println("end of second loop")

	for i := range 3 { // this is using iteration with range over an integer
		fmt.Println("range ", i)
	}

	fmt.Println("end of third loop")

	for {	//basically a while(true) loop
		fmt.Println("loop")
		break //break the loop using "break"
	}

	fmt.Println("end of fourth loop")

	for n := range 6 { 
		if n%2 == 0 {
			continue // skip a iteration using "continue"
		}
		fmt.Println(n)
	}

	fmt.Println("end of fifth loop")
}