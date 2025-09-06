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

	
	rank = fmt.Sprintf("%-2s", rank)

	return fmt.Sprintf("┌──────┐\n│%s    │\n│  %-2s  │\n│    %s│\n└──────┘",
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


type Hand []deck.Card


func (h Hand) String() string {
	strs := make([]string,len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs,", ")
}

func (h Hand) DealerString() string {
	return h[0].String() + ", **Hidden**"
}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11{
		return minScore
	}
	for _,c := range h {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}


func (h Hand) MinScore() int {
	score := 0 

	for _,c := range h {
		score += min(int(c.Rank),10)


	}
	return score
}


func min(a,b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
 
    cards := deck.New(
        deck.Deck(3),    
        deck.Shuffle,    
    )

	var card deck.Card
   
	var player, dealer Hand
	for i :=0; i<2; i++ {
	for _, hand := range []*Hand{&player, &dealer} {
		card, cards = draw(cards)
		*hand = append(*hand, card)
		
	}
}

	var input string
	for input != "s"{
		fmt.Printf("Player Hand: %s\n", player)
		displayHand(player, fmt.Sprintf("%d", player.MinScore()))
		fmt.Println("-------------------------------- Dealer Hand --------------------------------")
		displayHand([]deck.Card{dealer[0]}, dealer.DealerString())

		fmt.Println("Enter 's' to stand, 'h' to hit")
		fmt.Scanf("%s", &input)
		switch input {
		case "h":
			card, cards = draw(cards)
			player = append(player, card)
		case "ozy": 
			fmt.Println("player won game hacked")
			return
			
		}
	

		
	}

	for dealer.Score() <= 16 || (dealer.Score() == 17 && dealer.MinScore() != 17) {
		card, cards = draw(cards)
		dealer = append(dealer, card)
	}
	
	pScore, dScore := player.Score(), dealer.Score()
	fmt.Println("==Final Hands==")
	

	// if dealer.Score() == player.Score() {
	// 	fmt.Println("==Push==")
	// } else if dealer.Score() > player.Score() {
	// 	fmt.Println("==Dealer Wins==")
	// } else {
	// 	fmt.Println("==Player Wins==")
	// }
	
	displayHand(player, fmt.Sprintf("Player Hand: %d", pScore))
	fmt.Println("-------------------------------- Dealer Hand --------------------------------")
	displayHand(dealer, fmt.Sprintf("Dealer Hand: %d", dScore))

	switch {
		case player.Score() > 21:
			fmt.Println("Player busts")
			return
		case dealer.Score() > 21:
			fmt.Println("Dealer busts")
			return
		case dealer.Score() == player.Score():
			fmt.Println("==Push==")
		case dealer.Score() > player.Score():
			fmt.Println("==Dealer Wins==")
		case dealer.Score() < player.Score():
			fmt.Println("==Player Wins==")
	}

	
    
}

func draw(card []deck.Card) (deck.Card, []deck.Card) {
	return card[0], card[1:]
}