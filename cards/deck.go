package main

import (
	"fmt"
	"strings"
)

type deck []string

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func (d deck) append(card string) deck {
	return append(d, card)
}

func (d deck) toString() string {
	return strings.Join(d, ",")
}

func deal(d deck, handSize int) (deck, deck) {
	//newDeck := d[0:count]
	return d[:handSize], d[handSize:]
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Diamonds", "Hearts", "Spades", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}
