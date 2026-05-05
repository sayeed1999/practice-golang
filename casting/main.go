package main

func main() {
	balanceInterface := getBalance()
	// [Success] This will succeed because balanceInterface is actually a *Balance
	balance, ok := balanceInterface.(*Balance)
	if !ok {
		panic("Failed to cast to Balance")
	}

	println("Amount:", balance.Amount)
	println("Currency:", balance.Currency)

	balanceCopy := balanceInterface.(*BalanceCopy)
	// [Failure] This will cause a panic due to incompatible types
	// but why?? Reason:
	// The reason for the panic is that Balance and BalanceCopy are two different struct types, even though they have the same fields. In Go, you cannot directly cast between different struct types, even if they have identical fields. The type assertion will fail because the underlying types are not the same, leading to a panic at runtime.
	println("Amount:", balanceCopy.Amount)
	println("Currency:", balanceCopy.Currency)
}

type Balance struct {
	Amount   int    "json:amount"
	Currency string "json:currency"
}

type BalanceCopy struct {
	Amount   int    "json:amount"
	Currency string "json:currency"
}

func getBalance() interface{} {
	return &Balance{Amount: 100, Currency: "USD"}
}
