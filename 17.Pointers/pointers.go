package main
import "fmt"


// go supports "pointer" allowing to pass references to values and records within this program
func zeroval(ival int){
	ival = 0 //here ival is setting the value its passed to zero
}

func zeroptr(iptr *int){
	*iptr = 0 //here iptr is setting the reference passed to zero
}

func main() {
	i := 1	
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i) // the &i syntax gives the memory address of i , (a pointer to i)
	fmt.Println("zeroptr:",i)

	fmt.Println("pointer:",  &i)// here the pointer is printed (the memory address for i)

	
}