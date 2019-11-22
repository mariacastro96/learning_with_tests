package wallet

import "testing"

func TestWallet(t *testing.T) {
	walletBalance := func(t *testing.T, wallet Wallet, want int) {
		t.Helper()
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}

	errorHandling := func(t *testing.T, err error, want string) {
		t.Helper()

		if err == nil {
			t.Error("wanted an error but didn't get one")
		}

		if err.Error() != want {
			t.Errorf("wanted error message: %q anf got %q", want, err)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		want := 10

		walletBalance(t, wallet, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{20}

		wallet.Withdraw(10)

		want := 10

		walletBalance(t, wallet, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{20}

		err := wallet.Withdraw(100)

		want := 20
		wantedError := "nope"

		walletBalance(t, wallet, want)

		errorHandling(t, err, wantedError)

	})

}
