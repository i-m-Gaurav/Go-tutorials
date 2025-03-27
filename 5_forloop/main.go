package main

import "fmt"

func main(){
	// go has only one loop for loop , we can implement the while using 
	// for keyword here is how.

	//while loop

	// i := 1

	// for i <= 6 {
	// 	fmt.Println(i);
	// 	i++;
	// }

	//for loop

	for i := 0 ; i< 6 ; i++ {
		
		if(i == 3){
			continue // skip the iteration
		}
		fmt.Println(i);
	}

	// there is one more thign called range.

	// here i start from 0 by default, and incrementation work automatically
	

	for i := range 11 {
		fmt.Println(i);
	}
}