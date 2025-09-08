package main

import (
	"fmt"
	"strings"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	deck "github.com/ozyftw/gojack"
)


var (
	
	redSuit = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")).Bold(true)
	blackSuit = lipgloss.NewStyle().Foreground(lipgloss.Color("#000000")).Bold(true)
	

	cardStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#FFFFFF")).
		Background(lipgloss.Color("#F8F8F8")).
		Padding(0, 1).
		Margin(0, 1)

	hiddenCardStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#666666")).
		Background(lipgloss.Color("#333333")).
		Foreground(lipgloss.Color("#666666")).
		Padding(0, 1).
		Margin(0, 1)


	titleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(0, 1).
		Bold(true)

	subtitleStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#7D56F4")).
		Bold(true)

	scoreStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#04B575")).
		Bold(true)

	instructionStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#626262")).
		Italic(true)

	winStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#04B575")).
		Bold(true)

	loseStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FF0000")).
		Bold(true)

	pushStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFA500")).
		Bold(true)

	buttonStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FFF")).
		Background(lipgloss.Color("#7D56F4")).
		Padding(0, 1).
		Margin(0, 1).
		Border(lipgloss.RoundedBorder())

	selectedButtonStyle = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#000")).
		Background(lipgloss.Color("#FFFF00")).
		Padding(0, 1).
		Margin(0, 1).
		Border(lipgloss.RoundedBorder()).
		Bold(true)
)

type gameState int

const (
	gameStart gameState = iota
	playerTurn
	dealerTurn
	gameOver
)

type Hand []deck.Card

func (h Hand) String() string {
	strs := make([]string, len(h))
	for i := range h {
		strs[i] = h[i].String()
	}
	return strings.Join(strs, ", ")
}

func (h Hand) DealerString() string {
	if len(h) == 0 {
		return ""
	}
	return h[0].String() + ", **Hidden**"
}

func (h Hand) Score() int {
	minScore := h.MinScore()
	if minScore > 11 {
		return minScore
	}
	for _, c := range h {
		if c.Rank == deck.Ace {
			return minScore + 10
		}
	}
	return minScore
}

func (h Hand) MinScore() int {
	score := 0
	for _, c := range h {
		score += min(int(c.Rank), 10)
	}
	return score
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type model struct {
	cards       []deck.Card
	player      Hand
	dealer      Hand
	state       gameState
	cursor      int
	gameResult  string
	showDealer  bool
	animating   bool
}

func initialModel() model {
	cards := deck.New(
		deck.Deck(3),
		deck.Shuffle,
	)

	m := model{
		cards:      cards,
		player:     Hand{},
		dealer:     Hand{},
		state:      gameStart,
		cursor:     0,
		showDealer: false,
	}

	// Deal initial cards
	for i := 0; i < 2; i++ {
		for _, hand := range []*Hand{&m.player, &m.dealer} {
			card, remaining := draw(m.cards)
			m.cards = remaining
			*hand = append(*hand, card)
		}
	}

	m.state = playerTurn
	return m
}

type tickMsg time.Time

func tickCmd() tea.Cmd {
	return tea.Tick(time.Millisecond*500, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.state {
		case playerTurn:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "up", "k":
				if m.cursor > 0 {
					m.cursor--
				}
			case "down", "j":
				if m.cursor < 1 {
					m.cursor++
				}
			case "enter", " ":
				if m.cursor == 0 { // Hit
					card, remaining := draw(m.cards)
					m.cards = remaining
					m.player = append(m.player, card)
					
					if m.player.Score() > 21 {
						m.state = gameOver
						m.gameResult = "Player Busts! Dealer Wins!"
						m.showDealer = true
					}
				} else { 
					m.state = dealerTurn
					m.showDealer = true
					return m, tickCmd()
				}
			case "ozy":
				m.state = gameOver
				m.gameResult = "Player Won! (Cheat Code Activated)"
				m.showDealer = true
			}
		case dealerTurn:
			
		case gameOver:
			switch msg.String() {
			case "ctrl+c", "q":
				return m, tea.Quit
			case "r", "enter":
				return initialModel(), nil
			}
		}

	case tickMsg:
		if m.state == dealerTurn {
			if m.dealer.Score() <= 16 || (m.dealer.Score() == 17 && m.dealer.MinScore() != 17) {
				card, remaining := draw(m.cards)
				m.cards = remaining
				m.dealer = append(m.dealer, card)
				return m, tickCmd()
			} else {
				// Determine winner
				m.state = gameOver
				pScore, dScore := m.player.Score(), m.dealer.Score()
				
				switch {
				case dScore > 21:
					m.gameResult = "Dealer Busts! Player Wins!"
				case dScore == pScore:
					m.gameResult = "Push! It's a Tie!"
				case dScore > pScore:
					m.gameResult = "Dealer Wins!"
				default:
					m.gameResult = "Player Wins!"
				}
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	var s strings.Builder

	// Title
	s.WriteString(titleStyle.Render("🃏 BLACKJACK 🃏"))
	s.WriteString("\n\n")

	// Player hand
	s.WriteString(subtitleStyle.Render("Your Hand"))
	s.WriteString(" ")
	s.WriteString(scoreStyle.Render(fmt.Sprintf("(Score: %d)", m.player.Score())))
	s.WriteString("\n")
	s.WriteString(renderHand(m.player, false))
	s.WriteString("\n")

	// Dealer hand
	s.WriteString(subtitleStyle.Render("Dealer's Hand"))
	if m.showDealer {
		s.WriteString(" ")
		s.WriteString(scoreStyle.Render(fmt.Sprintf("(Score: %d)", m.dealer.Score())))
	}
	s.WriteString("\n")
	s.WriteString(renderHand(m.dealer, !m.showDealer))
	s.WriteString("\n")

	// Game state specific content
	switch m.state {
	case playerTurn:
		if m.player.Score() > 21 {
			s.WriteString(loseStyle.Render("BUST! You went over 21!"))
			s.WriteString("\n\n")
		}
		
		// Action buttons
		s.WriteString("Choose your action:\n\n")
		
		hitButton := "Hit"
		standButton := "Stand"
		
		if m.cursor == 0 {
			hitButton = selectedButtonStyle.Render(hitButton)
			standButton = buttonStyle.Render(standButton)
		} else {
			hitButton = buttonStyle.Render(hitButton)
			standButton = selectedButtonStyle.Render(standButton)
		}
		
		s.WriteString(lipgloss.JoinHorizontal(lipgloss.Left, hitButton, standButton))
		s.WriteString("\n\n")
		s.WriteString(instructionStyle.Render("Use ↑/↓ or j/k to navigate, Enter/Space to select"))

	case dealerTurn:
		s.WriteString(instructionStyle.Render("Dealer is playing..."))

	case gameOver:
		s.WriteString("\n")
		
	
		if strings.Contains(m.gameResult, "Player Wins") || strings.Contains(m.gameResult, "Dealer Busts") {
			s.WriteString(winStyle.Render("🎉 " + m.gameResult + " 🎉"))
		} else if strings.Contains(m.gameResult, "Push") || strings.Contains(m.gameResult, "Tie") {
			s.WriteString(pushStyle.Render("🤝 " + m.gameResult + " 🤝"))
		} else {
			s.WriteString(loseStyle.Render("💔 " + m.gameResult + " 💔"))
		}
		
		s.WriteString("\n\n")
		s.WriteString(buttonStyle.Render("Press 'r' or Enter to play again"))
	}

	s.WriteString("\n\n")
	s.WriteString(instructionStyle.Render("Press 'q' to quit"))

	return s.String()
}

func renderHand(hand Hand, hideSecond bool) string {
	if len(hand) == 0 {
		return ""
	}

	var cards []string
	for i, card := range hand {
		if hideSecond && i == 1 {
			cards = append(cards, renderHiddenCard())
		} else {
			cards = append(cards, renderCard(card))
		}
	}

	return lipgloss.JoinHorizontal(lipgloss.Top, cards...)
}

func renderCard(card deck.Card) string {
	var suit string
	var suitStyle lipgloss.Style

	switch card.Suit.String() {
	case "Spade":
		suit = "♠"
		suitStyle = blackSuit
	case "Heart":
		suit = "♥"
		suitStyle = redSuit
	case "Diamond":
		suit = "♦"
		suitStyle = redSuit
	case "Club":
		suit = "♣"
		suitStyle = blackSuit
	default:
		suit = "?"
		suitStyle = blackSuit
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

	cardContent := fmt.Sprintf("%s\n%s\n%s", 
		fmt.Sprintf("%-2s", rank),
		suitStyle.Render(suit),
		fmt.Sprintf("%2s", rank))

	return cardStyle.Render(cardContent)
}

func renderHiddenCard() string {
	cardContent := "??\n?\n??"
	return hiddenCardStyle.Render(cardContent)
}

func draw(cards []deck.Card) (deck.Card, []deck.Card) {
	return cards[0], cards[1:]
}

func main() {
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
	}
}