package main

import (
	"fmt"
	"time"
)

func main(){
	i:=1

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	default:
		fmt.Println("Other value")
	}

	// we have one more type of switch, 
	// that is with multiple condition.

	switch time.Now().Weekday() {

	case time.Saturday, time.Sunday :
		fmt.Println("It's weekend")
	default:
		fmt.Printf("Its tuesday");

	}

	// type switch in go

	whichType := func(i interface{}) {

		switch i.(type) {
		case int:
			fmt.Println("Its an integer")
		case float64:
			fmt.Println("Its a float value")
		case bool:
			fmt.Println("Its a bool")
		default:
			fmt.Println("Other value")
		}
	
	}

	whichType(3.14);


	  
}