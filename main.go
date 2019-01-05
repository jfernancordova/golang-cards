package main

func main(){
	cards := newDeck()
	cards.saveToFile("myCards")
}

func newCard() string {
	return "Ace of Spades"
}