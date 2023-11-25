package main

func main() {
	// cards := newDeck()
	// cards.saveToFile("my_cards.txt")

	cards := newDeckFromFile("my_cards.txt")
	cards.print()
	cards.shuffle()
	cards.print()
}
