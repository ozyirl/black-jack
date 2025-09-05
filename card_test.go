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

func TestDefaultSort(t *testing.T){
	cards := New(DefaultSort)
	expected := Card{Suit: Spade, Rank: Ace}
	if cards[0] != expected {
		t.Errorf("Expected Ace of Spades, got %v", cards[0])
	}
}


func TestJokers(t *testing.T){
	cards := New(Jokers(3))
	count := 0
	for _,c:= range cards{
		if c.Suit == Joker {
			count ++
		}
	}


	if count != 3 {
		t.Error("Expected 3 jokers, got")
	}
}


func TestFilter(t *testing.T){
	filter :=func(card Card) bool {
		return card.Rank == Two || card.Rank == Three
	}

	cards := New(Filter(filter))
	for _,c := range cards {
		if c.Rank == Two || c.Rank == Three {
			t.Error("expected twos and threes")
		}
	}
}


func TestDeck(t *testing.T){
	cards := New(Deck(3))
	//13 ranks * 4 suits * 3 decks
	if len(cards) != 13 * 4 * 3{
		t.Errorf("Expected %d cards, recieved %d cards", 13*4*3, len(cards))
	}
}