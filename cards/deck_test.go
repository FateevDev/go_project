package main

import "testing"

func TestNewDeck(t *testing.T) {
	const deckLength = 52
	const firstCard = "Ace of Diamonds"
	const lastCard = "King of Clubs"

	d := newDeck()

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
