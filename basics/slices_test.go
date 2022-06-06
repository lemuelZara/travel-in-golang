package basics

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Slices(t *testing.T) {
	t.Run("should return a nil slice with zero length", func(t *testing.T) {
		var slice []int

		assert.Len(t, slice, 0)
		assert.Nil(t, slice)
	})

	t.Run("should append new values inside slice", func(t *testing.T) {
		s1 := []string{"lemuel"}
		s2 := []string{"coelho", "zara"}

		s1 = append(s1, s2...)
		s1 = append(s1, "a", "b", "c")

		assert.Equal(t, []string{"lemuel", "coelho", "zara", "a", "b", "c"}, s1)
	})

	t.Run("should cut slice", func(t *testing.T) {
		names := []string{"João", "Pedro", "José", "Maria", "Joana"}

		women := names[3:5]
		men := names[0:3]

		assert.Equal(t, []string{"Maria", "Joana"}, women)
		assert.Equal(t, []string{"João", "Pedro", "José"}, men)
	})

	t.Run("should created new slice is the same pointer address from first slice", func(t *testing.T) {
		names := []string{"Maria", "João", "Pedro", "José", "Joana"}
		newNames := names[2:4]
		newNames2 := newNames[:1]

		assert.Equal(t, &newNames2[0], &names[2])
	})

	t.Run("should increase slice capacity", func(t *testing.T) {
		names := []string{"Maria", "João", "Pedro", "José", "Joana"}

		assert.Equal(t, 5, len(names))
		assert.Equal(t, 5, cap(names))

		names = append(names, "Lemuel")

		assert.Equal(t, 6, len(names))
		assert.Equal(t, 10, cap(names))
	})

	t.Run("should correct manipulate slices", func(t *testing.T) {
		nums := []int{56, 89, 15, 25, 30, 50}
		mine := append([]int(nil), nums[:3]...)
		mine[0], mine[1], mine[2] = -50, -100, -150

		assert.Equal(t, []int{-50, -100, -150}, mine)
		assert.Equal(t, []int{56, 89, 15}, nums[:3])

		items := []string{
			"pacman", "mario", "tetris", "doom", "galaga", "frogger",
			"asteroids", "simcity", "metroid", "defender", "rayman",
			"tempest", "ultima",
		}
		items = items[5:8]
		sort.Strings(items)

		assert.Equal(t, []string{"asteroids", "frogger", "simcity"}, items)
	})
}
