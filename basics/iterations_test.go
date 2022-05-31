package basics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Iterations(t *testing.T) {
	t.Run("should repeat word at times", func(t *testing.T) {
		result := RepeatWordAtTimes("a", 5)
		expected := "aaaaa"

		assert.Equal(t, result, expected)
	})

	t.Run("should find even number by range", func(t *testing.T) {
		start := 10
		end := 20
		result := FindEvenNumbersByRange(start, end)
		expected := []int{10, 12, 14, 16, 18, 20}

		assert.Equal(t, result, expected)
	})
}
