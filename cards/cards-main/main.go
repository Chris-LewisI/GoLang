package main

import "fmt"

func main() {
	// syntax for making a slice
	// cards := newDeck()
	// hand, remainingCards := deal(cards, 5)

	// cards.print()
	// hand.print()
	// remainingCards.print()

	// greeting := "Hi there!"
	//[]byte turns a variable into its byte form
	// fmt.Println([]byte(greeting))

	cards := newDeck()
	fmt.Println(cards.toString())
}
