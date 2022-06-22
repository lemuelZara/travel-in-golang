package basics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mobydick  = Book{title: "moby dick", price: 10.0}
	minecraft = Game{title: "minecraft", price: 20.0}
	tetris    = Game{title: "tetris", price: 15.0}
	rubik     = Puzzle{title: "rubik's cube", price: 5.0}
	yoda      = Toy{title: "yoda", price: 55.0}
)

func Test_Interfaces(t *testing.T) {
	tests := []struct {
		name     string
		result   interface{}
		expected interface{}
	}{
		{
			name:     "should return perimeter of rectangle",
			result:   checkPerimeter(Rectangle{10, 10}),
			expected: 40.0,
		},
		{
			name:     "should return area of rectangle",
			result:   checkArea(Rectangle{10, 10}),
			expected: 100.0,
		},
		{
			name:     "should return perimeter of circle",
			result:   checkPerimeter(Circle{10}),
			expected: 62.83185307179586,
		},
		{
			name:     "should return perimeter of circle",
			result:   checkArea(Circle{10}),
			expected: 314.1592653589793,
		},
		{
			name:     "should return correct store values",
			result:   checkStoreValuesList(),
			expected: "moby dick\nminecraft\ntetris\nrubik's cube\nyoda\n",
		},
		{
			name:   "should discount store list values",
			result: checkStoreDiscountList(),
			expected: List{
				Book{title: "moby dick", price: 10},
				&Game{title: "minecraft", price: 10},
				&Game{title: "tetris", price: 7.5},
				Puzzle{title: "rubik's cube", price: 5},
				&Toy{title: "yoda", price: 27.5},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.result)
		})
	}
}

func checkPerimeter(shape Shape) float64 {
	return shape.Perimeter()
}

func checkArea(shape Shape) float64 {
	return shape.Area()
}

func checkStoreValuesList() string {
	var store List

	store = append(store, mobydick, minecraft, tetris, rubik, yoda)

	return store.Print()
}

func checkStoreDiscountList() List {
	var store List

	store = append(store, mobydick, &minecraft, &tetris, rubik, &yoda)

	store.Discount(.5)

	return store
}
