package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"

	fmt.Println("Fruit list is: ", fruitList)
	fmt.Println("Fruit length is: ", len(fruitList)) //4 -> takes the size assigned

	var vejList = [3]string{"Potato", "Beans", "Mushrooms"}
	fmt.Println("Vej List is: ", vejList)

}
