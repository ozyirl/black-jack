package deck

import (
	"fmt"
	"testing"
)

func ExampleCard(){
	fmt.Println(Card{Suit: Spade, Rank: Ace})
	fmt.Println(Card{Suit: Heart, Rank: Ace})
	fmt.Println(Card{Suit: Diamond, Rank: Ace})
	fmt.Println(Card{Suit: Club, Rank: Ace})
	fmt.Println(Card{Suit: Joker})
	// Output:
	// Ace of Spades
	// Ace of Hearts
	// Ace of Diamonds
	// Ace of Clubs
	// Joker
}



func TestNew (t *testing.T){
	cards := New()
	//13 ranks * 4 suits
	if len(cards) != 13*4 {
		t.Errorf("Expected 52 cards, got %d", len(cards))
	}

	
}