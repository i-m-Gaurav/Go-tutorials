package main

import "fmt"


func printNum(num int){
	// note here so the num that is coming in this function, is indeed 54, but 
	// it coming by value okay, that's why we can change it here in this function,
	// or we can say, a copy of that number is coming in this function, and here in
	// this function we are modifying that copied number 54, that's why the original variable
	// in the main function still has the value of 54.
	// if we want to change that value in this function then we have to use the 
	// reference of that particular variable.
	num = 3
	fmt.Println("this is the printNum",num)
}

func printNum2(num *int){
	*num = 10
	fmt.Println("this is the inside the printNum2", *num)
}

 

func main(){

	num := 54

	// lets see the memory address of the num
	fmt.Println("address of num: " ,&num)

	printNum(num)
	fmt.Println("this number in the main:= ", num)


	printNum2(&num)
	fmt.Println("this is the changed number", num)

}




