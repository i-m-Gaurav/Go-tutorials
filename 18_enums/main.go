package main

import "fmt"

// so the enums in the go are like
// fixed set of named constants,
// imagine it like, traffic lights, it has only three colors, red, green, yellow
// it can't be other than that,

// since go does not have enum keyword like other language,
// here we use const and iota keyword to implements the enums.

// basic enumeration with iota


// here see , when we are using the iota, in each block of const, it start from 0 and 
// get incremented by one, so we can explicitly assign iota to each variable or
// we can just assign ones, and it will auto increment automatically for that
// const block
const (
	Red = iota
	Yellow = iota 
	Green = iota
)

const (
	R = iota
	Y 
	G
)

// and if we have to use some string, not iota we can do it like this

type OrderStatus string

const (
	Received OrderStatus  = "received"
	Confirmed			= "confirmed"
	Prepared 			= "prepared"
	Delivered 			= "delivered"
)


func ChangeOrderStatus(status OrderStatus){
	fmt.Println("Order status changed to : ", status)
}

// here in the above enum, it will work the same.


// here this iota is a predefined identifier, that starts from 0 and 
// get incremented by 1 for each next one.
// so here red will be 1 and yellow will be 2 and green will be 3.

// lets use that enums

func getTrafficAction(lightState int) string {
	switch lightState{
	case R:
			return "Stop!"
	case Y:
			return "Ready!"
	case G:
			return "Go!"
	default:
		return "invalid state!"
	}
	
}






func main(){

	fmt.Println("Traffic light states")
	fmt.Println("Red:", Red)
	fmt.Println("Yellow:", Yellow)
	fmt.Println("Green:", Green)

	fmt.Println(("\nActions:"))

	currentState := Yellow
	fmt.Printf("Current state (%d): %s\n",currentState, getTrafficAction(currentState))


	invalidState := 5
	fmt.Printf("Current state (%d): %s\n", invalidState,getTrafficAction(invalidState))


	ChangeOrderStatus(Prepared)




}