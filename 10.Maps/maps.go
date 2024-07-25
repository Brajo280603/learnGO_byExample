package main

import (
	"fmt"
	"maps"
)

func main(){
	m := make(map[string]int) //use this to create an empty map (same as dict or hashes in other languages)

	m["k1"] = 7		//set key/value pairs
	m["k2"] = 13

	fmt.Println("map:", m)

	v1 := m["k1"]	//accessing a value with a key
	fmt.Println("v1:", v1)

	v3 := m["k3"]	//if no key exists it returns zero value (0)
	fmt.Println("v3:", v3)

	fmt.Println("len:", len(m))	//use len() to get the numbers of key/value pairs

	delete(m, "k2")		//use delete to remove a specific value with a key
	fmt.Println("map:", m)

	clear(m)	// use clear to clear the whole map
	fmt.Println("map:", m)

	_, prs := m["k2"]	//the optional second return value (prs in this case) shows the if the key was present in the map, this can be used to disambiguate between missing keys and keys with zero values like 0 or "". here because we didn't need the value itself , so ignored it with _ (blank) indentifier
	fmt.Println("prs:", prs)


	n := map[string]int{"foo": 1, "bar": 2} 	//declare and initialize a new map in the same line
	fmt.Println("map:", n)

	n2 := map[string]int{"foo": 1, "bar": 2}
	if maps.Equal(n, n2){		//same function as slices, checks if both maps are same or not
		fmt.Println("n == n2")
	}

	// NOTE : maps appear in the form [k:v k:v] when printed with fmt.Println
}