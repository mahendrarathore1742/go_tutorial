package main

import "fmt"

func main() {

	fmt.Println("welcome to loop section")

	day := []string{"Sun", "Mon", "Tue", "Wed", "Fri", "Sat", "thu"}

	fmt.Println(day)

	// for i,value:= range day{
	// 	fmt.Println(i,value)
	// }

	// for i:= 0 ; i< len(day);i++{
	// 	fmt.Println(i,day[i])
	// }

	for i := range day {
		fmt.Println(i, day[i])
	}

}
