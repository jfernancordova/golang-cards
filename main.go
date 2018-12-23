package main

import "fmt"

func main(){
	card := newCard()

	fmt.Println(card) 
}

//go is always going to expect us 
//to label that type of data that 
//is being exchanged around our different
//functions inside of our program.
func newCard() string {
	return "Ace of Spades"
}