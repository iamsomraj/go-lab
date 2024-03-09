package main

func main() {
	cards := newCards()
	hand, remainingCards := deal(cards, 5)
	hand.print()
	remainingCards.print()
}
