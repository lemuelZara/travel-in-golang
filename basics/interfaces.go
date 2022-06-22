package basics

import (
	"fmt"
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

/* - - - - - */

type Book struct {
	title string
	price float64
}

func (b Book) Print() string {
	return b.title
}

type Game struct {
	title string
	price float64
}

func (g Game) Print() string {
	return g.title
}

func (g *Game) Discount(d float64) {
	g.price -= (g.price * d)
}

type Puzzle struct {
	title string
	price float64
}

func (p Puzzle) Print() string {
	return p.title
}

type Toy struct {
	title string
	price float64
}

func (t Toy) Print() string {
	return t.title
}

func (t *Toy) Discount(d float64) {
	t.price -= (t.price * d)
}

type Printer interface {
	Print() string
}

type List []Printer

func (l List) Print() (r string) {
	if len(l) == 0 {
		return "Sorry"
	}

	for _, v := range l {
		r += fmt.Sprintf("%s\n", v.Print())
	}

	return r
}

func (l List) Discount(d float64) {
	for _, it := range l {
		if it, ok := it.(interface{ Discount(float64) }); ok {
			it.Discount(d)
		}
	}
}
