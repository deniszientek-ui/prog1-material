package coord

import "fmt"

func ExampleCoord() {
	c1 := Coord{8.5, 49.5}
	c2 := Coord{0, 0}

	d1 := c1.Distance(c2)

	fmt.Println(c1)
	fmt.Println(c2)
	fmt.Println(d1)

	// Nur Längengrad c1
	fmt.Println(c1.Longtitude)

	// Den Breitengrad von c2 verändern:
	c1.Latitude = 42
	fmt.Println(c2)
	// Output:

}
