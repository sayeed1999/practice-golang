package pointersnerrors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(30)}
		wallet.Withdraw(Bitcoin(20))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(30))

		if err == nil {
			t.Errorf("expected an error for insufficient funds, got none")
		}
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got '%d', want '%d'", got, want)
	}
}
