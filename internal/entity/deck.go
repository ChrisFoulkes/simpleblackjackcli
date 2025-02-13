package entity

import (
	"math/rand"
)

type Deck []Card

var ranks = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
var suits = []string{"♠", "♥", "♦", "♣"}

func NewDeck() Deck {
	deck := Deck{}
	for _, rank := range ranks {
		for _, suit := range suits {
			value := getCardValue(rank) // Assign value once
			deck = append(deck, NewCard(rank, suit, value))
		}
	}
	return deck
}

// Shuffle randomizes the deck order
func (d *Deck) Shuffle() {
	rand.Shuffle(len(*d), func(i, j int) {
		(*d)[i], (*d)[j] = (*d)[j], (*d)[i]
	})
}

func (d *Deck) Draw() (Card, bool) {
	if len(*d) == 0 {
		*d = NewDeck()
	}
	card := (*d)[0]
	*d = (*d)[1:] // Remove drawn card
	return card, true
}

// getCardValue assigns a fixed numerical value to each rank
func getCardValue(rank string) int {
	switch rank {
	case "A":
		return 11 // Aces default to 11, will adjust in Hand logic
	case "K", "Q", "J", "10":
		return 10
	default:
		return int(rank[0] - '0') // Convert "2"-"9" to int
	}
}

func (d *Deck) DrawMany(count int) []Card {

	cards := make([]Card, 0, count)

	// Draw 'count' number of cards
	for i := 0; i < count; i++ {
		// If deck is empty, create new deck
		if len(*d) == 0 {
			*d = NewDeck()
		}

		// Draw a card and append to our slice
		card, ok := d.Draw()
		if ok {
			cards = append(cards, card)
		}
	}

	return cards
}
