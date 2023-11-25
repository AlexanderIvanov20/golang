package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	contact
}

func (p person) print() {
	fmt.Printf("%+v\n", p)
}

func (pp *person) updateName(firstName string) {
	(*pp).firstName = firstName
}
