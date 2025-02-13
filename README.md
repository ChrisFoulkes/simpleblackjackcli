# Go Blackjack Game

A terminal-based Blackjack game implemented in Go using the Bubble Tea framework for the terminal user interface.

## Description

This is a simple implementation of the classic Blackjack card game where players compete against a dealer. The game features:
- Interactive terminal UI with command history
- Card drawing and hand evaluation
- Dealer AI that follows standard Blackjack rules
- Unicode card suit symbols (♠, ♥, ♦, ♣)

## Prerequisites

- Go 1.24 or higher
- The following dependencies (automatically installed via go.mod):
  - github.com/charmbracelet/bubbletea
  - github.com/charmbracelet/lipgloss

## Installation

1. Clone the repository
2. Navigate to the project directory
3. Install dependencies:
   ```bash
   go mod tidy
   ```


## Running the Game

To start the game, run:
```bash
go run main.go
```


## How to Play

1. Type `draw` and press Enter to start a new hand
2. After receiving your initial cards, you can:
   - Select "Draw" to get another card
   - Select "Show" to end your turn and see the dealer's play
3. Type `quit` to exit the game

## Game Rules

- The goal is to get a hand value closer to 21 than the dealer without going over
- Number cards (2-10) are worth their face value
- Face cards (J, Q, K) are worth 10
- Aces are worth 11 but automatically adjust to 1 if the hand would bust
- The dealer must draw on 16 and stand on 17

## Project Structure

- `main.go` - Entry point and program initialization
- `internal/game/game.go` - Core game logic and state management
- `internal/ui/console.go` - Terminal UI implementation
- `internal/entity/` - Game entities (cards, deck, hand)
  - `card.go` - Card representation
  - `deck.go` - Deck management and card drawing
  - `hand.go` - Hand evaluation and card collection

