package main

import (
	"fmt"
)

type traingle struct {
	base   float64
	height float64
}

type square struct {
	side float64
}

type shape interface {
	getArea() float64
}

func main() {
	s := square{
		side: 2,
	}
	t := traingle{
		height: 2,
		base:   2,
	}
	printArea(s)
	printArea(t)
}

func (s square) getArea() float64 {
	return s.side * s.side
}

func (t traingle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func printArea(s shape) {
	fmt.Println("Area:", s.getArea())
}
