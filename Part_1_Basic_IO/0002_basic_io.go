package main

import "fmt"

func main() {
	// We introduce the concept of a variable to which
	// we reserve a little box in our computer's memory
	// to hold the string which we are going to type
	// when prompted to provide our favorite food and
	// favorite drink
	fmt.Printf("What is your favorite food? ")

	// Create a string var for favorite food
	var favoriteFood string

	// Take input from the user
	fmt.Scanln(&favoriteFood)

	// Ask the user for their favorite drink
	fmt.Printf("What is your favorite drink? ")

	// Create a string var for favorite food
	var favoriteDrink string

	// Take input from the user
	fmt.Scanln(&favoriteDrink)

	// Here we  provide a response back based to the
	// console based on our two variables which we
	// inputted above
	fmt.Printf("I love %s and %s as well!", favoriteFood, favoriteDrink)
}
