package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the value")

	input, _ := reader.ReadString('\n')

	fmt.Println("This is value\n", input)

	newRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println(err)
	} else {

		fmt.Println(newRating + 1)
	}

}
