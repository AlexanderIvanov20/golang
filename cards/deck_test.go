package main

import (
	"os"
	"testing"
)

const expectedNumberOfCards = 52

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != expectedNumberOfCards {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Two of Hearts" {
		t.Errorf("Expected first card of \"Two of Hearts\", but got \"%s\"", d[0])
	}

	if d[len(d)-1] != "Ace of Spades" {
		t.Errorf("Expected last card of \"Ace of Spades\", but got \"%s\"", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_deck_testing"

	os.Remove(filename)

	d := newDeck()
	d.saveToFile(filename)

	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(loadedDeck))
	}

	os.Remove(filename)
}
