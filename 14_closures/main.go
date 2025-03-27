package main

import "fmt"

// so the clousre
// let's understand it in pretty simple way
// lets say a function okay, what happen when we call it
// it just complete its execution, and clear all its
// variables state and memory,
// but in the closure, there is a outer scope varible, that
// is used inside another function, in the parent fuction,
// so the dependecny of the outter varible on the inner fuction always
// stays, so that even if we call the function, that variable with modified value
// will always be there,

func counter() func() int {
	count := 0

	return func() int {
		count++
		return count
	}

}



func main() {

	result := counter()

	fmt.Println(result())
	fmt.Println(result())
	fmt.Println(result())
	fmt.Println(result())
}
