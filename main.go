package main

import "fmt"

func main() {
	cards := []string{newCard(), "ace of diamonts"}
	cards = append(cards, "six of spades")
	fmt.Println(cards)

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "five of diamonds"
}
