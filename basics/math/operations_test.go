package math

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Operations(t *testing.T) {
	t.Run("should round float number", func(t *testing.T) {
		expected := 12
		result := Round(12.43)

		assert.Equal(t, result, expected)
	})

	t.Run("should fixed float number", func(t *testing.T) {
		expected := 49.3919
		result := ToFixed(49.3918687, 4)

		assert.Equal(t, result, expected)
	})

	t.Run("should return correct sum", func(t *testing.T) {
		expected := 7.33
		result := Sum(2.34, 4.99)

		assert.Equal(t, result, expected)
	})

	t.Run("should return correct subtraction", func(t *testing.T) {
		expected := -2.65
		result := Sub(2.34, 4.99)

		assert.Equal(t, result, expected)
	})

	t.Run("should return correct multiplication", func(t *testing.T) {
		expected := 11.68
		result := Mult(2.34, 4.99)

		assert.Equal(t, result, expected)
	})

	t.Run("should return correct division", func(t *testing.T) {
		expected := 0.47
		result := Div(2.34, 4.99)

		assert.Equal(t, result, expected)
	})
}
