package main

import "log"

func main() {
	cards, err := newDeckFromFile("./files/file.txt")
	if err != nil {
		log.Fatal("Error: ", err)
	}
	cards.shuffle()
	cards.print()
}
