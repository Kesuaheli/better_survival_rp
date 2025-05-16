package main

type cardColor string

const (
	CARD_COLOR_DIAMONDS cardColor = "♦"
	CARD_COLOR_HEARTS   cardColor = "♥"
	CARD_COLOR_SPADES   cardColor = "♠"
	CARD_COLOR_CLUBS    cardColor = "♣"
)

var cardDeckColors = []cardColor{
	CARD_COLOR_DIAMONDS,
	CARD_COLOR_HEARTS,
	CARD_COLOR_SPADES,
	CARD_COLOR_CLUBS,
}

func (c cardColor) Name() string {
	switch c {
	case CARD_COLOR_DIAMONDS:
		return "diamonds"
	case CARD_COLOR_HEARTS:
		return "hearts"
	case CARD_COLOR_SPADES:
		return "spades"
	case CARD_COLOR_CLUBS:
		return "clubs"
	default:
		return "?"
	}
}
