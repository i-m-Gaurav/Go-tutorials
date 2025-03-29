package main

import (
	"fmt"
	"sync"
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


func worker(id int,wg *sync.WaitGroup){
	// here we will call the done method form the wait group with the defer
	// keyword,

	defer wg.Done()
	// what does defer do it will call the done no matter what how the function ended.

	fmt.Printf("worker %d: Starting\n", id)

	// do some waiting
	time.Sleep((time.Second))
	fmt.Printf("Worker %d Finished\n", id)
}

func main(){

		for i:= 0 ; i<=10 ;i++{

			// we just have to put go keyword to run the go routine.
			// now this function will run on separate thread,

		
			go printNum(i)
		}




		// lets understand one thing here, here as you can see we are blocking the main function
		// for two seconds, but this is for test only , so that we can see what actually that go routing is doing,
		// if we remove this sleep, and run the program we won't be able to see the output of the go routine, 
		// because, if main function complete its execution and then go routine complete its execution, there is no way 
		// that it will report that i have completed , here is the output, so we will see nothing in the terminal.


		// so for that we have wait group in go routines,
		// the main task for these are, the main function will wait for the go routines, to complete its execution, then it 
		// will complete its own execution, in this way we can see the output of the go routines.

		// here is how we use it, 

		var wg sync.WaitGroup

		workerNumber := 3
		fmt.Printf("Main: Launching %d workers...\n",workerNumber)

		for i := 1; i<= workerNumber ; i++{
			// here we will increment the counter, before launching the go routine
			// lets do that
			wg.Add(1)
			// so we did it, we are telling that we have one go routing to done.

			// now launch the routine
			go worker(i, &wg)
		}

		fmt.Println("Main: all worker launched. waiting fo them to finish...")


		wg.Wait()
		// this wait will wait for all the go routine to complete their work

		fmt.Println("main : all worker have finished . done")





		// time.Sleep(time.Second * 2)

		

}