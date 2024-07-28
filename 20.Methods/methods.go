package main 

import "fmt"

//GO supports methods defined on struct type

type rect struct {
	width, height int
}

func (r *rect) area() int { // this area method has a receiver type of *rect
	return r.width * r.height
}

func (r rect) perim() int { // methods can be defined for either pointer or value receiver types. this is an example of a value receiver
	return 2*r.width + 2*r.height
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("area : ", r.area())	//here we call the two methods defined for our structs

	fmt.Println("perim : ",r.perim())


	//GO automatically handles conversion between values and pointers for method calls. A pointer receiver type can be used to avoid copying on method calls or to allow the method to mutate the receiving struct.
	rp := &r	

	fmt.Println("area : ", rp.area())
	fmt.Println("perim : ", rp.perim())
}