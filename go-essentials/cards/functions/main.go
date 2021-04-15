package main

import "fmt"

func main() {
	card := newCard()

	fmt.Println(card)
}

//declare function and identify the return type
func newCard() string {
	return "Five of Diamonds"
}
