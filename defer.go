package main

import "fmt"

// defer keyword is working like stack(LIFO)

func myDefer() {

	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}

}

func main() {

	// defer fmt.Println("4")
	// defer fmt.Println("3")
	// defer fmt.Println("2")
	// defer fmt.Println("1")
	// fmt.Println("Hello")

	myDefer()

}
