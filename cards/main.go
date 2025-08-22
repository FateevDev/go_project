package main

func main() {
	cards := newDeck()

	cards.saveToFile("cards.txt")

	cards = newDeckFromFile("cards.txt")

	//hand, remainingDeck := deal(cards, 5)
	//
	//cards.print()
	//hand.print()
	//remainingDeck.print()
}
