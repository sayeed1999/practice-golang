package pointersnerrors

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("insufficient funds, cannot withdraw")

type Stringer interface {
	String() string
}

// Let's make a Bitcoin type from an existig type int
type Bitcoin int

// Implement the Stringer interface on type Bitcoin
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

// Note: - If we dont use pointer for the methods, with each method there will be a copy of the Wallet
// this means that any changes made to the Wallet in the method will not be reflected outside of that method.
// - Using a pointer receiver allows us to modify the original Wallet instance directly.

// Task: Try removing the pointers and run the tests like func (w Wallet) Deposit(amount int) {}

// Getter for balance (Conventionally, we need a separate getter method!)
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}

	w.balance -= amount
	return nil
}
