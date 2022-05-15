package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to maths in golang")

	// var num1 int = 2
	// var num2 float64 = 4
	// fmt.Println("The sum is: ", num1+int(num2)) -> not good bcz sum dsnt consider decimal.

	// random number
	// rand.Seed(time.Now().UnixNano()) // -> random function of maths package
	// fmt.Println(rand.Intn(5) + 1)

	// random from crypto	-> more accurate

	myRandomNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)

}
