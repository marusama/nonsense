package main

import (
	"fmt"

	"github.com/marusama/nonsense/cards"
)

func main() {
	// chan_chan.Run2()
	// chan_func.Run()

	//d := cards.NewDeck54()
	////cards.Shuffle(d)
	//
	//for _, c := range d {
	//	fmt.Println(c, cards.CardToUnicode(c))
	//}

	//d := cards.NewDeckTarot()

	//for i := range d {
	//	d[i] = d[i].SetBack(true)
	//}

	//printDeck(d, cards.CardToUnicode)
	//printDeck(d, cards.CardToUnicode)
	//for i := 0; i < 256; i++ {
	//	c := cards.Card(i)
	//	fmt.Println(fmt.Sprintf("%03d %08b %s %s %s", i, i, c, cards.CardToUnicode(c), cards.CardToShortString(c)))
	//}

	d := cards.NewDeckFull()
	//cards.Shuffle(d)
	for _, c := range d {
		fmt.Println(fmt.Sprintf("%08b %s %s %s", byte(c), c, cards.CardToUnicode(c), cards.CardToShortString(c)))
	}
}

func printDeck(d cards.Deck, f func(cards.Card) string) {
	for i, c := range d {
		fmt.Print(f(c))
		fmt.Print(" ")
		if (i+1)%14 == 0 {
			fmt.Println()
		}
	}
	fmt.Println()
}
