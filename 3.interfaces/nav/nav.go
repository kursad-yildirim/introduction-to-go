package main

import (
	"fmt"
	"math"
)

type point struct {
	x, y int
}

type currentLcoation struct {
	point
	moment string
}

type startPosition struct {
	point
	routeId string
}

type destinationPosition struct {
	point
	routeId string
}

type landmarkLocation struct {
	point
	name           string
	visitorRanking int
}

type bookmarkedLoacation struct {
	point
	name   string
	folder string
}

type mover interface {
	move(x, y int)
	getPoint() point
}

func main() {
	p := point{1, 3}
	c := currentLcoation{
		point:  point{1, 1},
		moment: "now",
	}
	s := startPosition{
		point: point{4, 5},
	}
	d := destinationPosition{
		point: point{76, 45},
	}
	fmt.Printf("point is %#v\n", p)
	fmt.Printf("I am at %v, %v\n ", c.point, c.moment)
	p.move(8, 9)
	c.move(8, 9)
	fmt.Printf("point is %#v\n", p)
	fmt.Printf("I am at %v, %v\n ", c.point, c.moment)

	fmt.Println(calcRoute(&s, &d))
	fmt.Println(calcRoute(&c, &d))
}

func (s startPosition) getPoint() point {
	return s.point
}

func (d destinationPosition) getPoint() point {
	return d.point
}

func (c currentLcoation) getPoint() point {
	return c.point
}

func (l landmarkLocation) getPoint() point {
	return l.point
}

func (b bookmarkedLoacation) getPoint() point {
	return b.point
}

func (p point) getPoint() point {
	return p
}

func (p *point) move(x, y int) {
	p.x = x
	p.y = y
}

func calcRoute(s mover, d mover) float64 {
	var r point
	speed := 10

	r.x = d.getPoint().x - s.getPoint().x
	r.y = d.getPoint().y - s.getPoint().y

	return math.Sqrt(float64(r.x*r.x+r.y*r.y)) / float64(speed)
}
