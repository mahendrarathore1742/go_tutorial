package main

import "fmt"

type User struct {
	Name   string
	Email  string
	status bool
	age    int
}

// create method for access user status and create method like this

func (u User) Status() {
	fmt.Println("Is User active: ", u.status)
}

// modifie the user Email
func (u User) NewMail() {
	u.Email = "msr@gmail.com"
	fmt.Println("Is User active: ", u.Email)
}

func main() {

	mahendra := User{"mahendra singh", "mahendra@gmail.com", true, 23}
	mahendra.Status()
	mahendra.NewMail()
	fmt.Println(mahendra.Email)
}
