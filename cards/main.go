package main

func main() {
	cards := newDeck()

	cards.saveToFile("cards")

	cards = newDeckFromFile("cards")

	//hand, remainingDeck := deal(cards, 5)
	//
	//cards.print()
	//hand.print()
	//remainingDeck.print()
}
