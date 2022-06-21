package basics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Pointers(t *testing.T) {
	t.Run("should deposit money in wallet", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		result := wallet.Balance()
		expected := Bitcoin(10)

		assert.Equal(t, expected, result)
	})

	t.Run("should not remove money in wallet", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)
		err := wallet.Withdraw(15)

		if err != nil {
			expected := ErrInsufficientBalance

			assert.Equal(t, expected, err)
		}
	})

	t.Run("should remove money in wallet", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)
		_ = wallet.Withdraw(5)

		result := wallet.Balance()
		expected := Bitcoin(5)

		assert.Equal(t, expected, result)
	})
}
