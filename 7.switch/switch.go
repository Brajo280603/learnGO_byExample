package main

import ( //to import multiple wrap them with parenthesis ()
	"fmt"
	"time"
)

func main(){
	
	i := 2
	fmt.Println("Write ", i , " as ")

	switch i { //basic switch
	case 1:
		fmt.Println("one")
	
	case 2:
		fmt.Println("two")

	case 3:
		fmt.Println("three")
	}

	switch time.Now().Weekday() { 
	case time.Saturday, time.Sunday:	//multiple expressions can be separated by commas , 
		fmt.Println("It's the weekend 🥳")
	
	default: 	//default case
		fmt.Println("It's a weekday ☹️")
	}

	t := time.Now()
	switch { //basically a if/else logic
	case t.Hour() < 12:  
		fmt.Println("It's before Noon 🌄")

	default:
		fmt.Println("It's after Noon 🌆")
	}

	whatAmI := func(i interface{}) { 
		switch t := i.(type) { // a type-switch compares types instead of values. 
		case bool:
			fmt.Println("I am a bool")

		case int:
			fmt.Println("I am a int")
		
		default:
			fmt.Printf("I don't know the type LOL 🤣 : %T\n",t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("guess who's back")
	
}