package main

// LEXER takes care of syntax like ; and brackets and more automatically in go

import "fmt"

func main() {
	fmt.Println("Welcome to pointers")

	// var ptr *int
	// fmt.Println("Value of the pointer is: ", ptr)

	myNum := 23
	var ptr = &myNum
	fmt.Println("The address of pointer is:", ptr)
	fmt.Println("The value of pointer is:", *ptr)

	*ptr = *ptr + 2
	fmt.Println("The updated value of myNum is:", myNum)

}
