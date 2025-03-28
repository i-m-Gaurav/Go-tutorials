package main

import "fmt"

// now here are are using  the generics in the golang.
// so as you can see in the below function, here we are giving a type 
// in the function parameter that it will only take the slice of  int
// but what if we can pass any type of slice then how will it work
// so for that we will use the generics





func printSlice(items []int){

	for index , item := range items {
		fmt.Println(index ,item)
	}

}

// now we are going to define this function using the generic type, not only int, or float , or string.

// now this function will accept any type of slice, and iterate over it.
// but this is still somewhat not good, because we are using the any type
func printSlice2[T any]( items []T ){

	fmt.Println("with generic type : ")
	for _, item := range items {
		fmt.Println(item);
	}
}


// here i used the empty interface {} which is similar to the any type in the golang.
func printSlice3[T interface{}]( items []T ){

	fmt.Println("with generic type with empty interface : ")
	for _, item := range items {
		fmt.Println(item);
	}
}


// now how can we limit what certain type our function should take, for this
// we can use the pipe operator with the types like this.

func printSlice4[T int | string ] (items []T){
	fmt.Println("This will only take the int and string : ")
	for _ , item := range items {
		fmt.Println(item)
	}
}


// now we are here using two certain types, we can add more into this,
// but in go we have a better way for this,
// there is a thing called comparable, if we use that , it has some comparable types in that
// and that all will be used in our functions, we can't put anything extra, and its
// pretty safe. so we can use that , like this.

func printSlice5[T comparable] (items []T) {
	fmt.Println("This is with the comparable: ")
	for _ , item := range items {
		fmt.Println(item)
	}
}






func main(){


	nums := []int {1,2,3,4,5,6}

	fruits := []string{"apple","grapes","banana"}

	vals := []bool{true,true,false}



	printSlice(nums)
	printSlice2(nums)
	printSlice2(fruits)

	printSlice3(fruits)

	fmt.Println("This will with int and string only")

	printSlice4(nums)
	printSlice4(fruits)

	// this will give error
	// bool does not satisfy int | string (bool missing in int | string)
	// printSlice4(vals)

	// this is with the comparable,
	printSlice5(nums)

	// here we can pass the slice of bools too.

	printSlice5(vals)



}