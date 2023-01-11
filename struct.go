package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {

	mahendra := User{"mahendra singh", "mahendra@gmail.com", true, 23}

	// No inheritance in golang
	fmt.Println("welcome to struct")
	fmt.Println("-----------------")

	fmt.Println(mahendra)
	fmt.Println("-----------------")

	fmt.Printf("User Detail := %+v\n", mahendra)

	fmt.Println("-----------------")
	fmt.Println(mahendra.Name)

	fmt.Println("-----------------")
	fmt.Println(mahendra.Email)

	fmt.Println("-----------------")
	fmt.Println(mahendra.Status)

	fmt.Println("-----------------")
	fmt.Println(mahendra.Age)

	fmt.Printf("The name is %v and age is %v and Email is %v and Status is %v", mahendra.Name, mahendra.Email, mahendra.Age, mahendra.Status)
	fmt.Println("\n")

}
