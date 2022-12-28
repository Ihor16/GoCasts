package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"blue":  "#4b0000",
		"green": "#0bdd00",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Printf("The hex code for %s is %s\n", color, hex)
	}
}
