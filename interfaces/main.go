package main

import (
	"fmt"
)

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// reciever can remove variable if not used, see below
func (englishBot) getGreeting() string {
	// Very custom logic for generating an english greeting
	return "Hi there"
}

func (spanishBot) getGreeting() string {
	// Very custom logic for generating an spanish greeting
	return "Hola"
}