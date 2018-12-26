package main

func main(){
	//Pattern
	cards := newDeck()
	cards.print()
}

func newCard() string {
	return "Ace of Spades"
}