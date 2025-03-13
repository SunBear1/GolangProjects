package main

func main() {
	var numberOfCards = 8
	var numberOfColors = 4
	deck := buildDeck(numberOfCards, numberOfColors)
	deck.printDeck()
	game := Game{}

	deck.shuffleDeck()
	game.dealHands(deck.cards)

	for {
		game.showPlayerHandStatus()
		game.makeTurn()
	}
}
