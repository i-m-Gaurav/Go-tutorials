package main

import "fmt"

func main(){
	age := 19

	if age > 19 {
		fmt.Println("Person is adult");
	} else if age <= 22 {
		fmt.Println("He is not that big");
	}

	// declare the variable in the if block

	if age:= 20; age >=18 {
		fmt.Println("Person is an adult",age);
	}

	// go does not have ternary operator,
	// we have to use the if else for that.

}