package pointersnerrors

type Wallet struct {
	balance int
}

// Note: - If we dont use pointer for the methods, with each method there will be a copy of the Wallet
// this means that any changes made to the Wallet in the method will not be reflected outside of that method.
// - Using a pointer receiver allows us to modify the original Wallet instance directly.

// Task: Try removing the pointers and run the tests like func (w Wallet) Deposit(amount int) {}

// Getter for balance (Conventionally, we need a separate getter method!)
func (w *Wallet) Balance() int {
	return w.balance
}

func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}
