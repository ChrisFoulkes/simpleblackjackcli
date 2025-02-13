package entity

import (
	"fmt"
)

// Card struct with predefined values
type Card struct {
	Rank  string
	Suit  string
	Value int // Numerical value for Blackjack
}

func NewCard(rank string, suit string, value int) Card {
	return Card{rank, suit, value}
}

// String method for displaying cards
func (c Card) String() string {
	return fmt.Sprintf("%s of %s", c.Rank, c.Suit)
}
