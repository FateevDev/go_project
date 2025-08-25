package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	// Arrange
	const deckLength = 16
	const firstCard = "Ace of Diamonds"
	const lastCard = "Four of Clubs"

	// Act
	d := newDeck()

	// Assert
	if len(d) != deckLength {
		t.Errorf("Expected deck length of %v, but got %v", deckLength, len(d))
	}

	if d[0] != firstCard {
		t.Errorf("Expected first card to be %v, but got %v", firstCard, d[0])
	}

	if d[deckLength-1] != lastCard {
		t.Errorf("Expected last card to be %v, but got %v", lastCard, d[deckLength-1])
	}
}

func TestSaveAndLoadFromFile(t *testing.T) {
	// Arrange
	const deckLength = 16
	const fileName = "_decktesting"

	// Act
	d := newDeck()

	deleteTestFile(fileName)

	d.saveToFile(fileName)
	deckFromFile := newDeckFromFile(fileName)

	// Assert
	if len(deckFromFile) != deckLength {
		t.Errorf("Expected deck length of %v, but got %v", deckLength, len(deckFromFile))
	}

	deleteTestFile(fileName)
}

func deleteTestFile(fileName string) {
	err := os.Remove(fileName + ".txt")

	if err != nil {
		return
	}
}
