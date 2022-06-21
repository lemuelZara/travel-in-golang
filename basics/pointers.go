package basics

import (
	"errors"
	"fmt"
)

var ErrInsufficientBalance = errors.New("insufficient balance to complete this transaction!")

type Bitcoin float64

type Wallet struct {
	balance Bitcoin
}

// (w *Wallet) is pointer for wallet
func (w *Wallet) Deposit(v Bitcoin) {
	fmt.Printf("(Deposit) Balance memory address: %v\n", &w.balance)

	w.balance += v
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(v Bitcoin) error {
	if v > w.balance {
		return ErrInsufficientBalance
	}

	w.balance -= v

	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}
