package main

type cardNumber string

const (
	CARD_NUMBER_2     cardNumber = "2"
	CARD_NUMBER_3     cardNumber = "3"
	CARD_NUMBER_4     cardNumber = "4"
	CARD_NUMBER_5     cardNumber = "5"
	CARD_NUMBER_6     cardNumber = "6"
	CARD_NUMBER_7     cardNumber = "7"
	CARD_NUMBER_8     cardNumber = "8"
	CARD_NUMBER_9     cardNumber = "9"
	CARD_NUMBER_10    cardNumber = "10"
	CARD_NUMBER_JACK  cardNumber = "J"
	CARD_NUMBER_QUEEN cardNumber = "Q"
	CARD_NUMBER_KING  cardNumber = "K"
	CARD_NUMBER_ACE   cardNumber = "A"
)

var cardDeckNumbers = []cardNumber{
	CARD_NUMBER_2,
	CARD_NUMBER_3,
	CARD_NUMBER_4,
	CARD_NUMBER_5,
	CARD_NUMBER_6,
	CARD_NUMBER_7,
	CARD_NUMBER_8,
	CARD_NUMBER_9,
	CARD_NUMBER_10,
	CARD_NUMBER_JACK,
	CARD_NUMBER_QUEEN,
	CARD_NUMBER_KING,
	CARD_NUMBER_ACE,
}

func (c cardNumber) Name() string {
	switch c {
	case CARD_NUMBER_2:
		return "2"
	case CARD_NUMBER_3:
		return "3"
	case CARD_NUMBER_4:
		return "4"
	case CARD_NUMBER_5:
		return "5"
	case CARD_NUMBER_6:
		return "6"
	case CARD_NUMBER_7:
		return "7"
	case CARD_NUMBER_8:
		return "8"
	case CARD_NUMBER_9:
		return "9"
	case CARD_NUMBER_10:
		return "10"
	case CARD_NUMBER_JACK:
		return "jack"
	case CARD_NUMBER_QUEEN:
		return "queen"
	case CARD_NUMBER_KING:
		return "king"
	case CARD_NUMBER_ACE:
		return "ace"
	default:
		return "0"
	}
}
