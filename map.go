package main

import "fmt"

func main() {

	fmt.Println("welcome to map")

	language := make(map[string]string)

	language["js"] = "Javascript"
	language["py"] = "python"
	language["java"] = "Java"

	fmt.Println(language)

	// how to delete form map;

	delete(language, "js")
	fmt.Println(language)

	for key, value := range language {
		fmt.Println(key, " :-", value)
	}

	//if you dont need key  only value so use replace key to this _:-

	for _, value := range language {
		fmt.Println(" :-", value)
	}

}
