package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type rectangle struct {
	side1 float64
	side2 float64
}

type square struct {
	side float64
}

type shape interface {
	calcArea() float64
}

func main() {
	t := triangle{4, 7}
	s := square{5}
	r := rectangle{6, 9}

	printArea(t)
	printArea(s)
	printArea(r)
}

func printArea(sh shape) {
	fmt.Println("Area is", sh.calcArea(), "cm^2")
}

func (t triangle) calcArea() float64 {
	return 0.5 * t.base * t.height
}
func (s square) calcArea() float64 {
	return s.side * s.side
}
func (r rectangle) calcArea() float64 {
	return r.side1 * r.side2
}
