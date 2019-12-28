package pointers_errors

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, w Wallet, expected Bitcoin) {
		t.Helper()
		var actual Bitcoin = w.GetBalance()

		if actual != expected {
			t.Errorf("actual: %s expected: %s", actual, expected)
		}
	}

	assertError := func(t *testing.T, err error, want error) {
		t.Helper()

		if err == nil {
			t.Fatal("wanted an error, but got none.")
		}

		if err != want {
			t.Errorf("got: %q, want %q", err, want)
		}
	}

	assertNoError := func(t *testing.T, err error) {
		t.Helper()
    if err != nil {
        t.Fatal("got an error but didn't want one")
    }
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(100))

		// Assert that nothing was withdrawn from the wallet.
		assertBalance(t, wallet, Bitcoin(20))

		// Assert if the Error is the correct error.
		assertError(t, err, ErrInsufficientFunds)
	})
}
