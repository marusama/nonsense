package cards

import "math/rand"

type Deck []Card

// Shuffle deck
// https://en.wikipedia.org/wiki/Fisher-Yates_shuffle
func Shuffle(d Deck) {
	for i := len(d) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d[i], d[j] = d[j], d[i]
	}
}
