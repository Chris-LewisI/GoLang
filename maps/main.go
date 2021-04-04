package main

import "fmt"

func main() {
	//multiple ways to initialize a map

	// var colors map[string]string

	colors := make(map[string]string)
	colors["white"] = "#ffffff"

	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"green": "#00ff00",
	// 	"blue":  "#0000ff",
	// }

	delete(colors, "white")

	fmt.Println(colors)
}
