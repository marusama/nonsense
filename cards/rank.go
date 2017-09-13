package cards

import "fmt"

type Rank byte

const (
	RankAce    = 0x1 // 0001
	RankTwo    = 0x2 // 0010
	RankThree  = 0x3 // 0011
	RankFour   = 0x4 // 0100
	RankFive   = 0x5 // 0101
	RankSix    = 0x6 // 0110
	RankSeven  = 0x7 // 0111
	RankEight  = 0x8 // 1000
	RankNine   = 0x9 // 1001
	RankTen    = 0xA // 1010
	RankJack   = 0xB // 1011
	RankKnight = 0xC // 1100
	RankQueen  = 0xD // 1101
	RankKing   = 0xE // 1110
	RankJoker  = 0xF // 1111
)

func (r Rank) String() string {
	switch r {
	case RankAce:
		return "Ace"
	case RankTwo:
		return "Two"
	case RankThree:
		return "Three"
	case RankFour:
		return "Four"
	case RankFive:
		return "Five"
	case RankSix:
		return "Six"
	case RankSeven:
		return "Seven"
	case RankEight:
		return "Eight"
	case RankNine:
		return "Nine"
	case RankTen:
		return "Ten"
	case RankJack:
		return "Jack"
	case RankKnight:
		return "Knight"
	case RankQueen:
		return "Queen"
	case RankKing:
		return "King"
	case RankJoker:
		return "Joker"
	}
	return fmt.Sprintf("invalid rank: %08b", byte(r))
}
