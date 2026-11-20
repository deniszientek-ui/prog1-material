package time

import "fmt"

func ExampleTime() {
	// Für 3
	l := FromHours(3)

	s := l.Seconds()
	m := l.Minutes()
	h := l.Hours()

	fmt.Println(s)
	fmt.Println(m)
	fmt.Println(h)

	// Output:
	// 10800
	// 180
	// 3
}
