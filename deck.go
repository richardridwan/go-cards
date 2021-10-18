package main

import "fmt"

// create a new type of deck
// which is a slice of string

type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, cardSuit := range cardSuits {
		for _, cardValue := range cardValues {
			card := cardValue + " of " + cardSuit
			cards = append(cards, card)
		}
	}

	return cards
}

func deal(d deck, n int) (deck, deck) {
	return d[:n], d[n:]
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i+1, card)
	}
}
