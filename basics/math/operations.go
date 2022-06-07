package math

import "math"

func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func ToFixed(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))

	return float64(Round(num*output)) / output
}

func Sum(num1, num2 float64) float64 {
	return num1 + num2
}

func Sub(num1, num2 float64) float64 {
	return ToFixed(num1-num2, 2)
}

func Mult(num1, num2 float64) float64 {
	return ToFixed(num1*num2, 2)
}

func Div(num1, num2 float64) float64 {
	return ToFixed(num1/num2, 2)
}
