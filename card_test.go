package deck

import "fmt"

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