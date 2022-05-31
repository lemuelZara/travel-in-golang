package basics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Conditionals(t *testing.T) {
	t.Run("should return correct type from a variable", func(t *testing.T) {
		result := GetTypeFromVariable(true)
		expected := "bool"

		assert.Equal(t, result, expected)

		result = GetTypeFromVariable(1)
		expected = "int"

		assert.Equal(t, result, expected)

		result = GetTypeFromVariable(6.7)
		expected = "float64"

		assert.Equal(t, result, expected)

		result = GetTypeFromVariable("a")
		expected = "string"

		assert.Equal(t, result, expected)

		result = GetTypeFromVariable([]int{1})
		expected = "Don't know type []int"

		assert.Equal(t, result, expected)
	})
}
