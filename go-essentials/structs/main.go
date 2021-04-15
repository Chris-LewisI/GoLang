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
	//turn value into address witrh "&value"
	// jimPointer := &jim //& is an operator that points at the memory address of the variable next to it
	// jimPointer.updateName("Jimmy")
	jim.updateName("Jimmy") //this is a shortcut in Go that allows us to bypass the pointer setup if the type matches the pointer type in the function
	jim.print()

	alex.print()
}

//turn address into value with "*address"
func (pointerToPerson *person) updateName(newFirstName string) { //*person is different from *pointerToPerson
	//*[type] this is a type description that means we are working  with a pointer to a person
	//*[pointerToPerson] this is an operator that gives us access to the value we wnt to manipulate or access
	(*pointerToPerson).firstName = newFirstName //* operator asks for the value that is stored at that memory address
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}
