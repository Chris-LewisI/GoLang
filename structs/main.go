package main

import "fmt"

//creating a struct
type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	// contact   contactInfo
	contactInfo
}

func main() {
	//using a struct, there are multiple ways you can initialize one
	/*var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"*/
	//alex := person{"Alex", "Anderson"}
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	//there are also multiple ways to access the data in the struct
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	//examples of embedded structs
	var alex person
	alex.firstName = "Alex"
	alex.lastName = "Anderson"
	alex.contactInfo.email = "alexanderson@email.com"
	alex.contactInfo.zip = 12345

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contactInfo: contactInfo{
			email: "jimparty@email.com",
			zip:   12345,
		},
	}

	jim.updateName("jimmy") //doesn't work because of pointers, will fix
	jim.print()
	alex.print()
}

func (p person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
