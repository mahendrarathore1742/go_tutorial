package main

import "fmt"

func main() {

	var result string

	loggincnt := 10

	if loggincnt < 10 {
		result = "Regular User"
	} else if loggincnt > 10 {
		result = "Watch out"
	} else {
		result = "Exactly 10 Login count "
	}

	fmt.Println(result)

	// create and check according  to  condition
	if Num := 112; Num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is not less than 10")
	}

}
