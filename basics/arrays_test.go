package basics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Arrays(t *testing.T) {
	t.Run("should return ordered array", func(t *testing.T) {
		result := BubbleSort([]int{10, 3, 2, 1, 8})
		expected := []int{1, 2, 3, 8, 10}

		assert.Equal(t, result, expected)
	})

	t.Run("should return last element for array", func(t *testing.T) {
		result := GetLastArrayElement([]int{10, 3, 2, 1, 8})
		expected := 8

		assert.Equal(t, result, expected)
	})

	t.Run("should return person by name", func(t *testing.T) {
		result := GetAverageGradeByStudentName("Maria")
		expected := 8.333333333333334

		assert.Equal(t, result, expected)
	})
}
