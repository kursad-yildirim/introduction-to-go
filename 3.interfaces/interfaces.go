package main

import "fmt"

type triangle struct {
	height int
	base   int
}

type rectangle struct {
	side1 int
	side2 int
}

type square struct {
	side int
}

type shape interface {
	pArea()
}

func main() {
	t := triangle{height: 2, base: 6}
	r := rectangle{4, 4}
	s := square{5}
	sh := []shape{t, r, s}

	fmt.Println(t, r, s)

	sh[0].pArea()
	sh[1].pArea()
	sh[2].pArea()

	printArea(t)
	printArea(r)
	printArea(s)
}

/*
func triangleArea(t triangle) float32 {
	return 0.5 * float32(t.base) * float32(t.height)
}

func rectangleArea(r rectangle) int {
	return r.side1 * r.side2
}

func squareArea(s square) int {
	return s.side * s.side
}
*/

func (t triangle) areaIs() float32 {
	return 0.5 * float32(t.base) * float32(t.height)
}

func (r rectangle) areaIs() int {
	return r.side1 * r.side2
}

func (s square) areaIs() int {
	return s.side * s.side
}

func (t triangle) pArea() {
	fmt.Println("Triangle: ", t.areaIs())
}

func (r rectangle) pArea() {
	fmt.Println("Rectanngle: ", r.areaIs())
}

func (s square) pArea() {
	fmt.Println("Square: ", s.areaIs())
}

func printArea(sh shape) {
	sh.pArea()
}
