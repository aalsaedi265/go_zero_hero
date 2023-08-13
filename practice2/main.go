
package main

import "fmt"

type shape interface {
	getArea() float64
	printArea()
}

type square struct {
	sideLength float64
}

type triangle struct {
	height float64
	base   float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (s square) printArea() {
	fmt.Printf("Area of square: %f\n", s.getArea())
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (t triangle) printArea() {
	fmt.Printf("Area of triangle: %f\n", t.getArea())
}

func main() {
	s := square{sideLength: 4}
	t := triangle{height: 4, base: 3}

	shapes := []shape{s, t}
	for _, sh := range shapes {
		sh.printArea()
	}
}