package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

// Create a new type of "deck" which is a slice of strings.
type deck []string

func newDeck() deck {
	cards := deck{}

	values := []string{
		"Two",
		"Three",
		"Four",
		"Five",
		"Six",
		"Seven",
		"Eight",
		"Nine",
		"Ten",
		"Jack",
		"Queen",
		"King",
		"Ace",
	}
	suits := []string{
		"Hearts",
		"Diamonds",
		"Clubs",
		"Spades",
	}

	for _, value := range values {
		for _, suit := range suits {
			cards = append(cards, fmt.Sprintf("%s of %s", value, suit))
		}
	}

	return cards
}

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return deck(strings.Split(string(bs), ";"))
}

func (d *deck) deal(handSize int) deck {
	cards := *d

	hand := cards[:handSize]
	*d = cards[handSize:]

	return hand
}

func (d deck) toString() string {
	return strings.Join(d, ";")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) shuffle() {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	for i, cards := range d {
		randIndex := r.Intn(len(cards) - 1)

		d[i], d[randIndex] = d[randIndex], d[i]
	}
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
