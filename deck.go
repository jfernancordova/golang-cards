package main

import "fmt"

// Create a new type of deck
// which is a slice of strings

type deck []string

//Receivers sets up methods on variables that we create
//d by convenction
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}