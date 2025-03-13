package main

import (
	"fmt"
	"math/rand"
)

type Card struct {
	name  string
	value int
	color string
}

type Deck struct {
	cards          []Card
	numberOfCards  int
	numberOfColors int
}

const (
	Clubs    = "♣️"
	Spades   = "♠️"
	Hearts   = "♥️"
	Diamonds = "♦️"
)

var (
	Ace   = Card{"Ace", 13, ""}
	King  = Card{"King", 12, ""}
	Queen = Card{"Queen", 11, ""}
	Jack  = Card{"Jack", 10, ""}
	Ten   = Card{"10", 9, ""}
	Nine  = Card{"9", 8, ""}
	Eight = Card{"8", 7, ""}
	Seven = Card{"7", 6, ""}
	Six   = Card{"6", 5, ""}
	Five  = Card{"5", 4, ""}
	Four  = Card{"4", 3, ""}
	Three = Card{"3", 2, ""}
	Two   = Card{"2", 1, ""}

	CardColors = [4]string{Spades, Hearts, Diamonds, Clubs}
	CardValues = [13]Card{Ace, King, Queen, Jack, Ten, Nine, Eight, Seven, Six, Five, Four, Three, Two}
)

func buildDeck(numberOfCards int, numberOfColors int) Deck {
	var cards []Card
	for i := 0; i < numberOfColors; i++ {
		color := CardColors[i]
		for j := 0; j < numberOfCards; j++ {
			card := CardValues[j]
			card.color = color
			cards = append(cards, card)
		}
	}
	return Deck{cards: cards, numberOfCards: numberOfCards, numberOfColors: numberOfColors}
}

func (d *Deck) printDeck() {
	for i := 0; i < d.numberOfColors; i++ {
		for j := 0; j < d.numberOfCards; j++ {
			idx := j + d.numberOfCards*i
			fmt.Printf("%s%s,", d.cards[idx].name, d.cards[idx].color)
		}
		fmt.Println()
	}
	fmt.Println()
}

func (d *Deck) shuffleDeck() {
	for i := range d.cards {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}
