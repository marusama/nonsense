package cards

import (
	"strconv"
)

func CardToShortString(c Card) string {
	if c.GetArcana() == ArcanaMinor {
		s := c.GetSuit()
		var suit string
		switch s {
		case SuitSpade:
			suit = "♠"
		case SuitHeart:
			suit = "♥"
		case SuitDiamond:
			suit = "♦"
		case SuitClub:
			suit = "♣"
		}
		r := c.GetRank()
		switch r {
		case RankAce:
			if c.IsBack() {
				return "*"
			}
			return "A" + suit
		case RankTwo, RankThree, RankFour, RankFive, RankSix, RankSeven, RankEight, RankNine:
			if c.IsBack() {
				return "*"
			}
			return strconv.Itoa(int(r)) + suit
		case RankTen:
			if c.IsBack() {
				return "*"
			}
			return "T" + suit
		case RankJack:
			if c.IsBack() {
				return "*"
			}
			return "J" + suit
		case RankKnight:
			if c.IsBack() {
				return "*"
			}
			return "N" + suit
		case RankQueen:
			if c.IsBack() {
				return "*"
			}
			return "Q" + suit
		case RankKing:
			if c.IsBack() {
				return "*"
			}
			return "K" + suit
		case RankJoker:
			switch JokerSuit(s) {
			case JokerSuitRed:
				if c.IsBack() {
					return "*"
				}
				return "RJ" // Red Joker
			case JokerSuitBlack:
				if c.IsBack() {
					return "*"
				}
				return "★" // Black Joker
			case JokerSuitWhite:
				if c.IsBack() {
					return "*"
				}
				return "☆" // White Joker
			}
		}
	} else { // Major Arcana
		t := c.GetTrump()
		if t >= Trump0TheFool && t <= Trump21TheWorld {
			if c.IsBack() {
				return "*"
			}
			return "T" + strconv.FormatInt(int64(t), 10)
		}
	}
	return c.String()
}
