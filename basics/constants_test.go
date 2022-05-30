package basics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Constants(t *testing.T) {
	t.Run("should return correct value for the constants", func(t *testing.T) {
		expctFloat, expctString := Constants()

		assert.Equal(t, expctFloat, 2201.29)
		assert.Equal(t, expctString, "fkgo_2")
	})
}
