package main

import (
	"fmt"
)

func main() {
	// creation method 1
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	// creation method 2
	// var colors map[string]string
	// colors = make(map[string]string)

	// creation method 3
	// colors := make(map[string]string)

	// add value
	// colors["white"] = "#ffffff"

	// delete value
	// delete(colors, "white")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
