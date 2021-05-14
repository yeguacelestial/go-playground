package main

import "fmt"

func arrays() {
	// Arrays
	var fruitArray [2]string

	// Assign values
	fruitArray[0] = "Apple"
	fruitArray[1] = "Orange"

	// Declare and assign
	fruitArr := [2]string{"Apple", "Orange"}

	fmt.Println(fruitArr)
}

func slices() {
	fruitSlice := []string{"Apple", "Orange"}
	fmt.Println(fruitSlice)
	fmt.Println(len(fruitSlice))
	fmt.Println(fruitSlice[1:2])
}

func main() {
	// arrays()
	slices()
}
