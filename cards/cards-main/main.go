package main

func main() {
	// syntax for making a slice
	cards := newDeck()
	hand, remainingCards := deal(cards, 5)

	// cards.print()
	hand.print()
	remainingCards.print()
}
