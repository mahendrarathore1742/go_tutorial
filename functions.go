package main

import "fmt"

func hello() {
	fmt.Println("hi I am come from hello function")
}

func add(a int, b int) int {
	return a + b
}

func addStr(a string, b string) string {
	return a + " " + b
}

func PrintArray(array []string) {
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
}

// spacial function if lots of values is comming as a argument but we dont know much  how many value is comming
func Proadder(value ...int) int {

	total := 0

	for _, val := range value {

		total += val
	}

	return total
}

// if you want to return 2 type  diffrent values ex - (int, string)

func twotype(a int) (int, string) {

	return a, "hiiii"
}

func main() {

	var a int = 10
	var b int = 20

	fmt.Println("hello from GoLang")
	fmt.Println(add(a, b))

	fmt.Println(addStr("mahendra", "Singh"))

	lang := []string{"js", "py", "go", "rust"}

	PrintArray(lang)

	fmt.Println(Proadder(1, 2, 3, 4, 5))

	fmt.Println(twotype(5))

}
