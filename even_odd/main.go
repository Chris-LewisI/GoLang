package main

import (
	"fmt"
)

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i, n := range numbers {
		if n%2 == 0 {
			fmt.Println(i+1, " | ", n, " is even")
		} else {
			fmt.Println(i+1, " | ", n, " is odd")
		}
	}
}
