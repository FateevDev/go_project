package main

import "fmt"

type Bot interface {
	getGreeting() string
}

type EnBot struct {
}

type SpBot struct {
}

func main() {
	en := EnBot{}
	sp := SpBot{}

	printGreeting(en)
	printGreeting(sp)
}

func printGreeting(b Bot) {
	fmt.Println(b.getGreeting())
}

func (EnBot) getGreeting() string {
	return "Hi there!"
}

func (SpBot) getGreeting() string {
	return "Hola!"
}
