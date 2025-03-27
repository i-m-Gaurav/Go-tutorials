package main

import "fmt"

func main(){

	// we use range for iterating over data structures,

	nums := []int{1,2,3,4}

	// for i := 0 ; i< len(nums) ; i++ {
	// 	fmt.Println(i)
	// }

	sum := 0;


	// here this underscore is index, here we are skipping, we can put some variable
	// like i to get the index.
	for i , num := range nums {
		sum = sum + num;
		fmt.Println(sum , i)
	}


	// now lets iterate on teh  map

	maps := map[string]string{"name":"gaurav","address":"patna"}


	for k , v := range maps {
		fmt.Println(k , v)
		
	}

	// here if we are using single variable, then it will give only key from the
	// map, if we want both key and value , then we must use the both variable.

	for k := range maps {
		fmt.Println(k)
	}

	// we can iterate over the strings too.

	for i , c := range "Golang" {
		fmt.Println(i , c)

		fmt.Println(i,  string(c))
	}




}