package basics

import (
	"math"
)

type Rectangle struct {
	length float64
	width  float64
}

type Circle struct {
	radius float64
}

type Shape interface {
	Perimeter() float64
	Area() float64
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c Circle) Area() float64 {
	return math.Pi * (c.radius * c.radius)
}
