package main

import (
	"fmt"
	"strings"

	deck "github.com/ozyftw/gojack"
)


func cardToASCII(card deck.Card) string {

	var suit string
	switch card.Suit.String() {
	case "Spade":
		suit = "♠"
	case "Heart":
		suit = "♥"
	case "Diamond":
		suit = "♦"
	case "Club":
		suit = "♣"
	default:
		suit = "?"
	}
	

	var rank string
	switch card.Rank.String() {
	case "Ace":
		rank = "A"
	case "Two":
		rank = "2"
	case "Three":
		rank = "3"
	case "Four":
		rank = "4"
	case "Five":
		rank = "5"
	case "Six":
		rank = "6"
	case "Seven":
		rank = "7"
	case "Eight":
		rank = "8"
	case "Nine":
		rank = "9"
	case "Ten":
		rank = "10"
	case "Jack":
		rank = "J"
	case "Queen":
		rank = "Q"
	case "King":
		rank = "K"
	default:
		rank = card.Rank.String()
	}
	

	return fmt.Sprintf("┌─────┐\n│%s    │\n│  %s  │\n│    %s│\n└─────┘", 
		rank, suit, rank)
}


func displayHand(cards []deck.Card, title string) {
	fmt.Println(title)
	if len(cards) == 0 {
		return
	}
	

	cardLines := make([][]string, len(cards))
	for i, card := range cards {
		cardLines[i] = strings.Split(cardToASCII(card), "\n")
	}
	

	for line := 0; line < 5; line++ { 
		for i, cardLine := range cardLines {
			fmt.Print(cardLine[line])
			if i < len(cardLines)-1 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func main() {
 
    cards := deck.New(
        deck.Deck(6),    
        deck.Shuffle,    
    )
    
  
    playerHand := []deck.Card{cards[0], cards[2]}
    dealerHand := []deck.Card{cards[1], cards[3]}
    
     

    displayHand(playerHand, "Player Hand:")
    displayHand(dealerHand, "Dealer Hand:")
}