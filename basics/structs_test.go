package basics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Structs(t *testing.T) {
	t.Run("should add users", func(t *testing.T) {
		userID := Add(User{ID: 123})

		assert.Equal(t, userID, GetAll()[2].ID)
	})

	t.Run("should get all persons", func(t *testing.T) {
		assert.Equal(t, 3, len(GetAll()))
	})
}
