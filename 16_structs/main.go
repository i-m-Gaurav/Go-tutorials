package main

import (
	"fmt"
	"time"
)

 type Person struct {
	Name string
	Age int
 }

 // here is another struct, its like defining types like we do in the ts

 type Employee struct {
	Id	int
	FirstName	string
	LastName	string
	Salary		float64
 }

 type Order struct {
	orderId	int
	Name	string
	status	string
	createdAt	time.Time
 }



// this is the example of the embedded structs or we call it nested structs too.
 type Contact struct {
	Email	string
	Phone	int
 }

 type Developer struct {
	Name string
	Contact Contact
 }

 // now lets see Promotion in the nested structs,
 // here we don't define the type of the struct that we are embedding. see example

 type GameDetails struct{
	GameType string
	Requirements string
 }

 type Game struct {
	Id	int
	Name string
	rating float64
	GameDetails // here we will left it like this.
 }

 type rectangle struct {
	width float64
	height float64
 }


// here in this method the struct is passed by value,
// so we cannot modify the values of width and height, for that
// we will use the  pointer.
 func (r rectangle) Area() float64 {
	return r.width * r.height
 }

 // here is how we will use the pointer.
 // here we can change the values of the width and height.

 func (r *rectangle) Scale(factor float64) {
	r.width = r.width * factor
	r.height = r.height * factor
 }

 // in the above method we are modifying the values of the 
 // height and width of the rectangle by multiplying by a factor.






 

 // now lets initialize these.
// first lets see how we can initialize it with zero value





func main(){

	var p Person

	// so right now it has value of 0,
	fmt.Println("This is the person",p)

	// 2nd way lets initialize it with some value.

	emp := Employee {
		Id: 1,
		FirstName: "Gaurav",
		LastName: "Kumar",
		Salary: 22000.00,
	}

	fmt.Println("this is the employee: ", emp)
	// we can get specific values by using like this
	fmt.Println("First Name ",emp.FirstName)
	fmt.Println("Last Name ", emp.LastName)

	// now there is a thing called positional initialization,
	// here we directly put the values, but the order must be same,

	order := Order{101,"Laptop","received",time.Now()}
	fmt.Println("this is the order: ",order)


	//now we can create anonymous structs, like we don't have to define type
	// this is useful when we are using something for the single time.

	point := struct{ x , y int}{10,20}

	fmt.Println("This is the anonymous struct: ", point)

	// for the developer

	dev := Developer{
		Name: "Gurav the dev",
		Contact: Contact{
			Email: "indiagauravkumar@gmail.com",
			Phone: 1010101010,
		},
	}

	fmt.Println("this is the dev, with embedded struct ", dev )


	game := Game{
		Id: 101,
		Name: "PUBG",
		rating: 4.4,
		GameDetails: GameDetails{
			GameType: "battle royale",
			Requirements: "win 10, 16 GB ram",
		},
	}

	fmt.Println("this is the game, with promotion embedding ", game)


	rec := rectangle{
		width: 4,
		height: 5,
	}

	fmt.Println("Area is : ", rec.Area())

	rec.Scale(2)

	fmt.Println("New Area with modified scale of rectangle:  ", rec.Area())



}