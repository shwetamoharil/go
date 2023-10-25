package main

import "fmt"

func main() {
	// Create a map
	// key -string value-string
	colors := map[string]string{
		"red":    "#ff0000",
		"yellow": "#ffff00",
		"blue":   "#00001",
		"green":  "#00002",
	}

	fmt.Println(colors)

	var colors2 map[string]string
	fmt.Println(colors2)

	colors3 := make(map[string]string)
	fmt.Println(colors3)

	// Adding entry in map
	colors["white"] = "#00000"
	fmt.Println(colors)

	// Delete entry in map
	delete(colors, "green")
	fmt.Println(colors)

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is ", hex)
	}
}
