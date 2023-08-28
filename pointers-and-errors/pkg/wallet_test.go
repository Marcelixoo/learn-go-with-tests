package banking_test

import (
	"testing"

	banking "github.com/Marcelixoo/learn-go-with-tests/pointers-and-errors/pkg"
)

func TestWallet(t *testing.T) {
	t.Run("it receives deposits", func(t *testing.T) {
		wallet := banking.NewWallet()

		wallet.Deposit(banking.Bitcoin(10))

		assertBalance(t, wallet, banking.Bitcoin(10))
	})

	t.Run("it allows withdraws", func(t *testing.T) {
		wallet := banking.NewWallet()
		wallet.Deposit(banking.Bitcoin(20))

		err := wallet.Withdraw(banking.Bitcoin(10))

		assertBalance(t, wallet, banking.Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("it prevents overdrafts", func(t *testing.T) {
		t.Run("balance is kept intact", func(t *testing.T) {
			wallet := banking.NewWallet()
			initial := banking.Bitcoin(20)

			wallet.Deposit(initial)
			wallet.Withdraw(banking.Bitcoin(100))

			assertBalance(t, wallet, initial)
		})
		t.Run("an error is returned", func(t *testing.T) {
			wallet := banking.NewWallet()

			wallet.Deposit(banking.Bitcoin(20))

			assertError(t, wallet.Withdraw(banking.Bitcoin(100)), banking.ErrInsufficientFunds)
		})
	})
}

func assertBalance(t testing.TB, wallet *banking.Wallet, want banking.Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got == nil {
		t.Fatal("expected to get an error but didn't")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Errorf("unexpected error found %q", got)
	}
}
