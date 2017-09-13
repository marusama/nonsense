package cards

import "fmt"

type Arcana byte

const (
	ArcanaMinor Arcana = 0 // 0
	ArcanaMajor Arcana = 1 // 1
)

func (a Arcana) String() string {
	switch a {
	case ArcanaMinor:
		return "Minor"
	case ArcanaMajor:
		return "Major"
	}
	return fmt.Sprintf("invalid arcana: %08b", byte(a))
}
