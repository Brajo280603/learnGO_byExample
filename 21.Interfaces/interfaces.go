package  main

import (
	"fmt"
	"math"
)


//"Interfaces" are named collections of method signatures
type geometry interface {	//This is a basic interface for geometric shapes
	area() float64
	perim() float64
}


//here the "geometry" interface is implemented on "rect" and "circle" types
type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}


// To implement an interface in GO, all methods in the interface are need to be implemented
func (r rect) area() float64 {// this is "geometry" implemented on "rects"
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {	// the implementation for circles
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry){	//if a variable has an interface type , then all the methods that are in the named interface can be called. 
	// this is a generic measure function taking advantage of this work on any geometry
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func main(){
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	//The circle and rect struct types both implement the geometry interface so we can use instances of these structs as arguments to measure.
	measure(r)
	measure(c)
}

