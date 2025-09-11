package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	age       int
	Gender    string
	ContactInfo
}

type ContactInfo struct {
	email string
	phone string
}

func main() {
	person := Person{
		firstName: "John",
		lastName:  "Doe",
		age:       30,
		Gender:    "m",
		ContactInfo: ContactInfo{
			email: "some@email.com",
			phone: "123-456-7890",
		},
	}

	person.setAge(31)
	person.Gender = "s"

	fmt.Printf("%#v", person)

	person.print()
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName + " (" + p.Gender + ") "
}

func (p *Person) setAge(age int) {
	p.age = age
}

func (p Person) print() {
	fmt.Println(p.fullName() + " is " + strconv.Itoa(p.age) + " years old.")
}
