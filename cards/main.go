package main

func main() {
	FILE_NAME := "my_cards.txt"

	cards := newDeckFromFile(FILE_NAME)
	cards.print()
}
