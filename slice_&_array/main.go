package main

import "fmt"

func main() {
	fmt.Println("This is slice programs")

	// declare a slice
	cards := []string{"Hi", "Hello", "Bye"}
	fmt.Println(cards)

	// adding new element to our slice
	// append does not modify the current slice, but instead create a new slice and we then assign it back to the card slice
	// Internally slice is just an array
	cards = append(cards, "New")
	fmt.Println(cards)

	// Iterate over slice
	for i, card := range cards {
		fmt.Println(i, card)
	}
}
