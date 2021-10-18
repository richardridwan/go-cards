package main

func main() {
	hand, remainingDeck := deal(newDeck(), 2)
	hand.print()
	remainingDeck.print()
}

func newCard() string {
	return "Five of Diamonds"
}
