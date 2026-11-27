package rank

import "fmt"

type Rank int

const (
	Zwei Rank = 2 + iota
	Drei
	Vier
	Fuenf
	Sechs
	Sieben
	Acht
	Neun
	Zehn
	Bube
	Dame
	König
	Ass
)

func (r Rank) String() string {
	if r >= Zwei && r <= Zehn {
		return fmt.Sprintf("%d", r)
	}

	switch r {

	case Bube:
		return "B"
	case Dame:
		return "D"
	case König:
		return "K"
	case Ass:
		return "A"
	default:
		return ""
	}
}
