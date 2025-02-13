package entity

type Hand struct {
	cards []Card
}

func NewHand() Hand {
	return Hand{
		cards: []Card{},
	}
}

func (h *Hand) AddCard(card Card) {
	h.cards = append(h.cards, card)
}

func (h *Hand) GetCards() []Card {
	return h.cards
}

func (h *Hand) GetTotalValue() int {
	total := 0
	aceCount := 0

	for _, card := range h.cards {
		value := getCardValue(card.Rank)
		total += value
		if card.Rank == "A" {
			aceCount++
		}
	}

	for i := 0; i < aceCount; i++ {
		if total > 21 {
			total -= 10
		}
	}

	return total
}
