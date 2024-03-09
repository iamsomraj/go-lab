package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type deck []string

func newCards() deck {
	cards := deck{}
	cardGroups := []string{"Spades", "Clubs", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jocker", "Queen", "King"}

	for _, group := range cardGroups {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+group)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	hand := d[:handSize]
	remainingCards := d[handSize:]
	return hand, remainingCards
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func toBytes(s string) []byte {
	return []byte(s)
}


func (d deck) writeToFile(fileName string) error {
	cards := d.toString()
	cardBytes := toBytes(cards)
	return ioutil.WriteFile(fileName, cardBytes, 0666) // This is deprecated
}
