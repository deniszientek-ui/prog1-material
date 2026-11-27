package Karten

import (
	"2026_10_27/karten/rank"
	"2026_10_27/karten/suit"
	"fmt"
)

func Example() {

	kreuz := suit.Kreuz
	dame := rank.Dame

	fmt.Println(kreuz)
	fmt.Println(dame)

	// Output:
	// ♣
	// D
}
