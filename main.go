package main

func main() {
	//slicing[startingIndexIncluding:upToNotIncluding]
	// hand, remainingDeck := deal(newDeck(), 2)

	cards := newDeckFromFile("cards")

	cards.shuffle()
	cards.print()

	// cards.saveToFile("cards")

	// hand.print()
	// remainingDeck.print()
}
