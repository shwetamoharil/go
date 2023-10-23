package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Craete a new type of 'deck'
// which is a slice of strings

type deck []string

// Receiver functions

func newDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suite := range cardSuites {
		for _, val := range cardValues {
			cards = append(cards, val+" of "+suite)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// here we dont use receiver since it would modify the actual cards copy, hence we pass it as an argument
func deal(d deck, handSize int) (deck, deck) {
	return d[0:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	file, err := os.ReadFile("my_cards")
	if err != nil {
		// option1 - log error and return call to newDeck()
		// option2 - log the error and entirely quit the program
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// convert byte slice to deck (slice of string)
	s := strings.Split(string(file), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		// swap
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
