package main

import (
	"fmt"
	// "math/rand"
	// "math/rand"
	// "sync"

	// "sync"
	"time"
)

// so here in this example , we are using the channel, these are used to send the data
// from one go routine to another .

// here we are sending the number 5 from the main go routine to the process go routine.

// syntax for sending data is [ data <-  ]
// syntax for receiving data is [ <- data ]
// its pretty simple.

// func process(numChan chan int, wg *sync.WaitGroup){

// 	defer wg.Done()

// 	fmt.Println("Processing the number : ", <- numChan)

// }

func getNumQueue(numbers chan int) {

	// defer wg.Done()


	for num := range numbers {
		fmt.Println("getting number queue: ", num)
		time.Sleep(time.Second)
	}

}


// now we will get the number from this function to the main using the channel.

func sum(result chan int, num2 int ,num3  int){

	numResult := num2 + num3

	result <- numResult

}

func channelAsWaitGroup(done chan bool){


	defer func(){ done <- true}()


	fmt.Println("processing...")

}


func emailSender(emailChannel chan string, done chan bool){

	defer func(){ done <- true}()
	
	for email := range emailChannel {
		fmt.Println("sending email to : ", email)
		time.Sleep(time.Millisecond)
	}

	
}


func main(){


	result := make(chan int)

	go sum(result, 4, 2)

	res := <- result

	fmt.Println(res)

	// we can use the behavior of wait group using the channel,
	// as we know that the channels are blocking, so 
	// we can use this feature, to make sure our go routine is synchronized 

	// here is how wo do it.

	done := make(chan bool)

	go channelAsWaitGroup(done)

	<- done // now this will block the main thread,


	// now lets understand about buffered channel.
	// so here we are using the channel, and we are giving a 
	// size so that this cannel become buffered, and it will not give the
	// deadlock error, as previously we have seen,

	emailChannel := make(chan string, 100)


	go emailSender(emailChannel,done)

	// emailChannel <- "gaurav@gmail.com"
	// emailChannel <- "gautam@gmai.com"

	// fmt.Println(<-emailChannel)
	// fmt.Println(<-emailChannel)


	for i:= 0 ; i< 100 ; i++{
		emailChannel <- fmt.Sprintf("%d@gmail.com", i)
	}

	fmt.Println("done sending...");

	close(emailChannel)

	<-done // block


	


	

	// var wg sync.WaitGroup

	

	// messageChannel := make(chan string)

	// messageChannel <- "ping"

	// msg :=  <- messageChannel

	//Note : Channels are blocking

	// lets send number to the go routings using channels.

	// numbers := make(chan int)

	// // wg.Add(1)

	// go getNumQueue(numbers)

	// // for i:= 0 ; i<= 1000000 ; i++{
	// // 	fmt.Println(i)
	// // }

	// // here we are sending random number to the getNumberQueue routine. 
	// for {
	// 	 numbers <- rand.Intn(10000)
	// }

	// wg.Wait()

	// var wg sync.WaitGroup
	





	// numChan := make(chan int)

	



	// wg.Add(1)
	// go process(numChan, &wg)

	// numChan <- 5


	// wg.Wait()


	//here are are again waiting for some time, so that we can see the 
	// data being passed and processed. 
	// now lets use the wait group here.
	// time.Sleep(time.Second)



	channel1 := make(chan int)
	channel2 := make(chan string)

	go func(){
		channel1 <- 10
	}()

	go func(){
		channel2 <- "pong"
	}()

	for i:= 0 ; i< 2 ;i++ {
		select {
		case channel1value := <- channel1:
			fmt.Println("received Data from channel 1", channel1value)
		case channel2value := <- channel2:
			fmt.Println("data receied from channel 2: ", channel2value)
		}
	}





}