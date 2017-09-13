package cards

import "fmt"

type Suit byte

const (
	SuitSpade   Suit = 0 // 00
	SuitHeart   Suit = 1 // 01
	SuitDiamond Suit = 2 // 10
	SuitClub    Suit = 3 // 11
)

func (s Suit) String() string {
	switch s {
	case SuitSpade:
		return "Spade"
	case SuitHeart:
		return "Heart"
	case SuitDiamond:
		return "Diamond"
	case SuitClub:
		return "Club"
	}
	return fmt.Sprintf("invalid suit: %08b", byte(s))
}

type JokerSuit byte

const (
	JokerSuitRed   JokerSuit = 1 // 01
	JokerSuitBlack JokerSuit = 2 // 10
	JokerSuitWhite JokerSuit = 3 // 11
)

func (s JokerSuit) String() string {
	switch s {
	case JokerSuitRed:
		return "Red"
	case JokerSuitBlack:
		return "Black"
	case JokerSuitWhite:
		return "White"
	}
	return fmt.Sprintf("invalid joker suit: %08b", byte(s))
}
