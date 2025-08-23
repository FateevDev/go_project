package main

func main() {
	cards := newDeck()

	cards.shuffle()

	cards.print()

	//hand, remainingDeck := deal(cards, 5)
	//
	//cards.print()
	//hand.print()
	//remainingDeck.print()
}
