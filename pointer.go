package main

import "fmt"

func main() {

	fmt.Println("welcome to pointer application")

	var myPointer *int

	fmt.Println(myPointer)

	var value = 15
	var ptr = &value

	fmt.Println(*ptr)

	*ptr = 16
	fmt.Println(value)

}
