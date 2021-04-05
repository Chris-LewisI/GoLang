package main

import "fmt"

type bot interface {
	//here we are saying that any type in our program that implements a function called "getGreeting" with a return type of "string" then that type is now part of the "bot" interface
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (eb englishBot) getGreeting() string {
	//if you're not using the receiver in the function than you can remove the reciever name entirely like so:
	//func (englishBot) getGreeting() string {
	// VERY custom logic for generating an English greeting
	return "Hi there!"
}

func (sb spanishBot) getGreeting() string {
	// VERY custom logic for generating a Spanish greeting
	return "Hola!"
}
