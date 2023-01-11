package main

import "fmt"

func main() {

	// var fruitList = []string{"mango","banana"};

	// fmt.Printf("this is Type of %T \n",fruitList)

	// fruitList = append(fruitList,"Apple");

	// fmt.Println(fruitList[1:])

	// fruitList = append(fruitList[1:]);
	// fmt.Println(fruitList)

	// fruitList = append(fruitList[:3]);

	// fmt.Println(fruitList)

	// var highScore = make([]int,4);

	// highScore[0] = 0;
	// highScore[1] = 1;
	// highScore[2] = 2;
	// highScore[3] = 3;

	// // we can put more data because the size is only 4 so that why we will use append method

	// highScore = append(highScore,666,4,5,23,23,4,23);

	// fmt.Println(highScore)

	// sort.Ints(highScore)
	// fmt.Println(highScore)

	// how to delete element from slices

	var course = []string{"React js", "Javascript", "python", "c++"}
	fmt.Println(course)

	var index int = 2

	course = append(course[:index], course[index+1:]...)

	fmt.Println(course)

}
