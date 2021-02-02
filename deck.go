package main

import "fmt"

//create of new type of 'deck'
// it is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"spades", "diamonds", "hearts", "clubs"}

	cardValues := []string{"ace", "two", "three", "four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+"of"+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
