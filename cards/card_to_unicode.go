package cards

func CardToUnicode(c Card) string {
	unicodeBack := string([]byte{0xF0, 0x9F, 0x82, 0xA0})
	if c.GetArcana() == ArcanaMinor {
		r := c.GetRank()
		s := c.GetSuit()
		switch r {
		case RankAce, RankTwo, RankThree, RankFour, RankFive, RankSix, RankSeven, RankEight, RankNine, RankTen,
			RankJack, RankKnight, RankQueen, RankKing:
			if c.IsBack() {
				return unicodeBack
			}
			var base []byte
			switch s {
			case SuitSpade: // F0 9F 82 A0
				base = []byte{0xF0, 0x9F, 0x82, 0xA0}
			case SuitHeart: // F0 9F 82 B0
				base = []byte{0xF0, 0x9F, 0x82, 0xB0}
			case SuitDiamond: // F0 9F 83 80
				base = []byte{0xF0, 0x9F, 0x83, 0x80}
			case SuitClub: // F0 9F 83 90
				base = []byte{0xF0, 0x9F, 0x83, 0x90}
			}
			base[3] += byte(r)
			return string(base)
		case RankJoker:
			switch JokerSuit(s) {
			case JokerSuitRed:
				if c.IsBack() {
					return unicodeBack
				}
				return string([]byte{0xF0, 0x9F, 0x82, 0xBF})
			case JokerSuitBlack:
				if c.IsBack() {
					return unicodeBack
				}
				return string([]byte{0xF0, 0x9F, 0x83, 0x8F})
			case JokerSuitWhite:
				if c.IsBack() {
					return unicodeBack
				}
				return string([]byte{0xF0, 0x9F, 0x83, 0x9F})
			}
		}
	} else { // Major Arcana
		t := c.GetTrump()
		if t >= Trump0TheFool && t <= Trump21TheWorld {
			if c.IsBack() {
				return unicodeBack
			}
			base := []byte{0xF0, 0x9F, 0x83, 0xA0}
			base[3] += byte(t)
			return string(base)
		}
	}
	return c.String()
}
