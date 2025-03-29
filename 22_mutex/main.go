package main

import (
	"fmt"
	"sync"
)

// race condition : multiple process trying to modify
// a same resource

// so here in this example , we have views 
// and the function inc() is in the for loop,
// and this function is trying to modify the views 500 times simultaneously 
// so it will create a race condition.

// that's where mutex comes,

// what it actually does that, when a variable or resource is being 
// modified then it locks that so that only one go routine can
// modify that.


type post struct {
	views int
	// lets create a mutex here, it also comes in the sync
	mu sync.Mutex
}


func(p *post ) inc(wg *sync.WaitGroup) {

	// defer wg.Done()

	defer func(){
		p.mu.Unlock()
		wg.Done()
	}()

	//here before the modification, we will lock and after
	// modification we will unlock it.

	p.mu.Lock()
	p.views = p.views + 1
	// p.mu.Unlock()



	// now note here we have moved the unlock in the defer
	// why, let say if some error comes during the 
	// modification logic then it won't come to the unlock
	// and hold the resource forever,
	// so we have put the unlock in the defer so that it must run after
	// the function execution.


	// there is one good practice that, we should use the lock
	// only for that line where modification is exactly happening.

	
}

func main() {

	var wg sync.WaitGroup


	myPost := post{views : 0}


	for i := 0 ; i< 500 ;i++{

		wg.Add(1)
		go myPost.inc(&wg)
	}
	
	wg.Wait()
	fmt.Println(myPost.views)
	
	

	

}