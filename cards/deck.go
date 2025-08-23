package main

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

type deck []string

// receiver of a method is the object on which the method is called
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

func (d deck) saveToFile(filename string) {
	err := os.WriteFile(filename+".txt", []byte(d.toString()), 0644)

	if err != nil {
		fmt.Println("Error writing deck file:", err)
		os.Exit(1)
	}
}

func (d deck) shuffle() {
	rand.Shuffle(len(d), func(i, newPosition int) { d[i], d[newPosition] = d[newPosition], d[i] })
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

func newDeckFromFile(filename string) deck {
	bytes, err := os.ReadFile(filename + ".txt")

	if err != nil {
		fmt.Println("Error reading deck file:", err)
		os.Exit(1)
	}

	if len(bytes) == 0 {
		fmt.Println("Deck file is empty")
		os.Exit(1)
	}

	deckSlice := strings.Split(string(bytes), ",")
	deck := deck(deckSlice)

	return deck
}
