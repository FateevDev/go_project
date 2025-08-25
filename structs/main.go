package main

import "fmt"

type Person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	var person Person

	//person := Person{firstName: "John", lastName: "Doe", age: 30}
	//
	person.firstName = "Jane"
	person.age = 32

	fmt.Println(person)
	fmt.Printf("%#v", person)
}
