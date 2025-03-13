package main

import (
	"fmt"
	"os"
)

type Game struct {
	playerHandA []Card
	playerHandB []Card
	warStack    []Card
}

func (g *Game) transferWarStackToPlayer(playersHand *[]Card) {
	for len(g.warStack) > 0 {
		warStackTopCard := g.warStack[len(g.warStack)-1]
		g.warStack = g.warStack[:len(g.warStack)-1]                     //removeCardFromTopOfDeck
		*playersHand = append([]Card{warStackTopCard}, *playersHand...) //insertAtBottomOfDeck
	}
}

func (g *Game) transferCardsBetweenPlayers(loosersHand *[]Card, winnersHand *[]Card) {
	g.hasSomeoneWon()

	loosersTopCard := (*loosersHand)[len(*loosersHand)-1]
	*loosersHand = (*loosersHand)[:len(*loosersHand)-1]            //removeCardFromTopOfDeck
	*winnersHand = append([]Card{loosersTopCard}, *winnersHand...) //insertAtBottomOfDeck

	winnersTopCard := (*winnersHand)[len(*winnersHand)-1]
	*winnersHand = (*winnersHand)[:len(*winnersHand)-1]            //removeCardFromTopOfDeck
	*winnersHand = append([]Card{winnersTopCard}, *winnersHand...) //insertAtBottomOfDeck
}

func (g *Game) dealHands(cards []Card) {
	g.playerHandA = cards[:len(cards)/2]
	g.playerHandB = cards[len(cards)/2:]
}

func (g *Game) showPlayerHandStatus() {
	fmt.Printf("Player A hand status(%d): (bottom)", len(g.playerHandA))
	if len(g.playerHandA) == 0 {
		fmt.Println("Player A has no cards!")
	} else {
		for i := 0; i < len(g.playerHandA)-1; i++ {
			fmt.Printf("%s%s,", g.playerHandA[i].name, g.playerHandA[i].color)
		}
		fmt.Printf("%s%s(top)\n", g.playerHandA[len(g.playerHandA)-1].name, g.playerHandA[len(g.playerHandA)-1].color)
	}
	if len(g.playerHandB) == 0 {
		fmt.Println("Player B has no cards!")
	} else {
		fmt.Printf("Player B hand status(%d): (bottom)", len(g.playerHandB))
		for i := 0; i < len(g.playerHandB)-1; i++ {
			fmt.Printf("%s%s,", g.playerHandB[i].name, g.playerHandB[i].color)
		}
		fmt.Printf("%s%s(top)\n", g.playerHandB[len(g.playerHandB)-1].name, g.playerHandB[len(g.playerHandB)-1].color)
	}
}

func (g *Game) hasSomeoneWon() {
	if len(g.playerHandA) == 0 {
		fmt.Printf("Player B won the whole game!\n")
		os.Exit(0)
	}
	if len(g.playerHandB) == 0 {
		fmt.Printf("Player A won the whole game!\n")
		os.Exit(0)
	}
}

func (g *Game) makeTurn() {
	g.hasSomeoneWon()
	cardA := g.playerHandA[len(g.playerHandA)-1]
	cardB := g.playerHandB[len(g.playerHandB)-1]
	fmt.Printf("Player A card: %s%s, Player B card: %s%s\n", cardA.name, cardA.color, cardB.name, cardB.color)
	if cardA.value > cardB.value {
		fmt.Println("Player A won!")
		g.transferCardsBetweenPlayers(&g.playerHandB, &g.playerHandA)
		g.transferWarStackToPlayer(&g.playerHandA)
	} else if cardA.value < cardB.value {
		fmt.Println("Player B won!")
		g.transferCardsBetweenPlayers(&g.playerHandA, &g.playerHandB)
		g.transferWarStackToPlayer(&g.playerHandB)
	} else {
		fmt.Println("WAR!")
		for j := 0; j < 2; j++ {
			g.transferCardsBetweenPlayers(&g.playerHandA, &g.warStack)
			g.transferCardsBetweenPlayers(&g.playerHandB, &g.warStack)
		}
	}
}
