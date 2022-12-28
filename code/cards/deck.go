package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// A type deck that is a slice of string
type deck []string

const DELIM = ","

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func newDeck() deck {
	cards := deck{}
	names := []string{"Diamonds", "Clubs"}
	values := []string{"Ace", "One", "Two", "Three"}
	for _, name := range names {
		for _, value := range values {
			cards = append(cards, value+" of "+name)
		}
	}
	return cards
}

// Returns (1) a new deck of size n and (2) old deck containing len-n cards
func deal(d deck, n int) (deck, deck) {
	return d[:n], d[n:]
}

// Joins deck elements by DELIM
func (d deck) toString() string {
	return strings.Join(d, DELIM)
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0644)
}

func newDeckFromFile(filename string) (deck, error) {
	content, err := ioutil.ReadFile(filename)
	return strings.Split(string(content), DELIM), err
}

func (d deck) shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d), func(i, j int) {
		d[i], d[j] = d[j], d[i]
	})
}
