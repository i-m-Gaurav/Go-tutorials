// here we will see the packages in the golang.
// main reason we use this, for code reusability, and better
// organization of the code.

package main

import (
	"github.com/fatih/color"

	"gaurav.com/tutorials/auth"
	"gaurav.com/tutorials/user"
)

func main() {

	auth.LoginWithCredentials("gaurav", "secret")
	session := auth.GetSession()

	println(session)

	user := user.User{
		Name:  "Gaurav",
		Email: "indiagauravkumar@gmail.com",
	}

	println("Name and Email is: ", user.Name, user.Email)
	color.Cyan(user.Name)

}
