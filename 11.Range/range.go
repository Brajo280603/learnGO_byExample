package main

import "fmt"

func main(){

	//Range Iterates over elements in a variety of data structures. 

	nums := []int {2,3,4} 
	sum := 0
	for _, num := range nums {//here range is used to sum the numbers in a slice. as we didnt need the index , we ignored it with blank (_) identifier
		sum += num
	}
	fmt.Println("sum:", sum)

	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	kvs := map[string]string{"a": "apple", "b":"banana"}	
	for k, v := range kvs {	//range on map iterates over key/value pairs
		fmt.Printf("%s -> %s\n", k, v) //NOTE : printf behaves like printf in C , while println just prints the values 
	}

	for k := range kvs {	//range can iterate over just the keys of the map
		fmt.Println("key:", k)
	}

	for i, c := range "go"{	//range on string iterate over UNICODE code points, first value is starting byte index of the rune and the second the rune itself
		fmt.Println(i, c)
	}
}