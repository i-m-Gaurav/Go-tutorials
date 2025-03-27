package main

import "fmt"

// lets create a simple function to add numbers

// its not returning any thing we are just printing the value 
// inside the function
func add(a int, b int)  {
	fmt.Println(a + b)
}


// but in this function it is returning something, so we have to give 
// the return type.
func add1(a int, b int) int {
	return a + b;
}

// on the above code as you can see it is returning a single value,
// but if we are returning multiple values then we have to do like this


// so this function only returning something, not expecting anything.
func getSome() (int,int,string) {
	return 1 ,2 ,"Gaurav"
}


// since functions in the golang is first class citizen, 
// so we can pass it inside a function, and we can also
// assign the function to a variable too.





func printNum(p func(a int) int) {
	p(1)
}

func processIt() func(a int) int {
	return func(a int) int {
		return 5;
	}
}


func main(){

	fmt.Println(getSome())

	fn1 := processIt();
	println(fn1)

	

	fn := func( a int) int {
		return 2
	}

	printNum(fn)

	a , _ , c := getSome();
	fmt.Println(a,c)




	// fmt.Println(printNum(fn(5)))


	add(4 , 5)
	fmt.Println(add1(3,3))
}