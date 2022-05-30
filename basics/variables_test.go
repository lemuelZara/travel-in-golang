package basics

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func Test_Variables(t *testing.T) {
	t.Run("should return correct value the variable", func(t *testing.T) {
		expctString, expctInt, expctFloat, expctBool, expctTime := Variables()

		assert.Equal(t, expctString, "Lemuel")
		assert.Equal(t, expctInt, 25)
		assert.Equal(t, expctFloat, 34.65)
		assert.Equal(t, expctBool, true)
		assert.Equal(t, expctTime, time.Date(2022, 1, 1, 0, 0, 0, 0, time.Local))
	})

	t.Run("should return correct type the variable", func(t *testing.T) {
		expctString, expctInt, expctFloat, expctBool, expctTime := Variables()

		assert.Equal(t, fmt.Sprintf("%T", expctString), "string")
		assert.Equal(t, fmt.Sprintf("%T", expctInt), "int")
		assert.Equal(t, fmt.Sprintf("%T", expctFloat), "float64")
		assert.Equal(t, fmt.Sprintf("%T", expctBool), "bool")
		assert.Equal(t, fmt.Sprintf("%T", expctTime), "time.Time")
	})
}
