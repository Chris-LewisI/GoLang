//call package main so that it builds nicely
package main

//import necessary packages
import (
	"fmt"
)

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	//simple for loop going through the full range of numbers in the slice
	for i, n := range numbers {
		//check if even or odd and print result
		if n%2 == 0 {
			fmt.Println(i+1, " | ", n, " is even")
		} else {
			fmt.Println(i+1, " | ", n, " is odd")
		}
	}
}
