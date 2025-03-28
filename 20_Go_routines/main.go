package main

import (
	"fmt"
	"time"
)

// what the hell is go routine.
// its a function that can run concurrently alongside other functions.
// basically its an incredible lightweight , independent executing task.

// in the main function all the work can do done normally, but if there are some
// blocking tasks , like some loops, then we use the go routine , so that they don't block the main
// thread , so by this, main function can run by its own pace, and till then go routines will complete their execution, independently.

func printNum(id int){

fmt.Println("doing task for : ", id)
	
}

func main(){

		for i:= 0 ; i<=10 ;i++{

			// we just have to put go keyword to run the go routine.
			go printNum(i)
		}
		

		time.Sleep(time.Second * 2)

		

}