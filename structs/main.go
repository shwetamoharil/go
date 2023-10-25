package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

// Defining struct
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// 1. declaring struct
	// testPerson := person{"Firstname", "Lastname"}
	// fmt.Println(testPerson)

	// // 2. declaring strcuct
	// testPerson2 := person{firstName: "Firstname2", lastName: "Lastname2"}
	// fmt.Println(testPerson2)

	// // 3. declaring struct
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Jone"
	// fmt.Println(alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Parton",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 12345,
		},
	}
	// jimPointer := &jim
	// jimPointer.updateName("Jimmy")
	// jim.printPerson()

	jim.updateName("Jimmy")
	jim.printPerson()

	// The below code will actually modify the mySlice
	mySlice := []string{"hi", "there", "how", "are", "you"}
	updateSlice(mySlice)
	fmt.Println(mySlice)
}

// creating receiver function
func (p person) printPerson() {
	fmt.Printf("%+v", p)
}

// This will not update the person jim
// Go is pass by value reference lang
// Go makes copy of struct and then updates that particular obj
// To make this work we will use pointer
// func (p person) updateName(newFirstName string) {
// 	p.firstName = newFirstName
// }

// *pointer means value of type pointer
func (p *person) updateName(newFirstName string) {
	(*p).firstName = newFirstName
}

func updateSlice(s []string) {
	s[0] = "bye"
}
