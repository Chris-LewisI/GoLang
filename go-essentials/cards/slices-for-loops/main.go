package main

import "fmt"

func main() {
	// syntax for making a slice
	cards := []string{newCard(), "Ace of Diamonds"}
	cards = append(cards, "Six of Spades")

	// i = index of this element in the array
	// card = current card being iterated over
	// range cards = takes "cards" slice and loops over the range of it
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "Five of Diamonds"
}
