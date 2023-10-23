package main

import "fmt"

func main() {
	fmt.Println("New deck")
	cards := newDeck()
	cards.shuffle()
	cards.print()
	// cards := newDeckFromFile("my_cards")
	// cards.print()
	// Type conversion
	// gretting := "Hi there"
	// fmt.Println([]byte(gretting))

	// cards := newDeck()
	// cards.saveToFile("my_cards")
	// cards.print()
	// hand, remaningDeck := deal(cards, 5)
	// hand.print()
	// remaningDeck.print()
}
