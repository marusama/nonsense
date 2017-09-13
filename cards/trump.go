package cards

import "fmt"

type Trump byte

const (
	Trump0TheFool          Trump = 0
	Trump1TheMagician      Trump = 1
	Trump2TheHighPriestess Trump = 2
	Trump3TheEmpress       Trump = 3
	Trump4TheEmperor       Trump = 4
	Trump5TheHierophant    Trump = 5
	Trump6TheLovers        Trump = 6
	Trump7TheChariot       Trump = 7
	Trump8Strength         Trump = 8
	Trump9TheHermit        Trump = 9
	Trump10WheelOfFortune  Trump = 10
	Trump11Justice         Trump = 11
	Trump12TheHangedMan    Trump = 12
	Trump13Death           Trump = 13
	Trump14Temperance      Trump = 14
	Trump15TheDevil        Trump = 15
	Trump16TheTower        Trump = 16
	Trump17TheStar         Trump = 17
	Trump18TheMoon         Trump = 18
	Trump19TheSun          Trump = 19
	Trump20Judgement       Trump = 20
	Trump21TheWorld        Trump = 21
)

func (t Trump) String() string {
	switch t {
	case Trump0TheFool:
		return "The Fool"
	case Trump1TheMagician:
		return "The Magician"
	case Trump2TheHighPriestess:
		return "The High Priestess"
	case Trump3TheEmpress:
		return "The Empress"
	case Trump4TheEmperor:
		return "The Emperor"
	case Trump5TheHierophant:
		return "The Hierophant"
	case Trump6TheLovers:
		return "The Lovers"
	case Trump7TheChariot:
		return "The Chariot"
	case Trump8Strength:
		return "Strength"
	case Trump9TheHermit:
		return "The Hermit"
	case Trump10WheelOfFortune:
		return "Wheel of Fortune"
	case Trump11Justice:
		return "Justice"
	case Trump12TheHangedMan:
		return "The Hanged Man"
	case Trump13Death:
		return "Death"
	case Trump14Temperance:
		return "Temperance"
	case Trump15TheDevil:
		return "The Devil"
	case Trump16TheTower:
		return "The Tower"
	case Trump17TheStar:
		return "The Star"
	case Trump18TheMoon:
		return "The Moon"
	case Trump19TheSun:
		return "The Sun"
	case Trump20Judgement:
		return "Judgement"
	case Trump21TheWorld:
		return "The World"
	}
	return fmt.Sprintf("invalid trump: %08b", byte(t))
}
