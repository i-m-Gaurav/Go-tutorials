package main

import "fmt"

func main(){

	//This is how we find the length of the array.

	// This array is of length 10
	var nums [10]int
	fmt.Println(len(nums))


	//This array is of length 20
	var num2 [20]int
	fmt.Println(len(num2))

	//now how we put the value in the array, 
	// like for 2nd index put some value like this

	num2[2] = 4
	fmt.Println(num2)
	// [0 0 4 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	// in the output we can see that 2nd index is now has value of 4

	// we can see the array by just printing the name of the array
	// here note one thing that when we initialized the array with int type, then all the
	// value initially is filled with 0

	// for string it will be empty string ""
	// for bool it will be false,
	// for the float it will be 0

	var vals [4]bool

	// let change value of the 3rd index

	vals[3] = true

	fmt.Println(vals)
	// [false false false true]

	var floats [5]float64

	fmt.Println(floats)


	// now lets initialize the value when we declare the array.

	fruits := [3]string{"tomato","grapes","apple"}
	fmt.Println(fruits)

	// lets make 2d array;

	num3 := [2][2] int {{2,3},{3,4}}

	fmt.Println(num3)


	// so in this case we have an array of the 2 x 3 so in the first matrix, for the 
	// third value there will be 0, because it always makes the 
	// square matrix.

	num4 := [2][3] int {{3,3},{2,5,6}}
	fmt.Println(num4)


	 
}