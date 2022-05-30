package math

import "math"

func Sum(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func Sub(num1 float64, num2 float64) float64 {
	result := math.Trunc(num1 - num2)

	return result
}

func Mult(num1 float64, num2 float64) float64 {
	return num1 * num2
}

func Div(num1 float64, num2 float64) float64 {
	return num1 / num2
}
