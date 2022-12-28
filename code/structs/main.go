package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{
		firstName: "Alex",
		lastName:  "Anderson",
		contactInfo: contactInfo{
			email:   "alex@gmail.com",
			zipCode: 94000,
		},
	}
	alex.updateName("Tom")
	alex.print()

}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (p *person) updateName(firstName string) {
	p.firstName = firstName
}
