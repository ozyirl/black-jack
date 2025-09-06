# GoJack ♠️ ♦️ ♣️ ♥️ - Playing Card Deck Package

A Go package for cards. Not just any cards — your deck, your rules.
Blackjack, Poker, Rummy, test harnesses, whatever. GoJack is the engine that deals, shuffles, and stacks like a casino-grade shoe.

## Installation

```bash
go get github.com/ozyirl/gojack
```

Pull it down. Build your table. Let’s go.

## Usage

### Basic Usage

```go
package main

import (
    "fmt"
    "github.com/ozyirl/gojack/deck"
)

func main() {
    // Create a standard 52-card deck
    cards := deck.New()
    fmt.Printf("Created deck with %d cards\n", len(cards))

    // Create a card
    card := deck.Card{Suit: deck.Spade, Rank: deck.Ace}
    fmt.Println(card) // Output: Ace of Spades
}
```

Minimal. Precise. Weaponized simplicity.

### Deck Options

GoJack runs on options — stack them, combine them, bend the game to your will.

```go
// Create a shuffled deck
cards := deck.New(deck.Shuffle)

// Create a sorted deck
cards := deck.New(deck.DefaultSort)

// Create multiple decks
cards := deck.New(deck.Deck(3)) // 3 standard decks combined

// Add jokers
cards := deck.New(deck.Jokers(2)) // Add 2 jokers

// Filter out specific cards
filterLowCards := func(card deck.Card) bool {
    return card.Rank == deck.Two || card.Rank == deck.Three
}
cards := deck.New(deck.Filter(filterLowCards)) // Remove 2s and 3s

// Combine multiple options
cards := deck.New(
    deck.Deck(2),        // 2 decks
    deck.Jokers(4),      // 4 jokers
    deck.Shuffle,        // shuffled
)
```

Every option is a lever. Pull it.

### Custom Sorting

You’re not stuck with defaults. You want a table that runs your order? Write it.

```go
// Use custom sort function
customSort := deck.Sort(func(cards []deck.Card) func(i, j int) bool {
    return func(i, j int) bool {
        // Custom sorting logic
        return cards[i].Rank < cards[j].Rank
    }
})
cards := deck.New(customSort)
```

Not just sorting. Reprogramming the deck.

## Types

### Card

```go
type Card struct {
    Suit Suit
    Rank Rank
}
```

### Suits

- `deck.Spade`
- `deck.Diamond`
- `deck.Club`
- `deck.Heart`
- `deck.Joker`

### Ranks

- `deck.Ace` through `deck.King`
- Standard ranks: Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King

## Functions

- `New(...func([]Card) []Card) []Card` - Create a new deck with optional modifications
- `DefaultSort([]Card) []Card` - Sort cards in default order
- `Shuffle([]Card) []Card` - Shuffle the deck randomly
- `Jokers(n int) func([]Card) []Card` - Add n jokers to the deck
- `Filter(func(Card) bool) func([]Card) []Card` - Remove cards matching the filter
- `Deck(n int) func([]Card) []Card` - Create n copies of the deck

## Example for Blackjack

```go
package main

import (
    "fmt"
    "github.com/ozyirl/gojack/deck"
)

func main() {
    // Create a blackjack deck (usually 6-8 decks, shuffled)
    cards := deck.New(
        deck.Deck(6),     // 6 decks
        deck.Shuffle,     // shuffled
    )

    // Deal cards
    playerHand := []deck.Card{cards[0], cards[2]}
    dealerHand := []deck.Card{cards[1], cards[3]}

    fmt.Printf("Player: %v, %v\n", playerHand[0], playerHand[1])
    fmt.Printf("Dealer: %v, [Hidden]\n", dealerHand[0])
}
```

You’re not just printing cards. You’re simulating the pit.

## Testing

Run the tests with:

```bash
go test
```

Run with verbose output:

```bash
go test -v
```

Why GoJack?

Because the world didn’t need another boring deck implementation.
It needed this — lean, sharp, production-ready.
GoJack is not about cards. It’s about control.
You’re not coding games. You’re running casinos.
