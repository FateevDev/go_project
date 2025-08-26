package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	age       int
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
		ContactInfo: ContactInfo{
			email: "some@email.com",
			phone: "123-456-7890",
		},
	}

	//fmt.Println(person)
	//fmt.Printf("%#v", person)
	person.setAge(31)
	person.print()
}

func (p Person) fullName() string {
	return p.firstName + " " + p.lastName
}

func (p Person) setAge(age int) {
	p.age = age
}

func (p Person) print() {
	fmt.Println(p.fullName() + " is " + strconv.Itoa(p.age) + " years old.")
}
