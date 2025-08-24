//go:generate stringer -type=Suit,Rank

package deck

import "fmt"

type Suit uint8

const (
	Spade Suit = iota
	Diamond 
	Club 
	Heart 
	Joker
)


type Rank uint8

const (
	_ Rank = iota
	Ace 
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Suit 
	Rank
	//suit
	//rank

}

func main() {
	card := Card{Suit: Spade, Rank: Ace}
	fmt.Println(card)
}


func (c Card) String() string {
	if c.Suit == Joker {
		return c.Suit.String()
	}

	

	return fmt.Sprintf("%s of %ss", c.Rank.String(), c.Suit.String())
}