package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our pizza app")
	fmt.Print("Please rate our pizza between 1 and 5: ")

	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input)

	// TrimSpace to trim \n to add 4 and 1
	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil { // it means err is not empty and has something
		fmt.Println(err)
	} else {
		fmt.Println("Added 1 to your rating,", numRating+1)
	}

}
