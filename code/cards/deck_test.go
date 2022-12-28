package main

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	assert.Equal(t, 8, len(d))
	assert.Equal(t, d[0], "Ace of Diamonds")
	assert.Equal(t, d[len(d)-1], "Three of Clubs")
}

func TestDeal(t *testing.T) {
	d := newDeck()
	deckSize := len(d)
	handSize := 5
	hand, d := deal(d, handSize)
	assert.Equal(t, handSize, len(hand))
	assert.Equal(t, deckSize-handSize, len(d))
}

func TestSaveToFileAndDeckFromFile(t *testing.T) {
	path := "./files/testing/_deck-testing"
	expected := newDeck()
	err := expected.saveToFile(path)
	if err != nil {
		assert.Fail(t, "Couldn't save the file")
	}
	actual, err := newDeckFromFile(path)
	if err != nil {
		assert.Fail(t, "Couldn't read the file")
		cleanup(path)
	}
	assert.Equal(t, expected, actual)
	cleanup(path)
}

// Removed file with path
func cleanup(path string) {
	_ = os.Remove(path)
}

func TestShuffle(t *testing.T) {
	sh1, sh2 := newDeck(), newDeck()
	sh1.shuffle()
	sh2.shuffle()
	assert.NotEqual(t, sh1, sh2)
}
