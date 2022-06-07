package basics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Interfaces(t *testing.T) {
	tests := []struct {
		name     string
		result   float64
		expected float64
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
