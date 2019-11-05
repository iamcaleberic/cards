package main

import "fmt"

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
	neo := person{
		firstName: "Neo",
		lastName:  "Anderson",
		contact: contactInfo{
			email:   "neo@matrix.system",
			zipCode: 101010,
		},
	}
	neo.updateName("Mr Neo")
	neo.print()

}

func (p *person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
