package main

import "fmt"

//creating a struct
type person struct {
	firstName string
	lastName  string
}

func main() {
	//using a struct, there are multiple ways you can initialize one
	/*var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"*/
	//alex := person{"Alex", "Anderson"}
	alex := person{firstName: "Alex", lastName: "Anderson"}

	//there are also multiple ways to access the data in the struct
	// fmt.Println(alex)
	fmt.Printf("%+v", alex)
}
