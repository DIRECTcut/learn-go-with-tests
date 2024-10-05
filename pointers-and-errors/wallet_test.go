package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(5))

		assertBalance(t, wallet, Bitcoin(5))
		assertNoError(t, err)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(15))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but did not want one")
	}
}

func assertError(t *testing.T, got error, want error) {
	if got == nil {
		t.Fatal("wanted an error but did not get one")
	}

	if got != want {
		t.Errorf("error message is incorrect, want '%s', got '%s'", want, got)
	}
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("Want %s got %s", want, got)
	}
}
