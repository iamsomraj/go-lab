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
	somraj := person{
		firstName: "Somraj",
		lastName:  "Mukherjee",
		contact: contactInfo{
			email:   "iamsomraj@gmail.com",
			zipCode: 700001,
		},
	}
	somraj.updateName()
	println(&somraj)
}

func (p *person) updateName() {
	(*p).firstName = "Souradip"
	(*p).lastName = "Ganguly"
}

func println(p *person) {
	fmt.Printf("%+v", *p)
}
