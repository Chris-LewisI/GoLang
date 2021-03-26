package main

import "fmt"

func main() {
	/*
		Go has a variety of different variables that can be used such as:
		bool: true or false
		string: "Hi!" "How's it going?"
		int: 0 -1000 99999
		float64: 10.000001
		and more!
	*/
	var card string = "Ace of Spades"
	card1 := "Ace of Spades" //this declaration of the variable card1 is the same as the declaration of the variable card
	//:= tells Go to figure out what kind of variable we are asking for by checking the value we provide it with after the equals sign
	card1 = "Five of Diamonds" // we do not need to use the colon again when changing the value of a predefined variable

	fmt.Println(card)
	fmt.Println(card1)
}
