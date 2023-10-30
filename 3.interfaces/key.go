package main

import "fmt"

type field byte

func main() {
	var x field = 1
	var y field = 2
	var z field = 3
	fmt.Println("x:", x)
	fmt.Println("y:", y)
	fmt.Println("z:", z)

	fmt.Printf("x: %d\n", x)
	fmt.Printf("y: %d\n", y)
	fmt.Printf("z: %d\n", z)
}

func (f field) String() string {
	switch f {
	case 1:
		return "one"
	case 2:
		return "two"
	case 3:
		return "three"
	}

	return "not my number"
}
