package basics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Hello(t *testing.T) {
	t.Run("should return correct salutation", func(t *testing.T) {
		expectedInPT := SayHello("Lemuel", "pt")
		expectedInEN := SayHello("Lemuel", "en")
		expectedInES := SayHello("Lemuel", "es")
		expectedInFR := SayHello("Lemuel", "fr")

		assert.Equal(t, expectedInPT, "Olá, Lemuel!")
		assert.Equal(t, expectedInEN, "Hello, Lemuel!")
		assert.Equal(t, expectedInES, "Hola, Lemuel!")
		assert.Equal(t, expectedInFR, "Bonjour, Lemuel!")
	})
}
