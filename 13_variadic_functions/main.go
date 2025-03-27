package main

import "fmt"


// this is the variadic function, so in short,
// if we want to pass unlimited number of arguments
// then we use the variadic function,
// one example of the variadic function is println 
// which is inbuilt, we can print anything with comma seperated
// now we have created this sum, that will take any number of
// int arguments,
// but it will take only int not anything
// for anything we will use the interface like this
// 
// func sum(nums ...interface{})

func sum(nums ...int) int {
	
	total := 0;
	for _ , a := range nums {
		total = total + a;
	}
	
	return total;
}


func main(){
	
	
	//Now note that here we are passing these numbers 
	// direcly, but let say we have a slice of numbers, 
	// then how will we pass it.
	// so for that we will use this syntax below
	
	numbers := []int{1,2,3,4,5}
	// this is the slice in which 5 numbers are present.
	// now pass it to that .
 	
	result2 := sum(numbers...)
	
	fmt.Println(result2)
	
	result := sum(3,4,5,6)
	fmt.Println(result);
	
}