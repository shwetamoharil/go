package main

import "fmt"

type bot interface {
	getGreeting() string
}
type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	// sb := spanishBot{}

	printGreeting(eb)
	// printSpanishGreeting(sb)
}

func (englishBot) getGreeting() string {
	// very custom logic to return greeting in english
	// Here in we are not using eb, hence we can eliminate and just specify the receiver type
	return "Hello There!"
}

func (sb spanishBot) getGreeting() string {
	// very custom logic to return greeting in spanish
	return "Hola!"
}

// Both the printGreeting have a similar logic and can be resused
// The only diff between them is the different argument type (englishBot, spanishBot)

// func printGreeting(eb englishBot) {
// 	fmt.Println(eb.getGreeting())
// }

// func printSpanishGreeting(sb spanishBot) {
// 	fmt.Println(sb.getGreeting())
// }

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}
