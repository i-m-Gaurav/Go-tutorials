package main

import "fmt"

const age = 43

func main(){

	const name  = "golang"

	fmt.Println(name,age);

	// if we have to do the multiple constant then do this

	const (
		port = 3000
		host = "localhost"
	)

	fmt.Println(port,host)
}