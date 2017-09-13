package cards

import "fmt"

// Card playing card
// 0 0 00 0000
// ^ ^ ^  ^
// | | |  rank bits
// | | suit bits
// | arcana bit (0 - minor)
// back bit
// 0 1 000000
// ^ ^ ^
// | | |
// | | trump bits
// | arcana bit (1 - major)
// back bit
type Card byte

const (
	RankMask  = 0x0F // 00 00 1111
	SuitMask  = 0x30 // 00 11 0000
	TrumpMask = 0x3F // 00 11 1111
	ArcanaBit = 0x40 // 01 00 0000
	BackBit   = 0x80 // 10 00 0000

	ArcanaBitOffset = 6
	SuitMaskOffset  = 4
)

func (c Card) GetRank() Rank {
	return Rank(c & RankMask)
}

func (c Card) GetSuit() Suit {
	return Suit((c & SuitMask) >> SuitMaskOffset)
}

func (c Card) GetTrump() Trump {
	return Trump(c & TrumpMask)
}

func (c Card) GetArcana() Arcana {
	return Arcana((c & ArcanaBit) >> ArcanaBitOffset)
}

func (c Card) IsBack() bool {
	return (c & BackBit) == BackBit
}

func (c Card) SetBack(back bool) Card {
	if back {
		return Card(c | BackBit)
	}
	return Card(byte(c) & ^byte(BackBit))
}

func (c Card) String() string {
	const back string = "Back"
	if c.GetArcana() == ArcanaMinor {
		r := c.GetRank()
		s := c.GetSuit()
		switch r {
		case RankAce, RankTwo, RankThree, RankFour, RankFive, RankSix, RankSeven, RankEight, RankNine, RankTen,
			RankJack, RankKnight, RankQueen, RankKing:
			if c.IsBack() {
				return back
			}
			return fmt.Sprintf("%s of %ss", r, s)
		case RankJoker:
			switch JokerSuit(s) {
			case JokerSuitRed:
				fallthrough
			case JokerSuitBlack:
				fallthrough
			case JokerSuitWhite:
				if c.IsBack() {
					return back
				}
				return fmt.Sprintf("%s %s", JokerSuit(s), r)
			}
		}
	} else { // Major Arcana
		t := c.GetTrump()
		if t >= Trump0TheFool && t <= Trump21TheWorld {
			if c.IsBack() {
				return back
			}
			return t.String()
		}
	}
	return fmt.Sprintf("invalid card: %08b", byte(c))
}
