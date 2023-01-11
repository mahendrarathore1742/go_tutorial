package main

import "fmt"

func main() {

	var my_array [4]int

	my_array[0] = 1

	my_array[1] = 1

	my_array[2] = 1

	my_array[3] = 1

	fmt.Println(my_array)

	var my_array1 = [4]int{1, 2, 3, 4}

	fmt.Println(len(my_array1))

}
