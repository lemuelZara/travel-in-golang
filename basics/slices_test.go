package basics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Slices(t *testing.T) {
	t.Run("should return a nil slice with zero length", func(t *testing.T) {
		var slice []int

		assert.Len(t, slice, 0)
		assert.Nil(t, slice)
	})

	t.Run("should return true if comparable slices length is equal", func(t *testing.T) {
		s1 := []string{"a"}
		s2 := []string{"b"}

		assert.True(t, len(s1) == len(s2))
	})

	t.Run("should return false if comparable slices length is not equal", func(t *testing.T) {
		s1 := []string{"abc", "dpa", "sdd"}
		s2 := []string{}

		assert.False(t, len(s1) == len(s2))
	})

	t.Run("should append new values inside slice", func(t *testing.T) {
		s1 := []string{"lemuel"}
		s2 := []string{"coelho", "zara"}

		s1 = append(s1, s2...)
		s1 = append(s1, "a", "b", "c")

		assert.Equal(t, []string{"lemuel", "coelho", "zara", "a", "b", "c"}, s1)
	})
}
