package game

import (
	"fmt"
	"game/internal/entity"
	"game/internal/ui"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

type Game struct {
	deck       entity.Deck
	hand       entity.Hand
	dealerHand entity.Hand
	console    *ui.Console
	running    bool
	input      string
}

// Add this new type for our timer messages
type dealerActionMsg struct{}

func New() *Game {
	return &Game{
		deck:       entity.NewDeck(),
		hand:       entity.NewHand(),
		dealerHand: entity.NewHand(),
		console:    ui.NewConsole(),
		running:    true,
		input:      "",
	}
}

func (g *Game) Init() tea.Cmd {
	return nil
}

func (g *Game) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case dealerActionMsg:
		if g.dealerHand.GetTotalValue() < 17 {
			g.console.Display("\nDealer must draw (below 17)...")
			card, _ := g.deck.Draw()
			g.dealerHand.AddCard(card)
			g.console.Display(fmt.Sprintf("Dealer drew: %s", card.String()))
			g.console.Display(fmt.Sprintf("Dealer's new total: %d", g.dealerHand.GetTotalValue()))
			return g, g.startDealerPlay() // Continue with next action
		} else {
			g.displayFinalState()
			g.determineWinner()
			return g, nil
		}

	case tea.KeyMsg:
		// Handle different states
		if g.console.GetGameState() == ui.StateDrawOption {
			switch msg.Type {
			case tea.KeyUp:
				g.console.NavigateOptions(-1)
				return g, nil
			case tea.KeyDown:
				g.console.NavigateOptions(1)
				return g, nil
			case tea.KeyEnter:
				// Handle option selection
				selected := g.console.GetSelected()
				g.console.SetGameState(ui.StateNormal)
				if selected == 0 {
					g.drawAdditionalCard(&g.hand)
					g.displayGameState()
					g.checkHandState()
					return g, nil
				} else {
					// Change this line to use showDealerHand
					return g, g.showDealerHand()
				}
			}
			return g, nil
		}

		switch msg.Type {
		case tea.KeyEnter:
			command := g.input
			g.console.AddToHistory(command)
			g.input = ""
			g.console.ClearInput()

			switch command {
			case "quit":
				g.running = false
				return g, tea.Quit
			case "draw":
				g.drawFirstHand(&g.hand)
				g.checkHandState()
				return g, nil
			default:
				g.console.Display("Unknown command!")
			}
			return g, nil

		case tea.KeyUp:
			g.console.NavigateHistory(-1)
			g.input = g.console.GetInput()
			return g, nil

		case tea.KeyDown:
			g.console.NavigateHistory(1) // Go forward in history
			g.input = g.console.GetInput()
			return g, nil

		case tea.KeyCtrlC:
			return g, tea.Quit

		case tea.KeyBackspace:
			if len(g.input) > 0 {
				g.input = g.input[:len(g.input)-1]
				g.console.UpdateInput(g.input)
			}

		default:
			g.input += msg.String()
			g.console.UpdateInput(g.input)
		}
	}
	return g, nil
}

func (g *Game) View() string {
	if !g.running {
		return "Goodbye!\n"
	}
	return g.console.View("Welcome to the Game!")
}

func (g *Game) drawFirstHand(hand *entity.Hand) {
	g.resetGame()
	cards := g.deck.DrawMany(2)
	for _, card := range cards {
		hand.AddCard(card)
	}
	g.displayGameState()
}

func (g *Game) resetGame() {
	g.console.ClearMessages()
	g.hand = entity.NewHand()
	g.dealerHand = entity.NewHand()
	g.deck.Shuffle()
}

func (g *Game) displayGameState() {
	g.console.ClearMessages()
	g.console.Display("\nCurrent Game State:")
	g.console.Display("==================")

	g.displayHand("Your Hand", &g.hand)

	if len(g.dealerHand.GetCards()) > 0 {
		g.displayHand("Dealer's Hand", &g.dealerHand)
	}

	g.console.Display("\n==================")
}

func (g *Game) displayHand(label string, hand *entity.Hand) {
	g.console.Display(fmt.Sprintf("\n%s:", label))
	g.console.Display("----------------")
	for _, card := range hand.GetCards() {
		g.console.Display(fmt.Sprintf("  %s", card.String()))
	}
	g.console.Display(fmt.Sprintf("Total value: %d", hand.GetTotalValue()))
}

func (g *Game) drawAdditionalCard(hand *entity.Hand) {
	card, _ := g.deck.Draw()
	hand.AddCard(card)
	g.displayGameState()
}

func (g *Game) checkHandState() {
	if g.hand.GetTotalValue() > 21 {
		g.console.Display("\nFinal Result:")
		g.console.Display("==================")
		g.console.Display("Bust! You lose!")
		g.console.SetGameState(ui.StateNormal)
	} else {
		g.console.SetGameState(ui.StateDrawOption)
	}
}

// Modify showDealerHand to handle the delayed actions
func (g *Game) showDealerHand() tea.Cmd {
	g.dealerInitialDraw()
	g.displayGameState()
	return g.startDealerPlay()
}

// New function to start the dealer play sequence
func (g *Game) startDealerPlay() tea.Cmd {
	return tea.Tick(1200*time.Millisecond, func(time.Time) tea.Msg {
		return dealerActionMsg{}
	})
}

func (g *Game) dealerInitialDraw() {
	if len(g.dealerHand.GetCards()) == 0 {
		cards := g.deck.DrawMany(2)
		for _, card := range cards {
			g.dealerHand.AddCard(card)
		}
	}
}

func (g *Game) displayFinalState() {
	g.console.Display("\nFinal Game State:")
	g.console.Display("==================")
	g.displayHand("Your Hand", &g.hand)
	g.displayHand("Dealer's Final Hand", &g.dealerHand)
	g.console.Display("\nFinal Result:")
	g.console.Display("==================")
}

func (g *Game) determineWinner() {
	playerTotal := g.hand.GetTotalValue()
	dealerTotal := g.dealerHand.GetTotalValue()

	switch {
	case dealerTotal > 21:
		g.console.Display(fmt.Sprintf("Dealer busts with %d! You win!", dealerTotal))
	case dealerTotal == 21:
		g.console.Display("Dealer has Blackjack! Dealer wins!")
	case dealerTotal > playerTotal && dealerTotal <= 21:
		g.console.Display(fmt.Sprintf("Dealer wins with %d vs your %d!", dealerTotal, playerTotal))
	case dealerTotal < playerTotal:
		g.console.Display(fmt.Sprintf("You win with %d vs dealer's %d!", playerTotal, dealerTotal))
	case dealerTotal == playerTotal:
		g.console.Display(fmt.Sprintf("It's a tie at %d!", playerTotal))
	}
}
