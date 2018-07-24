package main

import (
	"fmt"
)

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// declaring and assigning using a struct method1
	// alex := person{"Alex", "Anderson"}

	// declaring and assigning using a struct method2
	// alex := person{firstName: "Alex", lastName: "Anderson"}

	// declaring and assigning using a struct method3
	// var alex person

	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"

	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)

	jim := person{
		firstName: "Jim",
		lastName:  "Party",
		contact: contactInfo{
			email:   "jim@gmail.com",
			zipCode: 9400,
		},
	}

	fmt.Printf("%+v", jim)
}
