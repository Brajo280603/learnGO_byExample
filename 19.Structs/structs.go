package main

import "fmt"


// In GO structs are typed collection of fields. They're useful for grouping data together to form records
type person struct{	// this person struct type has name and age fields
	name string
	age int
}

func newPerson(name string) *person {	// newPerson contructs a new person struct with the given name

	p := person{name: name}
	p.age = 42
	return &p	// we can safely return a pointer to local variable as a local variable will survive the scope of the function
}

func main(){
	fmt.Println(person{"Bob",  20}) // this syntax creates a new struct

	fmt.Println(person{name: "Alice", age : 30}) // the fields can be also named when initializing a struct

	fmt.Println(person{name: "Fred"}) // omitted fields will be zero-valued

	fmt.Println(&person{name: "Ann", age:40}) // an prefix yields a pointer to the struct

	fmt.Println(newPerson("Jon"))	//It's idiomatic to encapsulate new struct creation in contructor functions


	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)	//accesss struct fields with a dot 

	sp := &s
	fmt.Println(sp.age)	// dots can be also used with struct pointers - the pointers are automatically dereferenced.

	sp.age = 51	//	structs are mutable
	fmt.Println(sp.age)

	dog := struct {	// if a struct type is only used for a single value, we don't have to give it a name. the value can have an anonymous struct type.
		name string
		isGood bool
	}{
		"Rex",
		true,
	}

	fmt.Println(dog)
}