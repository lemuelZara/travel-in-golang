package basics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Slices(t *testing.T) {
	t.Run("should sum all nums inside slice", func(t *testing.T) {
		r := Sum([]int{1, 2, 3, 4, 5, 6})
		expected := 21

		assert.Equal(t, expected, r)
	})

	t.Run("should return a slice containing the sum of the other slices", func(t *testing.T) {
		n1 := []int{1, 2, 3}
		n2 := []int{4, 5, 6}

		r := SumAll(n1, n2)
		expected := []int{6, 15}

		assert.Equal(t, expected, r)
	})

	t.Run("should create nil slice", func(t *testing.T) {
		var s []int

		assert.Nil(t, s)
		assert.Equal(t, 0, len(s))
		assert.Equal(t, 0, cap(s))
	})

	t.Run("should grow a slice", func(t *testing.T) {
		var s []int

		s = append(s, 1, 2)
		assert.Equal(t, 2, len(s))
		assert.Equal(t, 2, cap(s))

		s = append(s, 3, 4)
		assert.Equal(t, 4, len(s))
		assert.Equal(t, 4, cap(s))

		s = append(s, 3, 4)
		assert.Equal(t, 6, len(s))
		assert.Equal(t, 8, cap(s))
	})

	t.Run("should cut a slice", func(t *testing.T) {
		var s []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		cut := s[:5]

		assert.Equal(t, []int{1, 2, 3, 4, 5}, cut)
		assert.Equal(t, 5, len(cut))
		assert.Equal(t, 10, cap(cut))

		cut = s[7:]

		assert.Equal(t, []int{8, 9, 10}, cut)
		assert.Equal(t, 3, len(cut))
		assert.Equal(t, 3, cap(cut))

		cut = s[2:5:5]

		assert.Equal(t, []int{3, 4, 5}, cut)
		assert.Equal(t, 3, len(cut))
		assert.Equal(t, 3, cap(cut))
	})

	t.Run("should share the same backing array", func(t *testing.T) {
		s := []int{1, 2, 3, 4, 5}

		// {2, 3, 4, 5}
		s2 := s[1:5]

		// {3, 4, 5}
		s3 := s2[1:4]

		// {4, 5}
		s4 := s3[1:3]

		assert.Equal(t, 2, len(s4))
		assert.Equal(t, 2, cap(s4))
		assert.Equal(t, &s[3], &s4[0])
	})

	t.Run("should preallocate the backing array", func(t *testing.T) {
		s := make([]string, 0, 5)

		assert.Equal(t, 0, len(s))
		assert.Equal(t, 5, cap(s))
	})

	t.Run("should copy elements between slices", func(t *testing.T) {
		s1 := []string{"Lemuel", "Coelho", "Zara"}
		s2 := make([]string, len(s1))

		copy(s2, s1)

		assert.Equal(t, []string{"Lemuel", "Coelho", "Zara"}, s2)
	})
}
