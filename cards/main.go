package main

func main() {
	cards := newDeck()

	hand, remainingDeck := deal(cards, 5)

	cards.print()
	hand.print()
	remainingDeck.print()
}
