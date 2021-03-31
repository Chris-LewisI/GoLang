package main

import "fmt"

//Create a new type of 'deck' which is a slice of strings.
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	//replace variables you declare but don't use with underscores to bypass Go errors of unused variables.
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//d is a receiver, a copy of the deck we are using and is called in the function as 'd'
func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}
