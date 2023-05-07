package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	cardNum := 16
	newDeckFirstCard := "Ace of Spades"
	newDeckLastCard := "Four of Clubs"

	if len(d) != 16 {
		t.Errorf("Expected deck length of %v, but got %v", cardNum, len(d))
	}

	if d[0] != newDeckFirstCard {
		t.Errorf("Expected first card of %v, but got %v", newDeckFirstCard, d[0])
	}

	if d[len(d) - 1] != newDeckLastCard {
		t.Errorf("Expected last card of %v, but got %v", newDeckLastCard, d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	testFileName := "_deck_testing"
	os.Remove(testFileName)

	deck := newDeck()
	deck.saveToFile(testFileName)

	loadedDeck := newDeckFromFile(testFileName)

	if len(loadedDeck) != 16 {
		t.Errorf("Expected 16 cards in deck, got %v", len(loadedDeck))
	}

	os.Remove(testFileName)
}